import axios from "axios"

export const postApi = (url = '', data = {}, type = 'json', isAuthen = true) => {
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

    if (isAuthen) {
        if (token == null) {
            return Promise.reject(new Error('need token'))
        } else {
            headers["Authorization"] = "Bearer " + token
        }
    }

    return axios({
        method: 'post',
        url: url,
        headers: headers,
        data: dataPush
    })
}

export const getApi = (url = '', data = {}, type = 'json', isAuthen = true) => {
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

    if (isAuthen) {
        if (token == null) {
            return Promise.reject(new Error('need token'))
        } else {
            headers["Authorization"] = "Bearer " + token
        }
    }

    return axios({
        method: 'get',
        url: url,
        headers: headers,
        data: dataPush
    })
}
