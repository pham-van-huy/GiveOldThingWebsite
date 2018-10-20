import { combineReducers } from 'redux'
import todos from './todos'
import login from './login-reduce'

export default combineReducers({
  todos,
  login
})
