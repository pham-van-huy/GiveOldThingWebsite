import React from 'react'
import { withRouter } from 'react-router-dom'
import Select from 'react-select'
import _ from 'lodash'
import ImageUploader from 'react-images-upload'
import './CreatePost.css'
import { postApi } from "../helper";

class CreatePost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            title: "",
            description: "",
            cities: [],
            selectedCity: 'select',
            categories: [],
            selectedCate: 'select',
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
            "Images[]": this.state.images
        }
        postApi("/api/posts/create", opts, "formdata").then((response) => {
            console.log('test 1', response)
            this.setState({isShowSuccess : true})
        })
    }
    render() {
        var { selectedCity, selectedCate } = this.state
        var successMessage = null
        if (this.state.isShowSuccess) {
            successMessage = (<div className="alert alert-success">
                <strong>Create successfully !!</strong>
            </div>)
        }
        return (
            <div className="container">
                <div className="col-md-6">
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
                        <div className="form-row">
                            <div className="form-group col-md-6">
                                <label htmlFor="title">Title</label>
                                <input type="text" name="title" onChange={this.onChangeTitle.bind(this)} className="form-control" id="title" placeholder="Tieu de" />
                            </div>
                        </div>
                        <div className="form-group">
                            <label htmlFor="description">Description</label>
                            <textarea type="text" name="description" onChange={this.onChangeDes.bind(this)} className="form-control" id="description" placeholder="" />
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
                        <button type="submit" className="btn btn-primary">Create</button>
                    </form>
                </div>
            </div>
        );
    }
}

export default CreatePost
