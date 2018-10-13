import axios from "axios"

export const postApi = (url = '', data = {}, type = 'json') => {
    var token = localStorage.getItem("token")
    var headers = {}
    var dataPush = null
    if (type == "json") {
        headers["Content-Type"] = "application/json; charset=utf-8"
        dataPush = JSON.stringify(data)
    } else if (type == "formdata") {
        headers["Content-Type"] = "multipart/form-data"
        dataPush = new FormData()
        for (var prop in data) {
            if (Array.isArray(data[prop])){
                data[prop].forEach(elem => {
                    dataPush.append(prop, elem)
                });
            } else {
                dataPush.append(prop, data[prop])
            }
        }
    }

    if (token != null) {
        headers["Authorization"] = "Bearer " + token
    } else {
        token = 'eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzk2ODc2NDgsImlhdCI6MTUzOTQyODQ0OCwic3ViIjoiMCJ9.dP7DKwkUHxQJ8dhYWwtdDwEkhzQAkYXfzOlewhDcj4-9URMRkAwGDRLHFagsySAdI0pk1y32GA5tkWkSAFT7fbAiEnciSKsWinJdr2TGWNgTVvISUWyf_-mMAWem_vVf30ajt3W0R5zPoRvgziro8jn4pIPd3rkecSKxz-l_mdhlw21kF9r2qa18RWTU2AblEryrkJAQsO-z791zO_6Eb5ZzcoDu6m5W9696eD8FuUj-Inj5pfYYZZQtLs0x-jEaXP9zuli_JevHl5bQ4VW-iJXImjiMnl5CQ99IMqz0LsGT6Uaivmmsw1gyyeo58r7yCNilhZNAjr4a1NUIKvevBQ'
        headers["Authorization"] = "Bearer " + token
    }

    return axios({
        method: 'post',
        url: url,
        headers: headers,
        data: dataPush
    })
}

export const getApi = (url = '') => {
    var token = localStorage.getItem("token")
    var headers = {
        "Content-Type": "application/json; charset=utf-8",
    }
    if (token != null) {
        headers["Authorization"] = "Bearer " + token
    }
    return fetch(url, {
        method: "GET", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, cors, *same-origin
        cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        credentials: "same-origin", // include, same-origin, *omit
        headers: headers,
        redirect: "follow", // manual, *follow, error
        referrer: "no-referrer", // no-referrer, *client
    }).then(res => res.json())
}
