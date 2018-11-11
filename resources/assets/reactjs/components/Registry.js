import React from 'react'
import {postApi} from '../helper'
import { Link, withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import { saveToken, saveInfoUser } from '../reduces/login-reduce'

class Registry extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            fullName: '',
            userName: '',
            email: '',
            password: '',
            comfirmPassword: '',
            message: null
        }
        this.handleChangeInput = this.handleChangeInput.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChangeInput(e) {
        this.setState({[e.target.name]: e.target.value})
    }

    componentDidMount(){

    }
    handleSubmit(e) {
        e.preventDefault()
        var data = {
            Name: this.state.fullName,
            Username: this.state.userName,
            Email: this.state.email,
            Password: this.state.password
        }

        postApi('/registry', data, 'json', false).then((response) => {
            if (response.status == 200) {
                this.props.history.push('/login', {message: response.data.message})
            } else {
                this.setState({message: response.data.message})
            }
        })
    }
    render() {
        let message = ''

        if (this.state.message != null) {
            message = (
                <p style={{color: "red"}}>{this.state.message}</p>
            )
        }

        return (
            <div id="loginModal" className="modal show" tabIndex="-1" role="dialog" aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h2 className="text-center"><img src="https://bootdey.com/img/Content/avatar/avatar1.png" className="img-circle" /><br/>Login</h2>
                        </div>
                        <div className="modal-body">
                            <form className="form col-md-12 center-block" onSubmit={this.handleSubmit.bind(this)}>
                                {message}
                                <div className="form-group">
                                    <input value={this.state.fullName} onChange={this.handleChangeInput} type="text" className="form-control input-lg" placeholder="Full name" name="fullName" required/>
                                </div>
                                <div className="form-group">
                                    <input value={this.state.email} onChange={this.handleChangeInput} type="email" className="form-control input-lg" placeholder="Email" name="email" required/>
                                </div>
                                <div className="form-group">
                                    <input value={this.state.username} onChange={this.handleChangeInput} type="text" className="form-control input-lg" placeholder="UserName" name="userName" required/>
                                </div>
                                <div className="form-group">
                                    <input value={this.state.password} onChange={this.handleChangeInput} type="password" className="form-control input-lg" placeholder="Password" name="password" required/>
                                </div>
                                <div className="form-group">
                                    <input value={this.state.comfirmPassword} onChange={this.handleChangeInput} type="password" className="form-control input-lg" placeholder="Password" name="comfirmPassword" required/>
                                </div>
                                <div className="form-group">
                                    <button type="submit" className="btn btn-primary btn-lg btn-block">Register</button>
                                    <span className="pull-right"><Link to="/login">Login</Link></span><span><a href="#">Need help?</a></span>
                                </div>
                            </form>
                        </div>
                        <div className="modal-footer">
                            <div className="col-md-12">
                                <button className="btn" data-dismiss="modal" aria-hidden="true">Cancel</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

const mapDispatchToProps = (dispatch) => ({
    saveToken: token => dispatch(saveToken(token)),
    saveInfoUser: infor => dispatch(saveInfoUser(infor))
})

const mapStateToProps = state => ({
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(Registry))
