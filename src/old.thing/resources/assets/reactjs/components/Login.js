import React from 'react'

class Login extends React.Component {
    constructor(props) {
        super(props)
        console.log(123)
    }

    render() {
        return (
            <div id="loginModal" className="modal show" tabIndex="-1" role="dialog" aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h2 className="text-center"><img src="https://bootdey.com/img/Content/avatar/avatar1.png" className="img-circle" /><br/>Login</h2>
                        </div>
                        <div className="modal-body">
                            <form className="form col-md-12 center-block">
                                <div className="form-group">
                                    <input type="text" className="form-control input-lg" placeholder="Email" />
                                </div>
                                <div className="form-group">
                                    <input type="password" className="form-control input-lg" placeholder="Password" />
                                </div>
                                <div className="form-group">
                                    <button className="btn btn-primary btn-lg btn-block">Sign In</button>
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
