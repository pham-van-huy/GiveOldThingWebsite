import React from 'react'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Sidebar from './Sidebar.js'
import ListPost from "../containers/Home/ListPost";

const Home = () => (
  <div>
    <div className="row row-offcanvas row-offcanvas-right">
      <div className="col-12 col-md-9">
        <div className="row">
          <ListPost></ListPost>
          <div className="col-6 col-lg-4">
            <h2>Heading</h2>
            <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo,
                tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem
                        malesuada magna mollis euismod. Donec sed odio dui. </p>
            <p><a className="btn btn-secondary" href="#"
              role="button">View details »</a></p>
          </div>
        </div>
      </div>

      <Sidebar />
    </div>

  </div>
)

export default Home
