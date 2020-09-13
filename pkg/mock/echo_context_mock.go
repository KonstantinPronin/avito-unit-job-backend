// Code generated by MockGen. DO NOT EDIT.
// Source: ./application/user/mocks/context.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo"
	io "io"
	multipart "mime/multipart"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockContext)(nil).Request))
}

// SetRequest mocks base method
func (m *MockContext) SetRequest(r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequest", r)
}

// SetRequest indicates an expected call of SetRequest
func (mr *MockContextMockRecorder) SetRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequest", reflect.TypeOf((*MockContext)(nil).SetRequest), r)
}

// SetResponse mocks base method
func (m *MockContext) SetResponse(r *echo.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResponse", r)
}

// SetResponse indicates an expected call of SetResponse
func (mr *MockContextMockRecorder) SetResponse(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponse", reflect.TypeOf((*MockContext)(nil).SetResponse), r)
}

// Response mocks base method
func (m *MockContext) Response() *echo.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(*echo.Response)
	return ret0
}

// Response indicates an expected call of Response
func (mr *MockContextMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockContext)(nil).Response))
}

// IsTLS mocks base method
func (m *MockContext) IsTLS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTLS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTLS indicates an expected call of IsTLS
func (mr *MockContextMockRecorder) IsTLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTLS", reflect.TypeOf((*MockContext)(nil).IsTLS))
}

// IsWebSocket mocks base method
func (m *MockContext) IsWebSocket() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWebSocket")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWebSocket indicates an expected call of IsWebSocket
func (mr *MockContextMockRecorder) IsWebSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWebSocket", reflect.TypeOf((*MockContext)(nil).IsWebSocket))
}

// Scheme mocks base method
func (m *MockContext) Scheme() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(string)
	return ret0
}

// Scheme indicates an expected call of Scheme
func (mr *MockContextMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockContext)(nil).Scheme))
}

// RealIP mocks base method
func (m *MockContext) RealIP() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RealIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// RealIP indicates an expected call of RealIP
func (mr *MockContextMockRecorder) RealIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RealIP", reflect.TypeOf((*MockContext)(nil).RealIP))
}

// Path mocks base method
func (m *MockContext) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockContextMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockContext)(nil).Path))
}

// SetPath mocks base method
func (m *MockContext) SetPath(p string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPath", p)
}

// SetPath indicates an expected call of SetPath
func (mr *MockContextMockRecorder) SetPath(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPath", reflect.TypeOf((*MockContext)(nil).SetPath), p)
}

// Param mocks base method
func (m *MockContext) Param(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param
func (mr *MockContextMockRecorder) Param(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockContext)(nil).Param), name)
}

// ParamNames mocks base method
func (m *MockContext) ParamNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamNames indicates an expected call of ParamNames
func (mr *MockContextMockRecorder) ParamNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamNames", reflect.TypeOf((*MockContext)(nil).ParamNames))
}

// SetParamNames mocks base method
func (m *MockContext) SetParamNames(names ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range names {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetParamNames", varargs...)
}

// SetParamNames indicates an expected call of SetParamNames
func (mr *MockContextMockRecorder) SetParamNames(names ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParamNames", reflect.TypeOf((*MockContext)(nil).SetParamNames), names...)
}

// ParamValues mocks base method
func (m *MockContext) ParamValues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamValues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamValues indicates an expected call of ParamValues
func (mr *MockContextMockRecorder) ParamValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamValues", reflect.TypeOf((*MockContext)(nil).ParamValues))
}

// SetParamValues mocks base method
func (m *MockContext) SetParamValues(values ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetParamValues", varargs...)
}

// SetParamValues indicates an expected call of SetParamValues
func (mr *MockContextMockRecorder) SetParamValues(values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParamValues", reflect.TypeOf((*MockContext)(nil).SetParamValues), values...)
}

// QueryParam mocks base method
func (m *MockContext) QueryParam(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParam", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryParam indicates an expected call of QueryParam
func (mr *MockContextMockRecorder) QueryParam(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParam", reflect.TypeOf((*MockContext)(nil).QueryParam), name)
}

// QueryParams mocks base method
func (m *MockContext) QueryParams() url.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParams")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

// QueryParams indicates an expected call of QueryParams
func (mr *MockContextMockRecorder) QueryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParams", reflect.TypeOf((*MockContext)(nil).QueryParams))
}

// QueryString mocks base method
func (m *MockContext) QueryString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryString")
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryString indicates an expected call of QueryString
func (mr *MockContextMockRecorder) QueryString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryString", reflect.TypeOf((*MockContext)(nil).QueryString))
}

