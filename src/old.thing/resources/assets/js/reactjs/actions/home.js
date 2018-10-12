import { getApi } from "../helper";
export const atcListPost = () => {
    return (dispatch => {
        // dispatch(fetchProductsBegin());
        getApi('/api/list-post').then(res => {
            console.log(res)
        })
        .then(res => res.json())
        .then(json => {
            console.log(json.products)
            dispatch(fetchPostSuccess(json.products));
            return json.products;
        })
    })()
}

export const fetchPostSuccess = posts => ({
    type: 'FETCH_POST_SUCCESS',
    payload: { posts }
});

