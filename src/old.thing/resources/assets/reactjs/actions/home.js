import { getApi } from "../helper";
export const actListPost = () => {
    return (dispatch) => {
        getApi('/api/posts').then(json => {
            dispatch(fetchPostSuccess(json.products));
            return json.products;
        })
    }
}

export const fetchPostSuccess = posts => ({
    type: 'FETCH_POST_SUCCESS',
    payload: { posts }
});

