package tigoweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//////////////////////////////////////Structure Cookie//////////////////////////////////////////////////////////////////

// 自定义Cookie结构体，可参看http.Cookie
type Cookie struct {
	Name  string
	Value string

	IsSecurity  bool   // 是否对cookie值进行加密
	SecurityKey string // 加密cookie用到的key

	Path       string    // 可选
	Domain     string    // 可选
	Expires    time.Time // 可选
	RawExpires string    // 只有在读取Cookie时有效

	// MaxAge=0 表示未指定“Max-Age”属性
	// MaxAge<0 表示现在删除cookie，相当于'Max-Age：0'
	// MaxAge>0 表示Max-Age属性存在并以秒为单位给出
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // 原始文本中未解析的属性值
}

// 获取cookie加密值
//   - IsSecurity如果设置为false，则返回原始值
//   - IsSecurity如果设置为true，则返回加密后的值
// 如果加密失败，则抛出异常
func (cookie *Cookie) GetCookieEncodeValue() (result string) {
	if !cookie.IsSecurity {
		return cookie.Value
	}
	value := []byte(cookie.Value)
	key := []byte(cookie.SecurityKey)
	result = Encrypt(value, key)
	return result
}

// 获取cookie解密值
//   - IsSecurity如果设置为false，则返回原始值
//   - IsSecurity如果设置为true，则返回加密后的值
// 如果解密失败，则抛出异常
func (cookie *Cookie) GetCookieDecodeValue() (result string) {
	if !cookie.IsSecurity {
		return cookie.Value
	}
	value := []byte(cookie.Value)
	key := []byte(cookie.SecurityKey)
	securityValue := Decrypt(value, key)
	result = string(securityValue)
	return result
}

// 转换为http/Cookie对象
func (cookie *Cookie) ToHttpCookie() http.Cookie {
	httpCookie := http.Cookie{
		Name:       cookie.Name,
		Value:      cookie.GetCookieEncodeValue(),
		Path:       cookie.Path,
		Domain:     cookie.Domain,
		Expires:    cookie.Expires,
		RawExpires: cookie.RawExpires,
		MaxAge:     cookie.MaxAge,
		Secure:     cookie.Secure,
		HttpOnly:   cookie.HttpOnly,
		Raw:        cookie.Raw,
		Unparsed:   cookie.Unparsed,
	}
	return httpCookie
}

// 将http/Cookie转换为Cookie
func (cookie *Cookie) ConvertFromHttpCookie(httpCookie http.Cookie) {
	cookie.Name = httpCookie.Name
	cookie.Value = httpCookie.Value

	cookie.Path = httpCookie.Path
	cookie.Domain = httpCookie.Domain
	cookie.Expires = httpCookie.Expires
	cookie.RawExpires = httpCookie.RawExpires

	cookie.MaxAge = httpCookie.MaxAge
	cookie.Secure = httpCookie.Secure
	cookie.HttpOnly = httpCookie.HttpOnly
	cookie.Raw = httpCookie.Raw
	cookie.Unparsed = httpCookie.Unparsed
}

// 为Cookie设置SecurityKey
func (cookie *Cookie) SetSecurityKey(key string) {
	cookie.SecurityKey = key
	cookie.IsSecurity = true
}

//////////////////////////////////////Structure BaseResponse////////////////////////////////////////////////////////////

// 定义BaseResponse类，其他Json数据类继承此类，用于BaseHandler.ResponseAsJson的参数。
type BaseResponse struct {
}

// 打印Json数据
func (baseResponse *BaseResponse) Print() {
	fmt.Println(baseResponse.ToJson())
}

// 序列化为Json字符串
func (baseResponse *BaseResponse) ToJson() string {
	// 将该对象转换为byte字节数组
	jsonResult, jsonErr := json.Marshal(baseResponse)
	if jsonErr != nil {
		return "To Json Failed!"
	}
	// 将byte数组转换为string
	return string(jsonResult)
}

//////////////////////////////////////Structure GlobalConfig////////////////////////////////////////////////////////////

// 全局配置对象
type GlobalConfig struct {
	IP       string   `json:"ip"`       // IP地址
	Port     int      `json:"port"`     // 端口
	Cert     string   `json:"cert"`     // https证书路径
	CertKey  string   `json:"cert_key"` // https密钥路径
	Cookie   string   `json:"cookie"`   // cookie加密解密的密钥
	Template string   `json:"template"` // 模板文件所在文件夹的路径
	Log      LogLevel `json:"log"`      // log相关属性配置
}

