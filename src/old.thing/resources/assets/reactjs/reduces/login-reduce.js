export const SAVE_TOKEN = 'SAVE_TOKEN'
export const UNAUTHENTICATED = 'UNAUTHENTICATED'
export const AUTHENTICATED = 'AUTHENTICATED'
export const AUTHENTICATION_ERROR = 'AUTHENTICATION_ERROR'

export function saveToken(token) {
    localStorage.setItem('token', token)
    return {
        type: SAVE_TOKEN,
        token: token,
    }
}

export function signOut() {
    localStorage.removeItem('token');
    return {
        type: UNAUTHENTICATED
    };
}

export function checkLogin() {
    var token = localStorage.getItem('token')
    if (token) {
        return {
            type: SAVE_TOKEN,
            token: token,
        }
    } else {
        return {
            type: UNAUTHENTICATED
        };
    }
}

const login = (state = {authenticated: false}, action) => {
    switch (action.type) {
        case SAVE_TOKEN:
            return { ...state, token: action.token, authenticated: true }
        case AUTHENTICATED:
            return { ...state, authenticated: true };
        case UNAUTHENTICATED:
            return { ...state, token: null, authenticated: false };
        case AUTHENTICATION_ERROR:
            return { ...state, token: null, error: action.payload };
        default:
            return state
    }
}

export default login