// FormValue mocks base method
func (m *MockContext) FormValue(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormValue", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// FormValue indicates an expected call of FormValue
func (mr *MockContextMockRecorder) FormValue(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormValue", reflect.TypeOf((*MockContext)(nil).FormValue), name)
}

// FormParams mocks base method
func (m *MockContext) FormParams() (url.Values, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormParams")
	ret0, _ := ret[0].(url.Values)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormParams indicates an expected call of FormParams
func (mr *MockContextMockRecorder) FormParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormParams", reflect.TypeOf((*MockContext)(nil).FormParams))
}

// FormFile mocks base method
func (m *MockContext) FormFile(name string) (*multipart.FileHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", name)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile
func (mr *MockContextMockRecorder) FormFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockContext)(nil).FormFile), name)
}

// MultipartForm mocks base method
func (m *MockContext) MultipartForm() (*multipart.Form, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipartForm")
	ret0, _ := ret[0].(*multipart.Form)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultipartForm indicates an expected call of MultipartForm
func (mr *MockContextMockRecorder) MultipartForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipartForm", reflect.TypeOf((*MockContext)(nil).MultipartForm))
}

// Cookie mocks base method
func (m *MockContext) Cookie(name string) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookie", name)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie
func (mr *MockContextMockRecorder) Cookie(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockContext)(nil).Cookie), name)
}

// SetCookie mocks base method
func (m *MockContext) SetCookie(cookie *http.Cookie) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCookie", cookie)
}

// SetCookie indicates an expected call of SetCookie
func (mr *MockContextMockRecorder) SetCookie(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockContext)(nil).SetCookie), cookie)
}

// Cookies mocks base method
func (m *MockContext) Cookies() []*http.Cookie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookies")
	ret0, _ := ret[0].([]*http.Cookie)
	return ret0
}

// Cookies indicates an expected call of Cookies
func (mr *MockContextMockRecorder) Cookies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookies", reflect.TypeOf((*MockContext)(nil).Cookies))
}

// Get mocks base method
func (m *MockContext) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContext)(nil).Get), key)
}

// Set mocks base method
func (m *MockContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set
func (mr *MockContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContext)(nil).Set), key, val)
}

// Bind mocks base method
func (m *MockContext) Bind(i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind
func (mr *MockContextMockRecorder) Bind(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockContext)(nil).Bind), i)
}

// Validate mocks base method
func (m *MockContext) Validate(i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockContextMockRecorder) Validate(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockContext)(nil).Validate), i)
}

// Render mocks base method
func (m *MockContext) Render(code int, name string, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", code, name, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Render indicates an expected call of Render
func (mr *MockContextMockRecorder) Render(code, name, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockContext)(nil).Render), code, name, data)
}

// HTML mocks base method
func (m *MockContext) HTML(code int, html string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTML", code, html)
	ret0, _ := ret[0].(error)
	return ret0
}

// HTML indicates an expected call of HTML
func (mr *MockContextMockRecorder) HTML(code, html interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTML", reflect.TypeOf((*MockContext)(nil).HTML), code, html)
}

// HTMLBlob mocks base method
func (m *MockContext) HTMLBlob(code int, b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTMLBlob", code, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// HTMLBlob indicates an expected call of HTMLBlob
func (mr *MockContextMockRecorder) HTMLBlob(code, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTMLBlob", reflect.TypeOf((*MockContext)(nil).HTMLBlob), code, b)
}

// String mocks base method
func (m *MockContext) String(code int, s string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", code, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// String indicates an expected call of String
func (mr *MockContextMockRecorder) String(code, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockContext)(nil).String), code, s)
}

// JSON mocks base method
func (m *MockContext) JSON(code int, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSON", code, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSON indicates an expected call of JSON
func (mr *MockContextMockRecorder) JSON(code, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockContext)(nil).JSON), code, i)
}

