
import { connect } from 'react-redux'
import Root from '../components/Root'
import * as actions from '../actions'

const mapStateToProps = (state, ownProps) => {
  return {
    isAuthenticated: state.auth.isAuthenticated,
    errorMessage: state.auth.errorMessage
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    // connect: () => {
    connect () {
      return dispatch(actions.connect())
    },
    disconnect () {
      return dispatch(actions.disconnect())
    },
    dispatch
  }
}

const RootContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(Root)

export default RootContainer
