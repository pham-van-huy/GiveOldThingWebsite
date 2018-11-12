package tigoweb

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Response interface {
	// Print()
}

// Handler的基础类，开发者开发的handler继承此类
type BaseHandler struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	RequestParams  map[string]interface{}
	ctxValMap      map[string]interface{}
}

// 初始化Handler的方法
func (baseHandler *BaseHandler) InitHandler(responseWriter http.ResponseWriter, request *http.Request) {
	baseHandler.Request = request
	baseHandler.ResponseWriter = responseWriter
}

// 解析json中的值
func (baseHandler *BaseHandler) ParseRequestParam() {
	defer baseHandler.Request.Body.Close()
	if baseHandler.GetHeader("Content-Type") == "application/json" {
		jsonData, _ := ioutil.ReadAll(baseHandler.Request.Body)
		json.Unmarshal([]byte(jsonData), &baseHandler.RequestParams)
	} else if baseHandler.GetHeader("Content-Type") == "multipart/form-data" {
		for key, vals := range baseHandler.Request.MultipartForm.Value {
			for _, val := range vals {
				baseHandler.RequestParams[key] = val
			}
		}
		for key, files := range baseHandler.Request.MultipartForm.File {
			if len(files) != 0 {
				baseHandler.RequestParams[key] = files
			}
		}
	} else if baseHandler.GetHeader("Content-Type") == "application/x-www-form-urlencoded" {
		baseHandler.Request.ParseForm()
		for key, vals := range baseHandler.Request.Form {
			for _, val := range vals {
				baseHandler.RequestParams[key] = val
			}
		}
	}
}

/////////////////////////////////////////////////////output/////////////////////////////////////////////////////////////
func (baseHandler *BaseHandler) ToJson(response Response) (result string) {
	// 将该对象转换为byte字节数组
	jsonResult, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		return ""
	}
	// 将byte数组转换为string
	return string(jsonResult)
}

// 向客户端响应一个Json结果
func (baseHandler *BaseHandler) ResponseAsJson(response Response) {
	// 将对象转换为Json字符串
	jsonResult := baseHandler.ToJson(response)
	// 设置http报文头内的Content-Type
	baseHandler.ResponseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(baseHandler.ResponseWriter, jsonResult)
}

// 向客户端响应一个Text结果
func (baseHandler *BaseHandler) ResponseAsText(result string) {
	fmt.Fprintf(baseHandler.ResponseWriter, result)
}

// 向客户端响应一个html结果
func (baseHandler *BaseHandler) ResponseAsHtml(result string) {
	baseHandler.ResponseWriter.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(baseHandler.ResponseWriter, result)
}

// 向客户端响应一个结果
func (baseHandler *BaseHandler) Response(result ...interface{}) {
	fmt.Fprintf(baseHandler.ResponseWriter, "%v", result)
}

// 渲染模板，返回数据
// 参数解析如下：
//   - data：表示传入的待渲染的数据
//   - templates：表示模板文件的路径，接受多个模板文件
func (baseHandler *BaseHandler) Render(data interface{}, templates ...string) {
	templateBasePath := globalConfig.Template
	var templatePath []string
	for _, value := range templates {
		value = templateBasePath + value
		templatePath = append(templatePath, value)
	}
	t, _ := template.ParseFiles(templatePath...)
	t.Execute(baseHandler.ResponseWriter, data)
}

// 向客户端永久重定向一个地址
func (baseHandler *BaseHandler) RedirectPermanently(url string) {
	baseHandler.ResponseWriter.WriteHeader(301)
	baseHandler.SetHeader("Location", url)
	fmt.Fprintf(baseHandler.ResponseWriter, "")
}

// 向客户端暂时重定向一个地址
func (baseHandler *BaseHandler) Redirect(url string, expire ...time.Time) {
	baseHandler.SetHeader("Location", url)
	baseHandler.ResponseWriter.WriteHeader(302)
	if len(expire) > 0 {
		expireTime := expire[0]
		expires := expireTime.Format("Mon, 02 Jan 2006 15:04:05 GMT")
		baseHandler.SetHeader("Expires", expires)
	}
	fmt.Fprintf(baseHandler.ResponseWriter, "")
}

/////////////////////////////////////////////////////cookie/////////////////////////////////////////////////////////////

// 设置cookie
func (baseHandler *BaseHandler) SetCookie(name string, value string) {
	cookie := http.Cookie{Name: name, Value: value}
	http.SetCookie(baseHandler.ResponseWriter, &cookie)
}

// 设置高级cookie选项
func (baseHandler *BaseHandler) SetCookieObject(cookie Cookie) {
	responseCookie := cookie.ToHttpCookie()
	http.SetCookie(baseHandler.ResponseWriter, &responseCookie)
}

// 设置加密cookie
func (baseHandler *BaseHandler) SetSecureCookie(name string, value string, key ...string) {
	securityKey := ""
	if len(key) > 0 {
		securityKey = key[0]
	} else {
		securityKey = globalConfig.Cookie
	}
	cookie := Cookie{
		Name:        name,
		Value:       value,
		IsSecurity:  true,
		SecurityKey: securityKey,
	}
	baseHandler.SetCookieObject(cookie)
}

// 获取cookie值，如果获取失败则返回空字符串
func (baseHandler *BaseHandler) GetCookie(name string) (value string) {
	cookie, err := baseHandler.Request.Cookie(name)
	if err != nil {
		return ""
	}
	value = cookie.Value
	return value
}

