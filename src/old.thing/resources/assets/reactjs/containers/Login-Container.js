import { connect } from 'react-redux'
import Login from '../components/Login'
import { withRouter } from 'react-router-dom'
import {saveToken} from '../reduces/login-reduce'

const mapDispatchToProps = (dispatch) => ({
    saveToken: token => dispatch(saveToken(token))
})

const mapStateToProps = state => ({
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(Login))
