import React from 'react'
import classnames from 'classnames'
import { Link, NavLink } from 'react-router-dom'

const DivList = styled.div`
    /* Adapt the colors based on primary prop */
    max-height: 155px;
    display: block;
    .ip-inner{
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        padding: 5px;
        border-bottom: 1px solid #e9ebee;
        margin-bottom: 3px;
        .card-img-top{
            max-height: 150px;
            width: 125px;
        }
        .ip-content{
            flex: 1;
            display: flex;
            flex-direction: column;
            padding-left: 12px;
            .ip-title{
                border: none;
                font-size: 14px;
                line-height: 17px;
                max-height: 35px;
                color: #000;
                margin: 0;
                font-family: Verdana,Arial,sans-serif;
                overflow: hidden;
                -webkit-line-clamp: 2;
                -webkit-box-orient: vertical;
                a{
                    color: inherit;
                }
            }
            .ip-info{
                .list-cate{
                    list-style: none;
                    padding-left: 0px;
                    li {
                        display: inline-block;
                    }
                }
            }
        }
    }
`;

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import styled from 'styled-components'

const ItemInList = ({post}) => {
    var firstImage = null 
    if (post.Images.length != 0) {
        firstImage = '/uploads/' + post.Images[0].Link
    }
    var nameCate = ''
    if (post.Category) {
        nameCate = post.Category.Name
    }
    return (
        <DivList className="col-md-12 item-post" key={post.ID}>
            <div className="ip-inner">
                <img className="card-img-top" src={firstImage} alt="Card image cap" />
                <div className="ip-content">
                    <h4 className="ip-title">
                        <Link to={`/posts/` + post.ID} className="navbar-brand" title="View Product">{post.Title}</Link>
                    </h4>
                    <p className="ip-description">{post.Description}</p>
                    <div className="ip-info">
                        <ul className="list-cate">
                            <li className="name-cate"><FontAwesomeIcon icon={'anchor'}/> {nameCate}</li>
                        </ul>
                    </div>
                </div>
            </div>
        </DivList>
    );
}

export default ItemInList