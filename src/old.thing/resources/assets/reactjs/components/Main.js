import React from 'react'
import Header from './Header'
import Footer from './Footer'
import { Route, Switch, Link, Router } from 'react-router-dom'
import { renderRoutes } from 'react-router-config'

const Main = ({ route }) => (
    <div id="wrap-container">
        <Header />
            {renderRoutes(route.routes)}
        <Footer />
    </div>
)

export default Main
