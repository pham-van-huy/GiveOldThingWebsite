import React from 'react'
import {postApi} from '../helper'
class Login extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            username: '',
            password: '',
        }
    }
    handleChangeEmail(e) {
        this.setState({username: e.target.value})
    }
    handleChangePassword(e) {
        this.setState({password: e.target.value})
    }
    componentDidMount(){
        
    }
    handleSubmit(e) {
        e.preventDefault()
        var data = {
            Username: this.state.username,
            Password: this.state.password
        }
        postApi('/token-auth', data, 'json', false).then((response) => {
            this.props.saveToken(response.data.token)
        })
    }
    render() {
        return (
            <div id="loginModal" className="modal show" tabIndex="-1" role="dialog" aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h2 className="text-center"><img src="https://bootdey.com/img/Content/avatar/avatar1.png" className="img-circle" /><br/>Login</h2>
                            <span className="row">For developement, after run seed, username: phann123, password: 123</span>
                        </div>
                        <div className="modal-body">
                            <form className="form col-md-12 center-block" onSubmit={this.handleSubmit.bind(this)}>
                                <div className="form-group">
                                    <input value={this.state.username} onChange={this.handleChangeEmail.bind(this)} type="text" className="form-control input-lg" placeholder="Email" />
                                </div>
                                <div className="form-group">
                                    <input value={this.state.password} onChange={this.handleChangePassword.bind(this)} type="password" className="form-control input-lg" placeholder="Password" />
                                </div>
                                <div className="form-group">
                                    <button type="submit" className="btn btn-primary btn-lg btn-block">Sign In</button>
                                    <span className="pull-right"><a href="#">Register</a></span><span><a href="#">Need help?</a></span>
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

export default Login