// 根据配置文件初始化全局配置变量
func (globalConfig *GlobalConfig) Init(configPath string) {
	if configPath == "" {
		return
	}
	if strings.HasSuffix(configPath, ".json") {
		globalConfig.initWithJson(configPath)
	}
}

// 根据json文件进行配置
func (globalConfig *GlobalConfig) initWithJson(configPath string) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	jsonErr := json.Unmarshal(raw, &globalConfig)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
		os.Exit(1)
	}
	InitLoggerWithObject(globalConfig.Log)
}

//////////////////////////////////////Json Param Structure /////////////////////////////////////////////////////////////

type ReqParams struct {
	Value interface{}
}

func ConvertParam(v interface{}) *ReqParams {
	t := &ReqParams{}
	t.Value = v
	return t
}

// 将json中解析出的参数格式化为string类型，失败则返回空字符串
func (jsonParam *ReqParams) ToString() string {
	valueType := reflect.TypeOf(jsonParam.Value).Name()
	switch valueType {
	case "string":
		return jsonParam.Value.(string)
	case "int":
		return strconv.Itoa(jsonParam.Value.(int))
	case "int64":
		return strconv.FormatInt(jsonParam.Value.(int64), 10)
	case "float64":
		return strconv.FormatFloat(jsonParam.Value.(float64), 'E', -1, 32)
	case "bool":
		val := jsonParam.Value.(bool)
		if val {
			return "true"
		}
		return "false"
	}
	Warning.Println("format to string failed")
	return ""
}

// 将json中的参数值转换为bool
func (jsonParam *ReqParams) ToBool(defaultValue ...bool) bool {
	if jsonParam.Value == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	valueType := reflect.TypeOf(jsonParam.Value).Name()
	switch valueType {
	case "string":
		val := jsonParam.Value.(string)
		if strings.ToLower(val) == "true" {
			return true
		}
		return false
	case "int":
		val := jsonParam.Value.(int)
		if val > 0 {
			return true
		}
		return false
	case "int64":
		val := jsonParam.Value.(int64)
		if val > 0 {
			return true
		}
		return false
	case "float64":
		val := jsonParam.Value.(float64)
		if val > 0 {
			return true
		}
		return false
	}
	result, success := jsonParam.Value.(bool)
	if !success {
		return false
	}
	return result
}

// 将json中解析出来的参数格式化为int64类型，失败则返回0
func (jsonParam *ReqParams) ToInt64() int64 {
	valueType := reflect.TypeOf(jsonParam.Value).Name()
	switch valueType {
	case "string":
		val, err := strconv.ParseInt(jsonParam.Value.(string), 10, 64)
		if err != nil {
			Warning.Println("parse value to int64 failed")
			return 0
		}
		return val
	case "float64":
		return int64(jsonParam.Value.(float64))
	case "bool":
		val := jsonParam.Value.(bool)
		if val {
			return 1
		}
		return 0
	}
	result, success := jsonParam.Value.(int64)
	if !success {
		Warning.Println("convert value to int64 failed")
		return 0
	}
	return result
}

func (jsonParam *ReqParams) ToInt() int {
	valueType := reflect.TypeOf(jsonParam.Value).Name()
	switch valueType {
	case "string":
		val, err := strconv.ParseInt(jsonParam.Value.(string), 10, 0)
		if err != nil {
			Warning.Println("parse value to int64 failed")
			return 0
		}
		return int(val)
	case "float64":
		return int(jsonParam.Value.(float64))
	case "bool":
		val := jsonParam.Value.(bool)
		if val {
			return 1
		}
		return 0
	}
	result, success := jsonParam.Value.(int)
	if !success {
		Warning.Println("convert value to int64 failed")
		return 0
	}
	return result
}

// 将json中解析出的参数格式化为float64类型，失败则返回0
func (jsonParam *ReqParams) ToFloat64() float64 {
	valueType := reflect.TypeOf(jsonParam.Value).Name()
	switch valueType {
	case "string":
		val, err := strconv.ParseFloat(jsonParam.Value.(string), 64)
		if err != nil {
			Warning.Println("parse value to float64 failed")
			return 0
		}
		return val
	case "bool":
		val := jsonParam.Value.(bool)
		if val {
			return 1
		}
		return 0
	}
	result, success := jsonParam.Value.(float64)
	if !success {
		Warning.Println("convert value to float64 failed")
		return 0
	}
	return result
}

// 将json中的参数值转换为目标对象
func (jsonParam *ReqParams) To(result interface{}) {
	if jsonParam.Value == nil {
		return
	}
	if jsonData, err := json.Marshal(jsonParam.Value); err != nil {
		return
	} else {
		json.Unmarshal(jsonData, &result)
	}
}
