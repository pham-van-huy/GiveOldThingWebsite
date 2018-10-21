// import { Route } from 'react-router-dom'
import Main from '../components/Main'
import Home from '../components/Home'
import Login from '../containers/login-container'
import CreatePost from '../components/CreatePost'
import DetailPost from '../components/DetailPost'
import NotAuthentication from '../components/NotAuthentication'
import RequireAuth from '../components/RequireAuth'

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
                // routes: [
                //   { path: '/child/:id/grand-child',
                //     component: GrandChild
                //   }
                // ]
            },
            {
                path: '/posts/:id',
                component: DetailPost,
                // routes: [
                //   { path: '/child/:id/grand-child',
                //     component: GrandChild
                //   }
                // ]
            },
            { path: '/login', component: NotAuthentication(Login) }
        ]
    }
]

export default routes
