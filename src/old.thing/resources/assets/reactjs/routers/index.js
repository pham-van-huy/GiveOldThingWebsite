// import { Route } from 'react-router-dom'
import Main from '../components/Main'
import Home from '../components/Home'
import Login from '../containers/login-container'
import CreatePost from '../components/post/CreatePost'
import DetailPost from '../components/post/DetailPost'
import NotAuthentication from '../components/NotAuthentication'
import RequireAuth from '../components/RequireAuth'
import Registry from '../components/Registry'

const routes = [
    {
        component: Main,
        routes: [
            {
                path: '/',
                exact: true,
                component: NotAuthentication(Home)
            },
            {
                path: '/posts/create',
                component: RequireAuth(CreatePost),
            },
            {
                path: '/posts/:id',
                component: DetailPost,
            },
            { path: '/login', component: NotAuthentication(Login) },
            { path: '/registry', component: NotAuthentication(Registry) }
        ]
    }
]

export default routes
