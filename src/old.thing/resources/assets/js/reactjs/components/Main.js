import React from 'react'
import AddTodo from '../containers/AddTodo'
import Header from './Header'
import Footer from './Footer'
import { Route, Switch, Link, Router } from 'react-router-dom'
import { renderRoutes } from 'react-router-config'

const Main = ({ route }) => (
    <div className="container">
        <Header />
        {renderRoutes(route.routes)}
        <Footer />
    </div>
)

export default Main
