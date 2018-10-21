import React from 'react'
import { withRouter } from 'react-router-dom'
import Select from 'react-select'
import _ from 'lodash'
import ImageUploader from '../externalCompo/react-images-upload'
import { postApi } from "../helper"
import MyMap from "./MyMap"

class CreatePost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            description: "",
            cities: [],
            selectedCity: 0,
            price: 0,
            categories: [],
            selectedCate: 0,
            images: [],
            isShowSuccess: false,
        };
    }
    onDrop(pictureFiles, pictureDataURLs) {
        this.setState({
            images: this.state.images.concat(pictureFiles),
        });
    }
    handleChangeCate(e) {
        this.setState({ selectedCate: e.target.value })
    }
    handleChangeCity(e) {
        this.setState({ selectedCity: e.target.value })
    }
    componentDidMount() {
        var options = [];

        fetch('/web/cities').then(res => res.json()).then((response) => {
            if (response.success) {
                _.forEach(response.data, (val) => {
                    options.push({ value: val.ID, label: val.Name })
                });
                this.setState({ cities: options })
            }
        })
        var options1 = [];
        fetch('/web/categories').then(res => res.json()).then((response) => {
            if (response.success) {
                _.forEach(response.data, (val) => {
                    options1.push({ value: val.ID, label: val.Name })
                });
                this.setState({ categories: options1 })
            }
        })
    }
    onChangeTitle(e) {
        this.setState({ title: e.target.value })
    }
    onChangeDes(e) {
        this.setState({ description: e.target.value })
    }
    onSubmit(e) {
        e.preventDefault();
        var opts = {
            Title: this.state.title,
            Description: this.state.description,
            CityId: this.state.selectedCity,
            CategoryId: this.state.selectedCate,
            UserId: 1,
            "Images[]": this.state.images,
            Price: this.state.Price
        }
        postApi("/api/posts/create", opts, "formdata").then((response) => {
            this.setState({isShowSuccess : true})
            this.props.history.push('/')
        })
    }
    onChangePrice(e) {
        this.setState({ price: e.target.value })
    }
    render() {
        var { selectedCity, selectedCate } = this.state
        var successMessage = null
        if (this.state.isShowSuccess) {
            successMessage = (<div className="alert alert-success">
                <strong>Create successfully !!</strong>
            </div>)
        }
        var cssBelowTitle = {
            marginTop: "14px"
        }
        return (
            <div className="container">
                <div className="row justify-content-center">
                    <div className="col-md-6 center-block">
                        <div className="page-header">
                            <h1>Tạo tin</h1>
                        </div>
                        <form id="create-post" onSubmit={this.onSubmit.bind(this)}>
                            {successMessage}
                            <div className="row">
                                <div className="city col-md-6">
                                    <select value={selectedCity} className="form-control" onChange={this.handleChangeCity.bind(this)}>
                                        {
                                            this.state.cities.map((val, i) => {
                                                return (<option key={i} value={val.value}>{val.label}</option>)
                                            })
                                        }
                                    </select>
                                </div>
                                <div className="category col-md-6">
                                    <select value={selectedCate} className="form-control" onChange={this.handleChangeCate.bind(this)}>
                                        {
                                            this.state.categories.map((val, i) => {
                                                return (<option key={i} value={val.value}>{val.label}</option>)
                                            })
                                        }
                                    </select>
                                </div>
                            </div>
                            <div className="form-group" style={cssBelowTitle}>
                                <input type="text" name="title" onChange={this.onChangeTitle.bind(this)} className="form-control" id="title" placeholder="Tieu de" />
                            </div>
                            <div className="form-group">
                                <textarea type="text" name="description" onChange={this.onChangeDes.bind(this)} className="form-control" id="description" placeholder="Mo ta" />
                            </div>
                            <div className="form-group">
                                <input type="number" name="price" onChange={this.onChangePrice.bind(this)} className="form-control" id="price" placeholder="Price" />
                            </div>
                            <div className="form-group">
                                <ImageUploader
                                    withPreview={true}
                                    withIcon={true}
                                    buttonText='Choose images'
                                    onChange={this.onDrop.bind(this)}
                                    imgExtension={['.jpg', '.gif', '.png', '.gif', '.jpeg']}
                                    maxFileSize={5242880}
                                />
                            </div>
                            <div className="form-group">
                                <MyMap />
                            </div>
                            <button type="submit" className="btn btn-primary">Create</button>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default CreatePost
