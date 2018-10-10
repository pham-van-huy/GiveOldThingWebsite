export const postApi = (url = '', data = {}) => {
    return fetch(url, {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, cors, *same-origin
        cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        credentials: "same-origin", // include, same-origin, *omit
        headers: {
            "Content-Type": "application/json; charset=utf-8",
            "Authorization": "Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzkyNDgwODYsImlhdCI6MTUzODk4ODg4Niwic3ViIjoiMCJ9.HEQtLf4yIaH-ylPll-g-BZCfWi0_tcAMXk3zpeeHUo1ekOJ8NqVbOLgY-8acqtgzTDDniM50IlR0gKepu0Ma8TxD_Vgx_Aio5TaHZdfnon5MbD03ZxFekq8Z1d2E5LDdQPR_qzaZG48LadZQlXgBBA0f8GjsCNRCf6u6HImwnpD0pOp9C41Wmc1hm56M05VMQd6o2Y82Vo8KY0llCMzG2nP3lSbKPQmBzCgRiWtC79o_G80wIEsxnZXVJN59EzQug6pjPOwXbS7gWPhdoVPtMUc4dYZkmuXMGOLmFXw2Y18dKnGIRwkN2wW1Qjm7dUtwxsvC8qbdQfv42K5zgD9OPQ",
        },
        redirect: "follow", // manual, *follow, error
        referrer: "no-referrer", // no-referrer, *client
        body: JSON.stringify(data), // body data type must match "Content-Type" header
    })
}

export const getApi = (url = '') => {
    return fetch(url, {
        method: "GET", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, cors, *same-origin
        cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        credentials: "same-origin", // include, same-origin, *omit
        headers: {
            "Content-Type": "application/json; charset=utf-8",
            "Authorization": "Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzkyNDgwODYsImlhdCI6MTUzODk4ODg4Niwic3ViIjoiMCJ9.HEQtLf4yIaH-ylPll-g-BZCfWi0_tcAMXk3zpeeHUo1ekOJ8NqVbOLgY-8acqtgzTDDniM50IlR0gKepu0Ma8TxD_Vgx_Aio5TaHZdfnon5MbD03ZxFekq8Z1d2E5LDdQPR_qzaZG48LadZQlXgBBA0f8GjsCNRCf6u6HImwnpD0pOp9C41Wmc1hm56M05VMQd6o2Y82Vo8KY0llCMzG2nP3lSbKPQmBzCgRiWtC79o_G80wIEsxnZXVJN59EzQug6pjPOwXbS7gWPhdoVPtMUc4dYZkmuXMGOLmFXw2Y18dKnGIRwkN2wW1Qjm7dUtwxsvC8qbdQfv42K5zgD9OPQ",
        },
        redirect: "follow", // manual, *follow, error
        referrer: "no-referrer", // no-referrer, *client
    })
}
