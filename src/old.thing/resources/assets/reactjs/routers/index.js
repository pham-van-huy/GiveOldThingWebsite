// import { Route } from 'react-router-dom'
import Main from '../components/Main'
import Home from '../components/Home'
import AddTodo from '../containers/AddTodo'
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
      }
    ]
  }
]

export default routes