// 获取加密cookie值，如果获取失败则返回空
func (baseHandler *BaseHandler) GetSecureCookie(name string, key ...string) (value string) {
	securityKey := ""
	if len(key) > 0 {
		securityKey = key[0]
	} else {
		securityKey = globalConfig.Cookie
	}
	httpCookie, err := baseHandler.Request.Cookie(name)
	if err != nil {
		return ""
	}
	cookie := Cookie{}
	cookie.ConvertFromHttpCookie(*httpCookie)
	cookie.IsSecurity = true
	cookie.SecurityKey = securityKey
	value = cookie.GetCookieDecodeValue()
	return value
}

// 获取cookie对象，多参数输入，参数如下：
//   - 无参数：默认cookieName为空字符串
//   - 一个参数：传入的参数为cookieName
//   - 多个参数：传入的第一个参数为cookieName，第二个参数为加密/解密cookie所用的Key，此时认为cookie是需要进行加密/解密处理的
func (baseHandler *BaseHandler) GetCookieObject(name ...string) (Cookie, error) {
	cookie := Cookie{}
	var cookieName, key string
	length := len(name)
	switch {
	case length < 1:
		cookieName = ""
	case length == 1:
		cookieName = name[0]
	case length > 1:
		cookieName = name[0]
		key = name[1]
	}
	httpCookie, err := baseHandler.Request.Cookie(cookieName)
	if err != nil {
		return cookie, nil
	}
	cookie.ConvertFromHttpCookie(*httpCookie)
	if len(key) > 0 {
		cookie.SetSecurityKey(key)
	}
	return cookie, nil
}

// 清除当前path下的指定的cookie
func (baseHandler *BaseHandler) ClearCookie(name string) {
	cookie := Cookie{
		Name:    name,
		Expires: time.Now(),
	}
	baseHandler.SetCookieObject(cookie)
}

// 清除当前path下所有的cookie
func (baseHandler *BaseHandler) ClearAllCookie() {
	cookies := baseHandler.Request.Cookies()
	for _, cookie := range cookies {
		baseHandler.ClearCookie(cookie.Name)
	}
}

/////////////////////////////////////////////////////input//////////////////////////////////////////////////////////////

// 获取header
func (baseHandler *BaseHandler) GetHeader(name string) (value string) {
	value = baseHandler.Request.Header.Get(name)
	return value
}

// 设置header
func (baseHandler *BaseHandler) SetHeader(name string, value string) {
	baseHandler.ResponseWriter.Header().Set(name, value)
}

// 根据key获取对应的参数值
//   - 如果Content-Type是application/json，则直接从http的body中解析出key对应的value
//   - 否则，根据key直接获取value
func (baseHandler *BaseHandler) DecodeStruct(v interface{}) error {
	err := mapstructure.Decode(baseHandler.RequestParams, &v)
	return err
}

// 根据key获取对应的参数值
//   - 如果Content-Type是application/json，则直接从http的body中解析出key对应的value
//   - 否则，根据key直接获取value
func (baseHandler *BaseHandler) GetParameter(key string) (value *ReqParams) {
	jsonValue := &ReqParams{}
	if baseHandler.GetHeader("Content-Type") == "application/json" {
		if value, ok := baseHandler.RequestParams[key]; ok {
			jsonValue.Value = value
		} else {
			jsonValue.Value = nil
		}
		return jsonValue
	}
	jsonValue.Value = baseHandler.Request.FormValue(key)
	return jsonValue
}

// 根据key获取对应的参数值，解析json数据，返回对应的value
func (baseHandler *BaseHandler) GetJsonValue(key string) interface{} {
	var mapResult map[string]interface{}
	jsonData, _ := ioutil.ReadAll(baseHandler.Request.Body)
	baseHandler.Request.Body.Close()
	//使用 json.Unmarshal(data []byte, v interface{})进行转换，返回 error 信息
	err := json.Unmarshal([]byte(jsonData), &mapResult)
	if err != nil {
		return ""
	}
	return mapResult[key]
}

//////////////////////////////////////////////////HTTP Method///////////////////////////////////////////////////////////

// 请求方法不合法
func (baseHandler *BaseHandler) methodNotAllowed() {
	baseHandler.ResponseWriter.WriteHeader(405)
}

func (baseHandler *BaseHandler) Get() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Put() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Post() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Connect() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Trace() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Head() {
}

func (baseHandler *BaseHandler) Delete() {
	baseHandler.methodNotAllowed()
}

func (baseHandler *BaseHandler) Options() {
	baseHandler.methodNotAllowed()
}

//////////////////////////////////////////////////Context Method////////////////////////////////////////////////////////

func (baseHandler *BaseHandler) SetCtxVal(key string, val interface{}) {
	baseHandler.ctxValMap[key] = val
}

func (baseHandler *BaseHandler) GetCtxVal(key string) interface{} {
	if val, isExisted := baseHandler.ctxValMap[key]; isExisted {
		return val
	} else {
		return nil
	}
}

//////////////////////////////////////////////////http message dump/////////////////////////////////////////////////////

// 获取http请求报文
func (baseHandler *BaseHandler) getHttpRequestMsg() string {
	req, err := httputil.DumpRequest(baseHandler.Request, true)
	if err != nil {
		return err.Error()
	}
	return string(req)
}
