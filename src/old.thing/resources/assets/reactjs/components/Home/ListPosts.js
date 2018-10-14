import React from 'react'
import {lifecycle, withState} from 'recompose'
import { getApi } from "../../helper"

class ListPosts extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        posts: []
      };
    }
  
    componentDidMount() {
        getApi('/api/posts').then(response => {
            if (response.data.success) {
                this.setState({ posts: response.data.data.records });
            }
        })
    }
    render() {
        var listPostsElem = this.state.posts.map(post => {
            var firstImage = null 
            if (post.Images.length != 0) {
                firstImage = '/uploads/' + post.Images[0].Link
            }
            return (
                <div className="col-12 col-md-6 col-lg-4" key={post.ID}>
                    <div className="card">
                        <img className="card-img-top" src={firstImage} alt="Card image cap" />
                        <div className="card-body">
                            <h4 className="card-title"><a href="product.html" title="View Product">{post.Title}</a></h4>
                            <p className="card-text">{post.Description}</p>
                            <div className="row">
                                <div className="col">
                                    <p className="btn btn-danger btn-block">99.00 $</p>
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
        return (
            <div className="row">
                {listPostsElem}
            </div>
        )
    }
}

export default ListPosts
