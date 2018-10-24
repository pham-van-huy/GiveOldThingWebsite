import React from 'react'
import {lifecycle, withState} from 'recompose'
import { getApi } from "../../helper"
import ReactPaginate from 'react-paginate'
import { Link, NavLink } from 'react-router-dom'
import './ListPosts.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import ButtonControlList from './ButtonControlList'
import ItemInList from './ItemInList'

const ListPostGrid = ({posts}) => {
    return posts.map(post => {
        var firstImage = null 
        if (post.Images.length != 0) {
            firstImage = "/uploads/" + post.Images[0].Link
        }
        return (
            <div className="col-12 col-md-6 col-lg-4" key={post.ID}>
                <div className="card">
                    <img className="card-img-top" src={firstImage} alt="Card image cap" />
                    <div className="card-body">
                        <h4 className="card-title">
                            <Link to={`/posts/` + post.ID} className="navbar-brand" title="View Product">{post.Title}</Link>
                        </h4>
                        <p className="card-text">{post.Description}</p>
                        <div className="row">
                            <div className="col">
                                <p className="btn btn-danger btn-block">{post.Price}</p>
                            </div>
                            <div className="col">
                                <a href="#" className="btn btn-success btn-block">Add to cart</a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    });
}

const ListPostCompo = ({posts}) => {
    return posts.map(post => {
        return (
            <ItemInList post={post} key={post.ID}/>
        )
    });
}

class ListPosts extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        posts: [],
        total_page: 1,
        perPage: 10,
        page: 1,
        styleUi: 'grid', // list
      };
    }
    loadPostsFromServer() {
        var data = {
            page: this.state.page,
            limit: this.state.perPage,
        }
        getApi('/api/posts', data).then(response => {
            if (response.data.success) {
                this.setState({ posts: response.data.data.records });
                this.setState({ total_page: response.data.data.total_page });
            }
        })
    }
    componentDidMount() {
        this.loadPostsFromServer()
    }
    handlePageClick(data) {
        let selected = data.selected;
    
        this.setState({page: selected}, () => {
          this.loadPostsFromServer();
        });
    }
    clickCondition(e) {

    }
    clickChooseStyle(e, t) {
        console.log('vao e', e, t)
        this.setState({styleUi: e})
    }
    render() {
        return (
            <React.Fragment>
                <div className="row">
                    <div className="w-100 condition">
                        <div className="tab-condition">
                            <div className="tab-each float-left active">
                                <a className="inner-tab" onClick={this.clickCondition.bind(this)}>
                                    All
                                </a>
                            </div>
                            <div className="tab-each float-left">
                                <a className="inner-tab" onClick={this.clickCondition.bind(this)}>
                                    Bán
                                </a>
                            </div>
                            <div className="tab-each float-left">
                                <a className="inner-tab" onClick={this.clickCondition.bind(this)}>
                                    Mua
                                </a>
                            </div>
                        </div>
                        <div className="control float-right">
                            <ButtonControlList active={this.state.styleUi == 'grid'} onClick={this.clickChooseStyle.bind(this, 'grid')} icon={['fas', 'th']}/>
                            <ButtonControlList active={this.state.styleUi == 'list'} onClick={this.clickChooseStyle.bind(this, 'list')} icon={['fas', 'list-ul']}/>
                        </div>
                    </div>
                    {this.state.styleUi == 'grid' ? <ListPostGrid posts={this.state.posts}/>: <ListPostCompo posts={this.state.posts}/>}
                </div>
                <ReactPaginate previousLabel={"previous"}
                    nextLabel={"next"}
                    breakLabel={<a href="">...</a>}
                    breakClassName={"break-me"}
                    pageCount={this.state.total_page}
                    marginPagesDisplayed={2}
                    pageRangeDisplayed={5}
                    onPageChange={this.handlePageClick}
                    containerClassName={"pagination"}
                    subContainerClassName={"pages pagination"}
                    activeClassName={"active"} 
                    pageClassName={"page-item"}
                    pageLinkClassName={"page-link"}
                    previousClassName={"page-item"}
                    previousLinkClassName={"page-link"}
                    nextClassName={"page-item"}
                    nextLinkClassName={"page-link"}
                    />
            </React.Fragment>
        )
    }
}

export default ListPosts
