import React from 'react'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

const Home = () => (
  <div className="row">
    <h1>This is Home page</h1>
    <Link to="/add">add</Link>
  </div>
)

export default Home
