
import React, { PropTypes } from 'react'
import { dateToString } from '../../utils.js'

import iconFastBackward from '../../../img/icon-fastBackward.png'
import iconBackward from '../../../img/icon-backward.png'
import iconPause from '../../../img/icon-pause.png'
import iconForward from '../../../img/icon-forward.png'
import iconFastForward from '../../../img/icon-fastForward.png'

import iconFastBackwardHover from '../../../img/icon-fastBackward-hover.png'
import iconBackwardHover from '../../../img/icon-backward-hover.png'
import iconPauseHover from '../../../img/icon-pause-hover.png'
import iconForwardHover from '../../../img/icon-forward-hover.png'
import iconFastForwardHover from '../../../img/icon-fastForward-hover.png'

class ControlPanel extends React.Component {

  constructor (props) {
    super(props)
    this.state = {
    }
  }

  render () {

    const { 
      currentDate,
      playbackMode,
      selectPlaybackMode
    } = this.props

    return (
      <div className="control-panel">
        <div className="control-panel_date-counter">
          { dateToString(new Date(currentDate)) }
        </div>
        <ul className="control-panel_playback-selector">
          <li
            className="control-panel_playback-selector_button"
            onClick={ () => selectPlaybackMode('fastBackward') }
            style={{
              backgroundColor: playbackMode === 'fastBackward' ? '#D0462C' : 'white'
            }}
          >
            <img
              src={ playbackMode === 'fastBackward' ? iconFastBackwardHover : iconFastBackward }
            />
          </li>
          <li
            className="control-panel_playback-selector_button"
            onClick={ () => selectPlaybackMode('backward') }
            style={{
              backgroundColor: playbackMode === 'backward' ? '#D0462C' : 'white'
            }}
          >
            <img
              src={ playbackMode === 'backward' ? iconBackwardHover : iconBackward }
            />
          </li>
          <li
            className="control-panel_playback-selector_button"
            onClick={ () => selectPlaybackMode('pause') }
            style={{
              backgroundColor: playbackMode === 'pause' ? '#D0462C' : 'white'
            }}
          >
            <img
              src={ playbackMode === 'pause' ? iconPauseHover : iconPause }
            />
          </li>
          <li
            className="control-panel_playback-selector_button"
            onClick={ () => selectPlaybackMode('forward') }
            style={{
              backgroundColor: playbackMode === 'forward' ? '#D0462C' : 'white'
            }}
          >
            <img
              src={ playbackMode === 'forward' ? iconForwardHover : iconForward }
            />
          </li>
          <li
            className="control-panel_playback-selector_button"
            onClick={ () => selectPlaybackMode('fastForward') }
            style={{
              backgroundColor: playbackMode === 'fastForward' ? '#D0462C' : 'white'
            }}
          >
            <img
              src={ playbackMode === 'fastForward' ? iconFastForwardHover : iconFastForward }
            />
          </li>
        </ul>
        <div className="control-panel_focus-selector">
          <p>Map Focus:</p>
          <ul>
            <li>
              Expedition
            </li>
            <li>
              Member
            </li>
            <li>
              Documents
            </li>
          </ul>
        </div>
      </div>
    )
  }

}

ControlPanel.propTypes = {

}

export default ControlPanel