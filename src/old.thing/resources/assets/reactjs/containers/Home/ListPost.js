import { connect } from 'react-redux'
import DomListPost from '../../components/Home/ListPosts'
import { actListPost } from '../../actions/home.js'
import { withRouter } from 'react-router-dom'

const mapDispatchToProps = (dispatch) => ({
    funcGetListPost: () => getApi('/api/list-post').then(res => res.json()).then(json => {
        dispatch(fetchPostSuccess(json.data))
    })
})

const mapStateToProps = state => ({
    posts: state.posts
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(DomListPost))
