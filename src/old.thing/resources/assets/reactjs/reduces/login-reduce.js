export const SAVE_TOKEN = 'SAVE_TOKEN'

export function saveToken(token) {
    return {
       type: SAVE_TOKEN,
       token: token,
    }
}

const login = (state = {}, action) => {
    switch (action.type) {
        case SAVE_TOKEN:
            localStorage.setItem('token', action.token)
            return state = {token: action.token}
        default:
            return state
    }
}

export default login
