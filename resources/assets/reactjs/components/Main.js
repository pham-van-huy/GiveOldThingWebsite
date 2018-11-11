import React from 'react'
import Header from './Header'
import Footer from './Footer'
import { Route, Switch, Link, Router } from 'react-router-dom'
import { renderRoutes } from 'react-router-config'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'

library.add(fas)

const Main = ({ route }) => (
    <div id="wrap-container">
        <Header />
            {renderRoutes(route.routes)}
        <Footer />
    </div>
)

export default Main
