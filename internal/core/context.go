package core

import (
	stdctx "context"
	"fmt"
	logs "github.com/danbai225/go-logs"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"sync"
	"web_demo/internal/data_obj/base"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/pkg/stringutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HandlerFunc func()

const (
	_BodyName        = "_body_"
	_PayloadName     = "_payload_"
	_SessionUserInfo = "_session_user_info"
	_AbortError      = "_abort_error_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(context)
	},
}
var contextMap = &sync.Map{}

func newContext(ctx *gin.Context, db mysql.Repo, rdb redis.Repo) Context {
	c := contextPool.Get().(*context)
	c.ctx = ctx
	c.db = db
	c.rdb = rdb
	c.traceID = stringutil.UUID()
	return c
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}
func GetContext() Context {
	value, ok := contextMap.Load(getGid())
	if ok {
		return value.(Context)
	}
	return nil
}
func setContext(c Context) {
	contextMap.Store(getGid(), c)
}
func delContext() {
	contextMap.Delete(getGid())
}

var _ Context = (*context)(nil)

type Context interface {

	// ShouldBindQuery 反序列化 querystring
	// tag: `form:"xxx"` (注：不要写成 query)
	ShouldBindQuery(obj interface{}) error

	// ShouldBindPostForm 反序列化 postform (querystring会被忽略)
	// tag: `form:"xxx"`
	ShouldBindPostForm(obj interface{}) error

	// ShouldBindForm 同时反序列化 querystring 和 postform;
	// 当 querystring 和 postform 存在相同字段时，postform 优先使用。
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error

	// ShouldBindJSON 反序列化 postjson
	// tag: `json:"xxx"`
	ShouldBindJSON(obj interface{}) error

	// ShouldBindURI 反序列化 path 参数(如路由路径为 /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error

	// Redirect 重定向
	Redirect(code int, location string)

	// PayloadCodeMsg 正确返回
	PayloadCodeMsg(data interface{}, code int64, msgStr string)
	// Payload 正确返回
	Payload(payload interface{})
	getPayload() interface{}
	// AbortError 错误返回
	AbortError(err error, code ...int64)
	getAbortError() interface{}
	// HTML 返回界面
	HTML(name string, obj interface{})

	// Header 获取 Header 对象
	Header() http.Header
	// GetHeader 获取 Header
	GetHeader(key string) string
	// SetHeader 设置 Header
	SetHeader(key, value string)

	// SessionUserInfo 当前用户信息
	SessionUserInfo() *SessionUserInfo
	SetSessionUserInfo(info *SessionUserInfo)

	// RequestInputParams 获取所有参数
	RequestInputParams() url.Values
	// RequestPostFormParams  获取 PostForm 参数
	RequestPostFormParams() url.Values
	// Request 获取 Request 对象
	Request() *http.Request
	// RawData 获取 Request.Body
	RawData() []byte
	// Method 获取 Request.Method
	Method() string
	// Host 获取 Request.Host
	Host() string
	// Path 获取 请求的路径 Request.URL.Path (不附带 querystring)
	Path() string
	// URI 获取 unescape 后的 Request.URL.RequestURI()
	URI() string
	// RequestContext 获取请求的 context (当 client 关闭后，会自动 canceled)
	RequestContext() StdContext

	// ResponseWriter 获取 ResponseWriter 对象
	ResponseWriter() gin.ResponseWriter
	// GetVal 本次请求缓存
	GetVal(key string) interface{}
	SetVal(key string, val interface{})
	GetGin() *gin.Context
	//GetDb 获取mysqldb
	GetDb() *gorm.DB
	//GetRdb 获取redis db
	GetRdb() redis.Repo
	// Info 日志
	Info(any ...any)
	// Err 日志
	Err(any ...any)
	// Warn 日志
	Warn(any ...any)
}

type context struct {
	ctx     *gin.Context
	db      mysql.Repo
	rdb     redis.Repo
	traceID string
}

type StdContext struct {
	stdctx.Context
}

// ShouldBindQuery 反序列化querystring
// tag: `form:"xxx"` (注：不要写成query)
func (c *context) ShouldBindQuery(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Query)
}

// ShouldBindPostForm 反序列化 postform (querystring 会被忽略)
// tag: `form:"xxx"`
func (c *context) ShouldBindPostForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.FormPost)
}

// ShouldBindForm 同时反序列化querystring和postform;
// 当querystring和postform存在相同字段时，postform优先使用。
// tag: `form:"xxx"`
func (c *context) ShouldBindForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Form)
}

// ShouldBindJSON 反序列化postjson
// tag: `json:"xxx"`
func (c *context) ShouldBindJSON(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.JSON)
}

