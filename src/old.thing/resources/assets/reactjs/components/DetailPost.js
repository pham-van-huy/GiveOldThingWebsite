import React from 'react'
import { withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import Select from 'react-select'
import _ from 'lodash'
import ImageUploader from '../externalCompo/react-images-upload'
import { postApi, getApi } from '../helper'
import MyMap from "./MyMap"
import './DetailPost.scss'
import Slider from 'react-slick'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

class Responsive extends React.Component {
    render() {
        var settings = {
            dots: true,
            infinite: false,
            speed: 500,
            slidesToShow: 4,
            slidesToScroll: 4,
            initialSlide: 0,
            responsive: [
                {
                    breakpoint: 1024,
                    settings: {
                        slidesToShow: 3,
                        slidesToScroll: 3,
                        infinite: true,
                        dots: true
                    }
                },
                {
                    breakpoint: 600,
                    settings: {
                        slidesToShow: 2,
                        slidesToScroll: 2,
                        initialSlide: 2
                    }
                },
                {
                    breakpoint: 480,
                    settings: {
                        slidesToShow: 1,
                        slidesToScroll: 1
                    }
                }
            ]
        };
        return (
            <div>
                <h2> Responsive </h2>
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
                    <div>
                        <h3>7</h3>
                    </div>
                    <div>
                        <h3>8</h3>
                    </div>
                </Slider>
            </div>
        );
    }
}

class DetailPost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            ID: 0,
            Title: "",
            Description: "",
            SelectedCity: 0,
            Price: 0,
            SelectedCate: 0,
            Images: [],
            IsShowSuccess: false,
        };
    }
    componentDidMount() {
        const { params } = this.props.match;
        getApi('/api/posts/' + params.id, 'json', false).then((response) => {
            if (response.status == 200 && response.data.success) {
                var post = response.data.data
                var generalAttri = _.pick(post, ['ID', 'Title', 'Description', 'Price', 'Images']);
                this.setState(generalAttri)
            }
        })
    }
    render() {
        var settings = {
            dots: true,
            infinite: true,
            speed: 500,
            slidesToShow: 1,
            slidesToScroll: 1
        };
        var listImageCompo = this.state.Images.map((value, key) => {
            var src = "/uploads/" + value.Link
            return (
                <div key={key} style={{height : "340px"}}>
                    <img className="image-slide" src={src} style={{height : "340px"}}/>
                </div>
            )
        })
        return (
            <React.Fragment>
                <div className="container-fluid" id="info-product">
                    <div className="row">
                        <Slider {...settings} className="w-100">
                            {listImageCompo}
                        </Slider>
                    
                        <div className="card w-100">
                            <div className="card-body">
                                <h5 className="card-title">{this.state.Title}</h5>
                                <p className="card-text">{this.state.Description}</p>
                                <p className="card-text price"><FontAwesomeIcon icon={['fas', 'money-bill-alt']} /> {this.state.Price + ' D'}</p>
                                <p className="card-text phone"><FontAwesomeIcon icon={['fas', 'phone']} /> <strong>{'0932047956'}</strong></p>
                                <a href="#" className="btn btn-primary">Add to bookmark</a>
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

const mapDispatchToProps = (dispatch) => ({

})

const mapStateToProps = state => ({
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(DetailPost))
