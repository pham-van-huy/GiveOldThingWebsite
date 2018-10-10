import { connect } from 'react-redux'
import DomListPost from '../../components/Home/ListPosts'
import { atcListPost } from '../../actions/home.js'
import { withRouter } from 'react-router-dom'

const mapDispatchToProps = () => ({
    funcGetListPost: () => atcListPost()
})

const mapStateToProps = state => ({
    posts: state.posts
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(DomListPost))
