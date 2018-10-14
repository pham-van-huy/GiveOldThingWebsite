import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import { createStore } from 'redux'
import rootReducer from './reduces'
import Main from './components/Main'
import { BrowserRouter } from 'react-router-dom'
import routes from './routers'
import { renderRoutes } from 'react-router-config'

const store = createStore(rootReducer)

render(
	<Provider store={store}>
        <BrowserRouter>
            {renderRoutes(routes)}
        </BrowserRouter>
	</Provider>,
	document.getElementById('app')
)
