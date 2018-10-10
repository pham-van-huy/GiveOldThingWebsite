const handlerHome = (state = {}, action) => {
    switch (action.type) {
        case 'FETCH_POST_SUCCESS':
            return state = action.payload.posts
        case 'TOGGLE_TODO':
            return state.map(todo =>
                (todo.id === action.id)
                    ? { ...todo, completed: !todo.completed }
                    : todo
            )
        default:
            return state
    }
}

export default handlerHome
