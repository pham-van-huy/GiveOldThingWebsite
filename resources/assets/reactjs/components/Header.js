import React from 'react'
import { NavLink, withRouter} from 'react-router-dom'
import { connect } from 'react-redux'

class Header extends React.Component {
    constructor(props) {
        super(props);
        this.state = {

        };
    }

    componentDidMount() {
    }

    moveCreatePost(e) {
        this.props.history.push('/posts/create')
    }

    render() {
        let auth = (
            <li className="nav-item">
                <NavLink className="nav-link" to={`/login`} activeClassName="active">Login</NavLink>
            </li>
        )

        if (this.props.authenticated) {
            auth = (
                <li className="nav-item dropdown">
                    <a className="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
                        aria-haspopup="true" aria-expanded="false">
                       bombay
                </a>
                    <div className="dropdown-menu" aria-labelledby="navbarDropdown">
                        <a className="dropdown-item" href="#">Profile</a>
                        <div className="dropdown-divider"></div>
                        <a className="dropdown-item" href="#">Logout</a>
                    </div>
                </li>
            )
        }
        return (
            <nav className="navbar navbar-expand-lg navbar-light bg-light">
                <NavLink to={`/`} className="navbar-brand" activeClassName="active">Logo</NavLink>
                <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
                    aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>

                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item">
                            <a className="nav-link" href="#">Link</a>
                        </li>
                        <li className="nav-item dropdown">
                            <a className="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
                                aria-haspopup="true" aria-expanded="false">
                                Dropdown
                </a>
                            <div className="dropdown-menu" aria-labelledby="navbarDropdown">
                                <a className="dropdown-item" href="#">Action</a>
                                <a className="dropdown-item" href="#">Another action</a>
                                <div className="dropdown-divider"></div>
                                <a className="dropdown-item" href="#">Something else here</a>
                            </div>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="javascript:" onClick={this.moveCreatePost.bind(this)}>CreatePost</a>
                        </li>
                        {auth}
                    </ul>
                    <form className="form-inline my-2 my-lg-0">
                        <input className="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search" />
                        <button className="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
                    </form>
                </div>
            </nav>
        )
    }
}

const mapDispatchToProps = (dispatch) => ({
})

const mapStateToProps = state => ({
    infoUser: state.login.infoUser,
    authenticated: state.login.authenticated
})

export default withRouter(connect(
    mapStateToProps,
    mapDispatchToProps
)(Header))

