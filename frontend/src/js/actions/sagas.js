
import { delay } from 'redux-saga';
import {
    all, put, take, race, takeLatest
    // takeEvery, select, all, call
} from 'redux-saga/effects';

import _ from 'lodash';

import * as ActionTypes from './types';
import FkApi from '../api/calls';

import { changePlaybackMode, focusExpeditionTime, chartDataLoaded } from './index';

import { PlaybackModes } from '../components/PlaybackControl';

import { FkGeoJSON } from '../common/geojson';

export function* walkGeoJson(summary, geojson) {
    const expedition = new FkGeoJSON(geojson);

    const start = new Date(summary.startTime);
    const end = new Date(summary.endTime);

    const tickDuration = 100;
    const walkDurationInSeconds = 30;
    const walkLength = end - start;
    const walkLengthInSeconds = (walkLength / 1000);

    let walkSecondsPerTick = (walkLengthInSeconds / walkDurationInSeconds) / (1000 / tickDuration);

    console.log("Start", start);
    console.log("End", end);
    console.log("TickDuration", tickDuration, "WalkDuration", walkDurationInSeconds);
    console.log("SecondsPerTick", walkSecondsPerTick);

    let time = start;

    while (true) {
        while (time < end) {
            const coordinates = expedition.getCoordinatesAtTime(time);

            yield put(focusExpeditionTime(time, coordinates, walkSecondsPerTick));

            const [ activityAction, focusFeatureAction, playbackAction, _ ] = yield race([
                take(ActionTypes.USER_MAP_ACTIVITY),
                take(ActionTypes.FOCUS_FEATURE),
                take(ActionTypes.CHANGE_PLAYBACK_MODE),
                delay(tickDuration)
            ]);

            if (activityAction || focusFeatureAction) {
                yield put(changePlaybackMode(PlaybackModes.Pause));
                break;
            }

            if (playbackAction) {
                let nextAction = playbackAction;
                if (playbackAction.mode === PlaybackModes.Pause) {
                    nextAction = yield take(ActionTypes.CHANGE_PLAYBACK_MODE);
                }
                switch (nextAction.mode) {
                case PlaybackModes.Beginning: {
                    time = expedition.getFirst().time();
                    break;
                }
                case PlaybackModes.Rewind: {
                    walkSecondsPerTick = -5;
                    break;
                }
                case PlaybackModes.Play: {
                    walkSecondsPerTick = 5;
                    break;
                }
                case PlaybackModes.Forward: {
                    walkSecondsPerTick = 20;
                    break;
                }
                case PlaybackModes.End: {
                    time = expedition.getLast().time();
                    break;
                }
                default: {
                    break;
                }
                }
            }

            time.setTime(time.getTime() + (walkSecondsPerTick * 1000));
        }

        console.log("DONE", start, end, time);

        time = expedition.getFirst().time();

        yield put(changePlaybackMode(PlaybackModes.Pause));

        yield take(ActionTypes.CHANGE_PLAYBACK_MODE);
    }
}

export function* refreshSaga(pagedGeojson) {
    while (true) {
        pagedGeojson = yield FkApi.getNextExpeditionGeoJson(pagedGeojson);
        if (pagedGeojson.geo.features.length === 0) {
            yield delay(10000);
        }
    }
}

export function* loadCharts() {
    yield takeLatest([ActionTypes.CHART_DATA_LOAD], function* (chartAction) {
        let page = yield FkApi.getSourceGeoJson(chartAction.chart.source.inputId);
        while (page.hasMore) {
            page = yield FkApi.getNextSourceGeoJson(page);
        }
        yield put(chartDataLoaded(chartAction.chart))
    });
}

export function* loadExpeditionDetails() {
    const cache = {};

    yield takeLatest([ActionTypes.API_EXPEDITION_SOURCES_GET.SUCCESS], function* (sourcesAction) {
        const ids = _(sourcesAction.response.deviceInputs).map('id').uniq().value();
        const newIds = _.difference(ids, _.keys(cache));
        const sources = yield all(newIds.map(id => FkApi.getSource(id)));
        const indexed = _(sources).keyBy('id').value();
        const features = yield all(sources.map(source => FkApi.getFeatureGeoJson(source.lastFeatureId)));
        console.log("LatestFeatures", features);
        Object.assign(cache, indexed);
    });
}

export function* loadActiveExpedition(projectSlug, expeditionSlug) {
    const [ expedition, pagedGeoJson, sources ] = yield all([
        FkApi.getExpedition(projectSlug, expeditionSlug),
        FkApi.getExpeditionGeoJson(projectSlug, expeditionSlug),
        FkApi.getExpeditionSources(projectSlug, expeditionSlug)
    ]);

    console.log(expedition, pagedGeoJson, sources);

    if (pagedGeoJson.geo.features.length > 0) {
        yield delay(1000)
        yield all([ walkGeoJson(expedition, pagedGeoJson.geo), refreshSaga(pagedGeoJson) ]);
    }
}

export function* loadSingleDevice(deviceId) {
    const [ pagedGeoJson, source ] = yield all([
        FkApi.getSourceGeoJson(deviceId),
        FkApi.getSource(deviceId),
    ]);

    if (pagedGeoJson.geo.features.length > 0) {
        const features = yield all([FkApi.getFeatureGeoJson(source.lastFeatureId)]);
        console.log("LatestFeatures", features);

        yield delay(1000)
        yield all([ walkGeoJson(source, pagedGeoJson.geo), refreshSaga(pagedGeoJson) ]);
    }
}

export function* loadActiveProject() {
    const projectSlug = 'www';

    const [ project, expeditions ] = yield all([
        FkApi.getProject(projectSlug),
        FkApi.getProjectExpeditions(projectSlug)
    ]);

    console.log(project);

    if (expeditions.length > 0) {
        yield all([
            loadExpeditionDetails(),
            loadCharts(),
            loadSingleDevice(125),
            // loadActiveExpedition(projectSlug, expeditions[0].slug)
        ])
        console.log("Done")
    }
}

export function* rootSaga() {
    yield all([
        loadActiveProject()
    ]);
}
