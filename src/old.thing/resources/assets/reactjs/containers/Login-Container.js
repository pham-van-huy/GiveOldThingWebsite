import { connect } from 'react-redux'
import Login from '../components/Login'
import { withRouter } from 'react-router-dom'

const mapDispatchToProps = (dispatch) => ({
})

const mapStateToProps = state => ({
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(Login))
