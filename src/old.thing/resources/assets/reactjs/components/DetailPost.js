import React from 'react'
import { withRouter } from 'react-router-dom'
import Select from 'react-select'
import _ from 'lodash'
import ImageUploader from '../externalCompo/react-images-upload'
import { postApi } from "../helper"
import MyMap from "./MyMap"
import './DetailPost.scss'
import Slider from "react-slick"

class DetailPost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            description: "",
            selectedCity: 0,
            price: 0,
            selectedCate: 0,
            images: [],
            isShowSuccess: false,
        };
    }
    componentDidMount() {
    }
    render() {
        var settings = {
            dots: true,
            infinite: true,
            speed: 500,
            slidesToShow: 1,
            slidesToScroll: 1
        };

        return (
            <React.Fragment>
                <div className="container">
                    <Slider {...settings}>
                        <div>
                            <h3>1</h3>
                        </div>
                        <div>
                            <h3>2</h3>
                        </div>
                        <div>
                            <h3>3</h3>
                        </div>
                        <div>
                            <h3>4</h3>
                        </div>
                        <div>
                            <h3>5</h3>
                        </div>
                        <div>
                            <h3>6</h3>
                        </div>
                    </Slider>
                    <div className="row">
                        <div className="col-md-8">
                            <div className="product-title">Corsair GS600 600 Watt PSU</div>
                            <div className="product-desc">The Corsair Gaming Series GS600 is the ideal price/performance choice for mid-spec gaming PC</div>
                            <div className="product-rating">
                                <i className="fa fa-star gold"></i>
                                <i className="fa fa-star gold"></i>
                                <i className="fa fa-star gold"></i>
                                <i className="fa fa-star gold"></i>
                                <i className="fa fa-star-o"></i>
                            </div>
                            <hr />
                            <div className="product-price">$ 1234.00</div>
                            <div className="product-stock">In Stock</div>
                            <hr />
                            <div className="btn-group cart">
                                <button type="button" className="btn btn-success">
                                    Add to cart
						</button>
                            </div>
                            <div className="btn-group wishlist">
                                <button type="button" className="btn btn-danger">
                                    Add to bookmark
						</button>
                            </div>
                        </div>
                    </div>
                </div>

                <div className="container-fluid">
                    <div className="col-md-12 product-info">
                        <ul id="myTab" className="nav nav-tabs" role="tablist">
                            <li className="nav-link active"><a href="#service-one" data-toggle="tab">DESCRIPTION</a></li>
                            <li className="nav-link"><a href="#service-two" data-toggle="tab">PRODUCT INFO</a></li>
                            <li className="nav-link"><a href="#service-three" data-toggle="tab">REVIEWS</a></li>
                        </ul>
                        <div id="myTabContent" className="tab-content">
                            <div className="tab-pane fade show active" id="service-one">
                                <section className="container product-info">
                                    <h3>Corsair Gaming Series GS600 Features:</h3>
                                    <li>It supports the latest ATX12V v2.3 standard and is backward compatible with ATX12V 2.2 and ATX12V 2.01 systems</li>
                                    <li>An ultra-quiet 140mm double ball-bearing fan delivers great airflow at an very low noise level by varying fan speed in response to temperature</li>
                                    <li>80Plus certified to deliver 80% efficiency or higher at normal load conditions (20% to 100% load)</li>
                                    <li>0.99 Active Power Factor Correction provides clean and reliable power</li>
                                    <li>Universal AC input from 90~264V — no more hassle of flipping that tiny red switch to select the voltage input!</li>
                                    <li>Extra long fully-sleeved cables support full tower chassis</li>
                                    <li>A three year warranty and lifetime access to Corsair’s legendary technical support and customer service</li>
                                    <li>Over Current/Voltage/Power Protection, Under Voltage Protection and Short Circuit Protection provide complete component safety</li>
                                    <li>Dimensions: 150mm(W) x 86mm(H) x 160mm(L)</li>
                                    <li>MTBF: 100,000 hours</li>
                                    <li>Safety Approvals: UL, CUL, CE, CB, FCC className B, TÜV, CCC, C-tick</li>
                                </section>
                            </div>
                            <div className="tab-pane fade" id="service-two">
                                <section className="container">

                                </section>

                            </div>
                            <div className="tab-pane fade" id="service-three">

                            </div>
                        </div>
                        <hr />
                    </div>
                </div>
            </React.Fragment>
        );
    }
}

export default DetailPost
