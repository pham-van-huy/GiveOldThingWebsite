package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type payload interface {
}

type response struct {
	Success bool    `json:"success"`
	Data    payload `json:"data"`
}

func ResSuccess(w http.ResponseWriter, data payload) {
	res := response{}

	res.Success = true
	res.Data = data

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func ResErrors(w http.ResponseWriter, data payload) {
	res := response{}

	res.Success = false
	res.Data = data

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}