// ShouldBindURI 反序列化path参数(如路由路径为 /user/:name)
// tag: `uri:"xxx"`
func (c *context) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

// Redirect 重定向
func (c *context) Redirect(code int, location string) {
	c.ctx.Redirect(code, location)
}

func (c *context) getPayload() interface{} {
	if payload, ok := c.ctx.Get(_PayloadName); ok != false {
		return payload
	}
	return nil
}
func (c *context) PayloadCodeMsg(data interface{}, code int64, msgStr string) {
	msg := base.ReturnMsg{
		Code:  code,
		Msg:   msgStr,
		Data:  data,
		Trace: c.traceID,
	}
	c.ctx.Set(_PayloadName, msg)
}
func (c *context) Payload(data interface{}) {
	msg := base.ReturnMsg{
		Code:  0,
		Msg:   "ok",
		Data:  data,
		Trace: c.traceID,
	}
	c.ctx.Set(_PayloadName, msg)
}

// AbortError 错误返回
func (c *context) AbortError(err error, code ...int64) {
	msg := base.ReturnMsg{
		Code:  1,
		Msg:   err.Error(),
		Data:  nil,
		Trace: c.traceID,
	}
	if len(code) > 0 && code[0] != 0 {
		msg.Code = code[0]
	}
	c.ctx.Set(_AbortError, msg)
	c.GetGin().Abort()
}
func (c *context) getAbortError() interface{} {
	if err, ok := c.ctx.Get(_AbortError); ok != false {
		return err
	}
	return nil
}

func (c *context) HTML(name string, obj interface{}) {
	c.ctx.HTML(200, name+".html", obj)
}

func (c *context) Header() http.Header {
	header := c.ctx.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
}

func (c *context) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

func (c *context) SetHeader(key, value string) {
	c.ctx.Header(key, value)
}

func (c *context) SessionUserInfo() *SessionUserInfo {
	if c.ctx == nil {
		return nil
	}
	val, ok := c.ctx.Get(_SessionUserInfo)
	if !ok {
		return &SessionUserInfo{}
	}

	return val.(*SessionUserInfo)
}

func (c *context) SetSessionUserInfo(info *SessionUserInfo) {
	c.ctx.Set(_SessionUserInfo, info)
}

// RequestInputParams 获取所有参数
func (c *context) RequestInputParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.Form
}

// RequestPostFormParams 获取 PostForm 参数
func (c *context) RequestPostFormParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.PostForm
}

// Request 获取 Request
func (c *context) Request() *http.Request {
	return c.ctx.Request
}

func (c *context) RawData() []byte {
	body, ok := c.ctx.Get(_BodyName)
	if !ok {
		return nil
	}
	return body.([]byte)
}

// Method 请求的method
func (c *context) Method() string {
	return c.ctx.Request.Method
}

// Host 请求的host
func (c *context) Host() string {
	return c.ctx.Request.Host
}

// Path 请求的路径(不附带querystring)
func (c *context) Path() string {
	return c.ctx.Request.URL.Path
}

// URI unescape后的uri
func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.ctx.Request.URL.RequestURI())
	return uri
}

// RequestContext (包装 Trace + Logger) 获取请求的 context (当client关闭后，会自动canceled)
func (c *context) RequestContext() StdContext {
	return StdContext{
		stdctx.Background(),
	}
}

// ResponseWriter 获取 ResponseWriter
func (c *context) ResponseWriter() gin.ResponseWriter {
	return c.ctx.Writer
}

// GetVal 获取 缓存对象
func (c *context) GetVal(key string) interface{} {
	get, _ := c.ctx.Get(key)
	return get
}

// SetVal 设置缓存对象
func (c *context) SetVal(key string, val interface{}) {
	c.ctx.Set(key, val)
}

// GetGin 获取gin
func (c *context) GetGin() *gin.Context {
	return c.ctx
}

func (c *context) GetDb() *gorm.DB {
	return c.db.GetDb()
}
func (c *context) GetRdb() redis.Repo {
	return c.rdb
}
func (c *context) Info(any ...any) {
	any = append([]interface{}{fmt.Sprintf("[%s]", c.traceID)}, any...)
	logs.InfoN(0, any...)
}
func (c *context) Err(any ...any) {
	any = append([]interface{}{fmt.Sprintf("[%s]", c.traceID)}, any...)
	logs.ErrN(0, any...)
}
func (c *context) Warn(any ...any) {
	any = append([]interface{}{fmt.Sprintf("[%s]", c.traceID)}, any...)
	logs.WarnN(0, any...)
}
