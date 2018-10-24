import React from 'react'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Sidebar from './Sidebar.js'
import ListPost from "../containers/Home/ListPost";

const Home = () => (
    <div className="container-fluid">
        <div className="row">
            <Sidebar />
            <div className="col">
                <ListPost />
            </div>
        </div>
    </div>
)

export default Home
