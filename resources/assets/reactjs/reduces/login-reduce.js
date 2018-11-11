export const SAVE_TOKEN = 'SAVE_TOKEN'
export const UNAUTHENTICATED = 'UNAUTHENTICATED'
export const AUTHENTICATED = 'AUTHENTICATED'
export const AUTHENTICATION_ERROR = 'AUTHENTICATION_ERROR'
export const SAVE_INFO_USER = 'SAVE_INFO_USER'

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

export function saveInfoUser(info) {
    localStorage.setItem('infoUser', JSON.stringify(info))
    return {
        type: SAVE_INFO_USER,
        info: info
    }
}

const login = (state = { authenticated: false, infoUser: JSON.parse(localStorage.getItem('infoUser'))}, action) => {
    switch (action.type) {
        case SAVE_TOKEN:
            return { ...state, token: action.token, authenticated: true }
        case AUTHENTICATED:
            return { ...state, authenticated: true };
        case UNAUTHENTICATED:
            return { ...state, token: null, authenticated: false };
        case AUTHENTICATION_ERROR:
            return { ...state, token: null, error: action.payload };
        case SAVE_INFO_USER:
            return { ...state, infoUser: action.info }
        default:
            return state
    }
}

export default login
