import { connect } from 'react-redux'
import Login from '../components/Login'
import { withRouter } from 'react-router-dom'
import { saveToken, saveInfoUser} from '../reduces/login-reduce'

const mapDispatchToProps = (dispatch) => ({
    saveToken: token => dispatch(saveToken(token)),
    saveInfoUser: infor => dispatch(saveInfoUser(infor))
})

const mapStateToProps = state => ({
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(Login))
