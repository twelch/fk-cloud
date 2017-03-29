// @flow weak

import React, { Component } from 'react'
import { Route, Link } from 'react-router-dom'
import ReactModal from 'react-modal';
import log from 'loglevel';

import type {
  Match as RouterMatch,
  Location as RouterLocation,
  RouterHistory
} from 'react-router-dom';

import { ProjectForm } from '../forms/ProjectForm';
import { ProjectExpeditionForm } from '../forms/ProjectExpeditionForm';
import { AdministratorForm } from '../forms/AdministratorForm';
import { FKApiClient } from '../../api/api';

import { RemoveIcon } from '../icons/Icons'
import '../../../css/projectsettings.css'

import type { APIProject, APINewProject, APINewAdministrator, APIAdministrator, APINewExpedition } from '../../api/types';

type Props = {
  project: APIProject;
  administrators: APIAdministrator[];
  onUpdate: (newSlug: ?string) => void;
  onExpeditionCreate: () => void;

  match: RouterMatch;
  location: RouterLocation;
  history: RouterHistory;
}

export class ProjectSettings extends Component {
  props: $Exact<Props>;
  state: {
    administrators: APIAdministrator[]
  }

  constructor(props: Props) {
    super(props);

    this.state = {
      administrators: []
    };

    this.loadAdministrators();
  }

  async loadAdministrators() {
    const { project } = this.props;
    const administratorsRes = await FKApiClient.get().getAdministrators(project.id);
    if (administratorsRes.type === 'ok' && administratorsRes.payload) {
      const administrators = administratorsRes.payload.administrators;
      this.setState({administrators: administrators});
    }
  }

  async onExpeditionCreate(e: APINewExpedition) {
    const { match, project } = this.props;

    const expeditionRes = await FKApiClient.get().createExpedition(project.id, e);
    if (expeditionRes.type === 'ok') {
      this.props.onExpeditionCreate();
    } else {
      return expeditionRes.errors;
    }
  }

  async onAdministratorAdd(e: APINewAdministrator) {
    const { project, match } = this.props;
    const administratorRes = await FKApiClient.get().addAdministrator(project.id, e);
    if (administratorRes.type === 'ok') {
      await this.loadAdministrators();
      this.props.history.push(`${match.url}`);
    } else {
      return administratorRes.errors;
    }
  }

  async onAdministratorDelete(e: APIAdministrator) {
    const { match } = this.props;
    const administratorRes = await FKApiClient.get().deleteAdministrator(e.project_id, e.user_id);
    if (administratorRes.type === 'ok') {
      await this.loadAdministrators();
      this.props.history.push(`${match.url}`);
    } else {
      return administratorRes.errors;
    }
  }

  async onProjectSave(project: APINewProject) {
    const projectRes = await FKApiClient.get().updateProject(this.props.project.id, project);
    if (projectRes.type !== 'ok') {
      return projectRes.errors;
    }

    log.debug('onProjectSave', projectRes, this.props.project);

    if (projectRes.payload.slug != this.props.project.slug && projectRes.payload) {
      this.props.onUpdate(projectRes.payload.slug);
    } else {
      this.props.onUpdate();
    }
  }

  render () {
    const { match, project } = this.props;
    let { administrators } = this.state;
    const projectSlug = project.slug;

    return (
      <div className="project">

        <Route path={`${match.url}/add-administrator`} render={props =>
          <ReactModal isOpen={true} contentLabel="Add Users" className="modal" overlayClassName="modal-overlay">
            <h2>Add User</h2>
            <AdministratorForm
              project={project}
              administrators={this.state.administrators}
              onCancel={() => this.props.history.push(`${match.url}`)}
              onSave={this.onAdministratorAdd.bind(this)} 
              saveText="Add" />
          </ReactModal> } />

        <h1>Project Settings</h1>
        <div className="row">
          <div className="two-columns">
            <h3>Main</h3>
            <ProjectForm
              name={project ? project.name : undefined}
              slug={project ? project.slug : undefined}
              description={project ? project.description : undefined}

              onSave={this.onProjectSave.bind(this)} />
          </div>
          <div className="two-columns">
            <h3>Users</h3>
            <p>
              Users you add to this project have administrative rights. They can create a new expedition, change its settings and add members to it. If you’re trying to add a member to an expedition instead, select an expedition and then go to <i>Teams</i>.
            </p>

            <table className="administrators-table">
              <thead>
                <tr>
                  <th></th>
                  <th>Users ({ administrators.length })</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
              { administrators.map((administrator, i) =>
                <tr key={i}>
                  <td>
                    <div className="user-avatar medium">
                      <img />
                    </div>
                  </td>
                  <td>
                    {administrator.user_id}
                  </td>
                  <td>
                    <div className="bt-icon medium" onClick={this.onAdministratorDelete.bind(this, administrator)}>
                      <RemoveIcon />
                    </div>
                  </td>
                </tr> )}
              </tbody>
            </table>

            <Link className="button secondary" to={`${match.url}/add-administrator`}>Add User</Link>
          </div>
        </div>
      </div>
    )
  }
}
