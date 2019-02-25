package echo

import (
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type RecordContext struct {
	response *echo.Response
	context  echo.Context
}

// NewResponse creates a new instance of Response.
func NewRecordContext(context echo.Context, w http.ResponseWriter, e *echo.Echo) (r *RecordContext) {
	return &RecordContext{response: echo.NewResponse(w, e), context: context}
}

// Request returns `*http.Request`.
func (c *RecordContext) Request() *http.Request {
	return c.context.Request()
}

// SetRequest sets `*http.Request`.
func (c *RecordContext) SetRequest(r *http.Request) {
	c.context.SetRequest(r)
}

// Response returns `*Response`.
func (c *RecordContext) Response() *echo.Response {
	return c.response
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (c *RecordContext) IsTLS() bool {
	return c.context.IsTLS()
}

// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
func (c *RecordContext) IsWebSocket() bool {
	return c.context.IsWebSocket()
}

// Scheme returns the HTTP protocol scheme, `http` or `https`.
func (c *RecordContext) Scheme() string {
	return c.context.Scheme()
}

// RealIP returns the client's network address based on `X-Forwarded-For`
// or `X-Real-IP` request header.
func (c *RecordContext) RealIP() string {
	return c.context.RealIP()
}

// Path returns the registered path for the handler.
func (c *RecordContext) Path() string {
	return c.context.Path()
}

// SetPath sets the registered path for the handler.
func (c *RecordContext) SetPath(p string) {
	c.context.SetPath(p)
}

// Param returns path parameter by name.
func (c *RecordContext) Param(name string) string {
	return c.context.Param(name)
}

// ParamNames returns path parameter names.
func (c *RecordContext) ParamNames() []string {
	return c.context.ParamNames()
}

// SetParamNames sets path parameter names.
func (c *RecordContext) SetParamNames(names ...string) {
	c.context.SetParamNames(names...)
}

// ParamValues returns path parameter values.
func (c *RecordContext) ParamValues() []string {
	return c.context.ParamValues()
}

// SetParamValues sets path parameter values.
func (c *RecordContext) SetParamValues(values ...string) {
	c.context.SetParamNames(values...)
}

// QueryParam returns the query param for the provided name.
func (c *RecordContext) QueryParam(name string) string {
	return c.context.QueryParam(name)
}

// QueryParams returns the query parameters as `url.Values`.
func (c *RecordContext) QueryParams() url.Values {
	return c.context.QueryParams()
}

// QueryString returns the URL query string.
func (c *RecordContext) QueryString() string {
	return c.context.QueryString()
}

// FormValue returns the form field value for the provided name.
func (c *RecordContext) FormValue(name string) string {
	return c.context.FormValue(name)
}

// FormParams returns the form parameters as `url.Values`.
func (c *RecordContext) FormParams() (url.Values, error) {
	return c.context.FormParams()
}

// FormFile returns the multipart form file for the provided name.
func (c *RecordContext) FormFile(name string) (*multipart.FileHeader, error) {
	return c.context.FormFile(name)
}

// MultipartForm returns the multipart form.
func (c *RecordContext) MultipartForm() (*multipart.Form, error) {
	return c.context.MultipartForm()
}

// Cookie returns the named cookie provided in the request.
func (c *RecordContext) Cookie(name string) (*http.Cookie, error) {
	return c.context.Cookie(name)
}

// SetCookie adds a `Set-Cookie` header in HTTP response.
func (c *RecordContext) SetCookie(cookie *http.Cookie) {
	c.context.SetCookie(cookie)
}

// Cookies returns the HTTP cookies sent with the request.
func (c *RecordContext) Cookies() []*http.Cookie {
	return c.context.Cookies()
}

// Get retrieves data from the context.
func (c *RecordContext) Get(key string) interface{} {
	return c.context.Get(key)
}

// Set saves data in the context.
func (c *RecordContext) Set(key string, val interface{}) {
	c.context.Set(key, val)
}

// Bind binds the request body into provided type `i`. The default binder
// does it based on Content-Type header.
func (c *RecordContext) Bind(i interface{}) error {
	return c.context.Bind(i)
}

// Validate validates provided `i`. It is usually called after `Context#Bind()`.
// Validator must be registered using `Echo#Validator`.
func (c *RecordContext) Validate(i interface{}) error {
	return c.context.Validate(i)
}

// Render renders a template with data and sends a text/html response with status
// code. Renderer must be registered using `Echo.Renderer`.
func (c *RecordContext) Render(code int, name string, data interface{}) error {
	return c.context.Render(code, name, data)
}

// HTML sends an HTTP response with status code.
func (c *RecordContext) HTML(code int, html string) error {
	return c.context.HTML(code, html)
}

// HTMLBlob sends an HTTP blob response with status code.
func (c *RecordContext) HTMLBlob(code int, b []byte) error {
	return c.context.HTMLBlob(code, b)
}

// String sends a string response with status code.
func (c *RecordContext) String(code int, s string) error {
	return c.context.String(code, s)
}

// JSON sends a JSON response with status code.
func (c *RecordContext) JSON(code int, i interface{}) error {
	return c.context.JSON(code, i)
}

// JSONPretty sends a pretty-print JSON with status code.
func (c *RecordContext) JSONPretty(code int, i interface{}, indent string) error {
	return c.context.JSONPretty(code, i, indent)
}

// JSONBlob sends a JSON blob response with status code.
func (c *RecordContext) JSONBlob(code int, b []byte) error {
	return c.context.JSONBlob(code, b)
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func (c *RecordContext) JSONP(code int, callback string, i interface{}) error {
	return c.context.JSONP(code, callback, i)
}

// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
// to construct the JSONP payload.
func (c *RecordContext) JSONPBlob(code int, callback string, b []byte) error {
	return c.context.JSONPBlob(code, callback, b)
}

// XML sends an XML response with status code.
func (c *RecordContext) XML(code int, i interface{}) error {
	return c.context.XML(code, i)
}

// XMLPretty sends a pretty-print XML with status code.
func (c *RecordContext) XMLPretty(code int, i interface{}, indent string) error {
	return c.context.XMLPretty(code, i, indent)
}

// XMLBlob sends an XML blob response with status code.
func (c *RecordContext) XMLBlob(code int, b []byte) error {
	return c.context.XMLBlob(code, b)
}

// Blob sends a blob response with status code and content type.
func (c *RecordContext) Blob(code int, contentType string, b []byte) error {
	return c.context.Blob(code, contentType, b)
}

// Stream sends a streaming response with status code and content type.
func (c *RecordContext) Stream(code int, contentType string, r io.Reader) error {
	return c.context.Stream(code, contentType, r)
}

// File sends a response with the content of the file.
func (c *RecordContext) File(file string) error {
	return c.context.File(file)
}

// Attachment sends a response as attachment, prompting client to save the
// file.
func (c *RecordContext) Attachment(file string, name string) error {
	return c.context.Attachment(file, name)
}

// Inline sends a response as inline, opening the file in the browser.
func (c *RecordContext) Inline(file string, name string) error {
	return c.context.Inline(file, name)
}

// NoContent sends a response with no body and a status code.
func (c *RecordContext) NoContent(code int) error {
	return c.context.NoContent(code)
}

// Redirect redirects the request to a provided URL with status code.
func (c *RecordContext) Redirect(code int, url string) error {
	return c.context.Redirect(code, url)
}

// Error invokes the registered HTTP error handler. Generally used by middleware.
func (c *RecordContext) Error(err error) {
	c.context.Error(err)
}

// Handler returns the matched handler by router.
func (c *RecordContext) Handler() echo.HandlerFunc {
	return c.context.Handler()
}

// SetHandler sets the matched handler by router.
func (c *RecordContext) SetHandler(h echo.HandlerFunc) {
	c.context.SetHandler(h)
}

// Logger returns the `Logger` instance.
func (c *RecordContext) Logger() echo.Logger {
	return c.context.Logger()
}

// Echo returns the `Echo` instance.
func (c *RecordContext) Echo() *echo.Echo {
	return c.context.Echo()
}

// Reset resets the context after request completes. It must be called along
// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
// See `Echo#ServeHTTP()`
func (c *RecordContext) Reset(r *http.Request, w http.ResponseWriter) {
	c.context.Reset(r, w)
}
