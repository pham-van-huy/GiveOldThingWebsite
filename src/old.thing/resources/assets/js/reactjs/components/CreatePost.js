import React from 'react'
import { withRouter } from 'react-router-dom'
import Select from 'react-select'
import _ from 'lodash'

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
            images: []
        };
    }
    handleChangeCate(e) {
        this.setState({selectedCate: e.target.value})
    }
    handleChangeCity(e) {
        this.setState({selectedCity: e.target.value})
    }
    componentDidMount() {
        var options = [];
          
        fetch('/web/cities').then(res => res.json()).then((response) => {
            if (response.success) {
                _.forEach(response.data, (val) => {
                    options.push({value: val.ID, label: val.Name})
                });
                this.setState({cities: options})
            }
        })

        fetch('/web/categories').then(res => res.json()).then((response) => {
            if (response.success) {
                _.forEach(response.data, (val) => {
                    options.push({value: val.ID, label: val.Name})
                });
                this.setState({categories: options})
            }
        })
    }
    onSubmit(e) {
        e.preventDefault();
    }
    render() {
        var {selectedCity, selectedCate} = this.state
        
        return (
            <div className="container">
                <div className="page-header">
                    <h1>Tạo tin</h1>      
                </div>
                <form id="create-post" onSubmit="onSubmit">
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
                            <input type="text" name="title" className="form-control" id="title" placeholder="Tieu de" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label htmlFor="description">Description</label>
                        <textarea type="text" name="description" className="form-control" id="description" placeholder="" />
                    </div>
                    <button type="submit" className="btn btn-primary">Create</button>
                </form>
            </div>
        );
    }
}

export default CreatePost
