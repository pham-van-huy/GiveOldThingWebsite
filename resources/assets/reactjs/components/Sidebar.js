import React from 'react'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

const Sidebar = () => (
    <div className="col-md-3" id="sidebar">
        <div className="card bg-light mb-3">
            <div className="card-header bg-primary text-white text-uppercase"><i className="fa fa-list"></i> Categories</div>
            <ul className="list-group category_block">
                <li className="list-group-item"><a href="category.html">Cras justo odio</a></li>
                <li className="list-group-item"><a href="category.html">Dapibus ac facilisis in</a></li>
                <li className="list-group-item"><a href="category.html">Morbi leo risus</a></li>
                <li className="list-group-item"><a href="category.html">Porta ac consectetur ac</a></li>
                <li className="list-group-item"><a href="category.html">Vestibulum at eros</a></li>
            </ul>
        </div>
        <div className="card bg-light mb-3">
            <div className="card-header bg-success text-white text-uppercase">Last product</div>
            <div className="card-body">
                <img className="img-fluid" src="https://dummyimage.com/600x400/55595c/fff" />
                <h5 className="card-title">Product title</h5>
                <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                <p className="bloc_left_price">99.00 $</p>
            </div>
        </div>
    </div>
)

export default Sidebar
