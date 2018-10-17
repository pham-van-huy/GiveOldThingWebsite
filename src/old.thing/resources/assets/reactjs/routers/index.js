// import { Route } from 'react-router-dom'
import Main from '../components/Main'
import Home from '../components/Home'
import Login from '../containers/Login-Container'
import CreatePost from '../components/CreatePost'

const routes = [
  { component: Main,
    routes: [
      { path: '/',
        exact: true,
        component: Home
      },
      { path: '/posts/create',
        component: CreatePost,
        // routes: [
        //   { path: '/child/:id/grand-child',
        //     component: GrandChild
        //   }
        // ]
      },
      {path: '/login', component: Login}
    ]
  }
]

export default routes
