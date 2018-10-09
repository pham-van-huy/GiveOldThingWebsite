import React from 'react'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Sidebar from './Sidebar.js'
import Header from './Header.js'
import Footer from './Footer.js'

const Home = () => (
  <div className="container">
    <div className="row row-offcanvas row-offcanvas-right">
      <div className="col-12 col-md-9">
        <p className="float-right hidden-md-up">
          <button type="button" class="btn btn-primary btn-sm" data-toggle="offcanvas">Toggle nav</button>
        </p>
        <div className="jumbotron">
          <h1>Hello, world!</h1>
          <p>This is an example to show the potential of an offcanvas layout pattern in Bootstrap. Try some
                    responsive-range viewport sizes to see it in action.</p>
        </div>

        <div className="row">
          <div className="col-6 col-lg-4">
            <h2>Heading</h2>
            <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo,
                tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem
                        malesuada magna mollis euismod. Donec sed odio dui. </p>
            <p><a className="btn btn-secondary" href="https://v4-alpha.getbootstrap.com/examples/offcanvas/#"
              role="button">View details »</a></p>
          </div>
        </div>
      </div>

      <Sidebar />
    </div>

      <hr></hr>
      <Footer />

  </div>
)
    
export default Home
