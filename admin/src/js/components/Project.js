// @flow weak

import React, { Component } from 'react'
import { Route, Link } from 'react-router-dom'
import ReactModal from 'react-modal';

import { ProjectForm } from './forms/ProjectForm';
import { ProjectExpeditionForm } from './forms/ProjectExpeditionForm';
import { InputForm } from './forms/InputForm';
import { FKApiClient } from '../api/api';

import type { APIProject, APIExpedition, APINewProject, APINewExpedition, APIInput, APINewInput } from '../api/types';

import '../../css/home.css'

type Props = {
  project: APIProject;
  onProjectUpdate: (newSlug: ?string) => void;

  match: Object;
  location: Object;
  history: Object;
}

export class Project extends Component {
  props: Props;
  state: {
    expeditions: APIExpedition[],
    inputs: APIInput[]
  }

  constructor(props: Props) {
    super(props);

    this.state = {
      expeditions: [],
      inputs: []
    };

    this.loadData();
  }

  async loadData() {
    const expeditionsRes = await FKApiClient.get().getExpeditionsByProjectSlug(this.props.project.slug);
    if (expeditionsRes.type === 'ok' && expeditionsRes.payload) {
      this.setState({ expeditions: expeditionsRes.payload.expeditions || [] })
    }

    const inputsRes = await FKApiClient.get().getInputsByProjectSlug(this.props.project.slug);
    if (inputsRes.type === 'ok' && inputsRes.payload) {
      this.setState({ inputs: inputsRes.payload.inputs || [] })
    }
  }

  async onExpeditionCreate(e: APINewExpedition) {
    const { project } = this.props;

    const expeditionRes = await FKApiClient.get().createExpedition(project.id, e);
    if (expeditionRes.type === 'ok') {
      await this.loadData();
      this.props.history.push(`/projects/${project.slug}`);
    } else {
      return expeditionRes.errors;
    }
  }

  async onInputCreate(i: APINewInput) {
    const { project } = this.props;

    const inputRes = await FKApiClient.get().createInput(project.id, i);
    if (inputRes.type === 'ok') {
      await this.loadData();
      this.props.history.push(`/projects/${project.slug}`);
    } else {
      return inputRes.errors;
    }
  }

  async onProjectSave(project: APINewProject) {
    // TODO: this isn't implemented on the backend yet!

    const projectRes = await FKApiClient.get().updateProject(this.props.project.id, project);
    if (projectRes.type === 'ok') {
      await this.loadData();
    } else {
      return projectRes.errors;
    }

    if (projectRes.slug != this.props.project.slug && projectRes.payload) {
      this.props.onProjectUpdate(projectRes.payload.slug);
    } else {
      this.props.onProjectUpdate();
    }
  }

  render () {
    const { project } = this.props;
    const projectSlug = project.slug;

    return (
      <div className="project">
        <Route path="/projects/:projectSlug/new-expedition" render={() =>
          <ReactModal isOpen={true} contentLabel="New expedition form">
            <h1>Create a new expedition</h1>
            <ProjectExpeditionForm
              projectSlug={projectSlug}
              onCancel={() => this.props.history.push(`/projects/${projectSlug}`)}
              onSave={this.onExpeditionCreate.bind(this)} />
          </ReactModal> } />

        <Route path="/projects/:projectSlug/new-input" render={() =>
          <ReactModal isOpen={true} contentLabel="New input form">
            <h1>Create a new input</h1>
            <InputForm
              projectSlug={projectSlug}
              onCancel={() => this.props.history.push(`/projects/${projectSlug}`)}
              onSave={this.onInputCreate.bind(this)} />
          </ReactModal> } />

        <div id="expeditions">
          <h4>Expeditions</h4>
          { this.state.expeditions.map((e, i) =>
            <div key={`expedition-${i}`} className="expedition-item">
              <Link to={`/projects/${projectSlug}/expeditions/${e.slug}`}>{e.name}</Link>
            </div> )}
          { this.state.expeditions.length === 0 &&
            <span className="empty">No expeditions!</span> }
        </div>
        <Link to={`/projects/${projectSlug}/new-expedition`}>Show new expedition modal</Link>

        <div id="inputs">
          <h4>Inputs</h4>
          { this.state.inputs.map((input, i) =>
            <div key={`input-${i}`} className="input-item">
              { JSON.stringify(input) }
            </div> )}
          { this.state.inputs.length === 0 &&
            <span className="empty">No inputs!</span> }
        </div>
        <Link to={`/projects/${projectSlug}/new-input`}>Show new input modal</Link>

        <h2>Edit project</h2>
        <ProjectForm
          name={project ? project.name : undefined}
          description={project ? project.description : undefined}
          onSave={this.onProjectSave.bind(this)} />
      </div>
    )
  }
}
