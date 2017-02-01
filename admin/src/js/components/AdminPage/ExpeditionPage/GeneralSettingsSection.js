import React, {PropTypes} from 'react'
import {findDOMNode} from 'react-dom'
import { Link } from 'react-router'
import autobind from 'autobind-decorator'
import ContentEditable from 'react-contenteditable'
import I from 'immutable'
import Dropdown from 'react-dropdown'
import Select from 'react-select'
import slug from 'slug'

class GeneralSettingsSection extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      addMemberValue: null,
      inputValues: {}
    }
  }

  render () {
    const {
      projectID,
      expedition,
      setExpeditionProperty,
      setExpeditionPreset,
      saveGeneralSettings,
      errors
    } = this.props

    console.log('aga rendering', expedition)

    return (
      <div id="teams-section" className="section">
        <div className="section-header">
          <h1>Create a New Expedition (1/2)</h1>
        </div>
        <p className="intro">
          Etiam eu purus in urna volutpat ornare. Etiam pretium ante non egestas dapibus. Mauris pretium, nunc non lacinia finibus, dui lectus molestie nulla, quis ultricies libero orci a sapien. Praesent bibendum leo vitae felis pellentesque, sit amet mattis nisi mattis.
        </p>

        <p className="input-label">
          Pick a name for your expedition:
        </p>
        <div className="columns-container">
          <div className="main-input-container">
            <input 
              type="text"
              value={expedition.get('name')}
              onFocus={(e) => {
                if (expedition.get('name') === 'New Expedition') {
                  setExpeditionProperty(['name'], '')
                }
              }}
              onChange={(e) => {
                setExpeditionProperty(['name'], e.target.value)
              }}
            />
            {
              !!errors && !!errors.get('name') &&
              errors.get('name').map((errors, i) => {
                return (
                  <p 
                    key={'errors-' + i}
                    className="errors"
                  >
                    {errors}
                  </p>
                )
              })
            }
          </div>
          <p className="input-description">
            Etiam pretium ante non egestas dapibus. Mauris pretium, nunc non lacinia finibus, dui lectus molestie nulla, quis ultricies libero orci a sapien.
          </p>
        </div>

        <p className="input-label">
          Your expedition will be available at the following address:
        </p>

        <p className="intro">
          {'https://' + projectID + '.fieldkit.org/' + slug(expedition.get('name'))}
        </p>

        {
          !!errors &&
          <p className="errors">
            We found one or multiple errors. Please check your information above or try again later.
          </p>
        }

        <a href="#" onClick={(e) => {
          e.preventDefault()
          saveGeneralSettings()
        }}>
          <div className="button hero">
            Save changes
          </div>
        </a>

      </div>
    )

  }
}

GeneralSettingsSection.propTypes = {
}

export default GeneralSettingsSection
