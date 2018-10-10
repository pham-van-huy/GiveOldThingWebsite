import React from 'react'

const ListPoists = ({ posts, funcGetListPost }) => {
    funcGetListPost()
    return (<ul>
        {posts}
    </ul>)
}

export default ListPoists
