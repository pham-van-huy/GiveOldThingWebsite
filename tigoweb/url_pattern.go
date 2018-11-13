// Copyright 2018 The Tigo Authors. All rights reserved.
package tigoweb

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

// URL路由中间件
type UrlPatternMidWare struct {
	Handler    interface{}
	requestUrl string
}

func (urlPatternMidWare UrlPatternMidWare) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	// 加载handler
	handler := reflect.ValueOf(urlPatternMidWare.Handler)
	// 获取init方法
	init := handler.MethodByName("InitHandler")

	paramPasser := handler.MethodByName("ParseRequestParam")

	middlewares := handler.MethodByName("RunBefore")

	requestMethod := MethodMapping[request.Method]
	Trace.Printf("%s %s", requestMethod, urlPatternMidWare.requestUrl)
	function := handler.MethodByName(requestMethod)
	initParams := []reflect.Value{reflect.ValueOf(responseWriter), reflect.ValueOf(request)}
	var functionParams []reflect.Value
	if init.IsValid() {
		init.Call(initParams)
	}
	if middlewares.IsValid() {
		middlewares.Call(functionParams)
	}
	if paramPasser.IsValid() {
		paramPasser.Call(functionParams)
	}
	if function.IsValid() {
		function.Call(functionParams)
	}
}

// URL路由，此处存储URL映射。
type UrlPattern struct {
}

// 向http服务挂载单个handler，注意：
//   - handler必须有一个Handle(http.ResponseWriter, *http.Request)函数
func (urlPattern *UrlPattern) AppendUrlPattern(router *mux.Router, uri string, v interface {
	Handle(http.ResponseWriter, *http.Request)
}) {
	router.HandleFunc(uri, v.Handle)
}

// 初始化url映射，遍历UrlMapping，将handler与对应的URL依次挂载到http服务上
func (urlPattern *UrlPattern) Init(router *mux.Router, urlMapping map[string]interface{}) {
	for key, value := range urlMapping {
		urlPatternMidWare := UrlPatternMidWare{
			Handler:    value,
			requestUrl: key,
		}
		urlPattern.AppendUrlPattern(router, key, &urlPatternMidWare)
	}
}
