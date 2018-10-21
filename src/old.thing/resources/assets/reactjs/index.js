import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import { createStore, applyMiddleware } from 'redux'
import rootReducer from './reduces'
import Main from './components/Main'
import { BrowserRouter } from 'react-router-dom'
import routes from './routers'
import { renderRoutes } from 'react-router-config'
import reduxThunk from 'redux-thunk'
import { checkLogin, SAVE_TOKEN } from './reduces/login-reduce';

const store = createStore(rootReducer, applyMiddleware(reduxThunk))

const token = localStorage.getItem('token');

if (token) {
  store.dispatch({ type: SAVE_TOKEN, token: token});
}

render(
	<Provider store={store}>
        <BrowserRouter>
            {renderRoutes(routes)}
        </BrowserRouter>
	</Provider>,
	document.getElementById('app')
)