// JSONPretty mocks base method
func (m *MockContext) JSONPretty(code int, i interface{}, indent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONPretty", code, i, indent)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONPretty indicates an expected call of JSONPretty
func (mr *MockContextMockRecorder) JSONPretty(code, i, indent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONPretty", reflect.TypeOf((*MockContext)(nil).JSONPretty), code, i, indent)
}

// JSONBlob mocks base method
func (m *MockContext) JSONBlob(code int, b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONBlob", code, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONBlob indicates an expected call of JSONBlob
func (mr *MockContextMockRecorder) JSONBlob(code, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONBlob", reflect.TypeOf((*MockContext)(nil).JSONBlob), code, b)
}

// JSONP mocks base method
func (m *MockContext) JSONP(code int, callback string, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONP", code, callback, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONP indicates an expected call of JSONP
func (mr *MockContextMockRecorder) JSONP(code, callback, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONP", reflect.TypeOf((*MockContext)(nil).JSONP), code, callback, i)
}

// JSONPBlob mocks base method
func (m *MockContext) JSONPBlob(code int, callback string, b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONPBlob", code, callback, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONPBlob indicates an expected call of JSONPBlob
func (mr *MockContextMockRecorder) JSONPBlob(code, callback, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONPBlob", reflect.TypeOf((*MockContext)(nil).JSONPBlob), code, callback, b)
}

// XML mocks base method
func (m *MockContext) XML(code int, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XML", code, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// XML indicates an expected call of XML
func (mr *MockContextMockRecorder) XML(code, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XML", reflect.TypeOf((*MockContext)(nil).XML), code, i)
}

// XMLPretty mocks base method
func (m *MockContext) XMLPretty(code int, i interface{}, indent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLPretty", code, i, indent)
	ret0, _ := ret[0].(error)
	return ret0
}

// XMLPretty indicates an expected call of XMLPretty
func (mr *MockContextMockRecorder) XMLPretty(code, i, indent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLPretty", reflect.TypeOf((*MockContext)(nil).XMLPretty), code, i, indent)
}

// XMLBlob mocks base method
func (m *MockContext) XMLBlob(code int, b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLBlob", code, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// XMLBlob indicates an expected call of XMLBlob
func (mr *MockContextMockRecorder) XMLBlob(code, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLBlob", reflect.TypeOf((*MockContext)(nil).XMLBlob), code, b)
}

// Blob mocks base method
func (m *MockContext) Blob(code int, contentType string, b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Blob", code, contentType, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Blob indicates an expected call of Blob
func (mr *MockContextMockRecorder) Blob(code, contentType, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blob", reflect.TypeOf((*MockContext)(nil).Blob), code, contentType, b)
}

// Stream mocks base method
func (m *MockContext) Stream(code int, contentType string, r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", code, contentType, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream
func (mr *MockContextMockRecorder) Stream(code, contentType, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockContext)(nil).Stream), code, contentType, r)
}

// File mocks base method
func (m *MockContext) File(file string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "File", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// File indicates an expected call of File
func (mr *MockContextMockRecorder) File(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockContext)(nil).File), file)
}

// Attachment mocks base method
func (m *MockContext) Attachment(file, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attachment", file, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Attachment indicates an expected call of Attachment
func (mr *MockContextMockRecorder) Attachment(file, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attachment", reflect.TypeOf((*MockContext)(nil).Attachment), file, name)
}

// Inline mocks base method
func (m *MockContext) Inline(file, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inline", file, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Inline indicates an expected call of Inline
func (mr *MockContextMockRecorder) Inline(file, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inline", reflect.TypeOf((*MockContext)(nil).Inline), file, name)
}

// NoContent mocks base method
func (m *MockContext) NoContent(code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContent", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// NoContent indicates an expected call of NoContent
func (mr *MockContextMockRecorder) NoContent(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContent", reflect.TypeOf((*MockContext)(nil).NoContent), code)
}

// Redirect mocks base method
func (m *MockContext) Redirect(code int, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redirect", code, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// Redirect indicates an expected call of Redirect
func (mr *MockContextMockRecorder) Redirect(code, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockContext)(nil).Redirect), code, url)
}

// Error mocks base method
func (m *MockContext) Error(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", err)
}

// Error indicates an expected call of Error
func (mr *MockContextMockRecorder) Error(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContext)(nil).Error), err)
}

// Handler mocks base method
func (m *MockContext) Handler() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// Handler indicates an expected call of Handler
func (mr *MockContextMockRecorder) Handler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockContext)(nil).Handler))
}

// SetHandler mocks base method
func (m *MockContext) SetHandler(h echo.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandler", h)
}

// SetHandler indicates an expected call of SetHandler
func (mr *MockContextMockRecorder) SetHandler(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandler", reflect.TypeOf((*MockContext)(nil).SetHandler), h)
}

// Logger mocks base method
func (m *MockContext) Logger() echo.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(echo.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockContextMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockContext)(nil).Logger))
}

// SetLogger mocks base method
func (m *MockContext) SetLogger(l echo.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", l)
}

// SetLogger indicates an expected call of SetLogger
func (mr *MockContextMockRecorder) SetLogger(l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockContext)(nil).SetLogger), l)
}

// Echo mocks base method
func (m *MockContext) Echo() *echo.Echo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo")
	ret0, _ := ret[0].(*echo.Echo)
	return ret0
}

// Echo indicates an expected call of Echo
func (mr *MockContextMockRecorder) Echo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockContext)(nil).Echo))
}

// Reset mocks base method
func (m *MockContext) Reset(r *http.Request, w http.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", r, w)
}

// Reset indicates an expected call of Reset
func (mr *MockContextMockRecorder) Reset(r, w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockContext)(nil).Reset), r, w)
}
