package ginHandler

type ErrorType uint64
type Error struct {
	Err  error
	Type ErrorType
	Meta any
	code *int
}

func (msg Error) Error() string

type errorMsgs []*Error
type Context struct {
	// writermem responseWriter
	// Request   *http.Request
	// Writer    ResponseWriter

	// Params   Params
	// handlers HandlersChain
	// index    int8
	// fullPath string

	// engine       *Engine
	// params       *Params
	// skippedNodes *[]skippedNode

	// This mutex protects Keys map.
	// mu sync.RWMutex

	// // Keys is a key/value pair exclusively for the context of each request.
	// Keys map[string]any

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	// // Accepted defines a list of manually accepted formats for content negotiation.
	// Accepted []string

	// // queryCache caches the query result from c.Request.URL.Query().
	// queryCache url.Values

	// // formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// // or PUT body parameters.
	// formCache url.Values

	// // SameSite allows a server to define a cookie attribute making it impossible for
	// // the browser to send this cookie along with cross-site requests.
	// sameSite http.SameSite
}

func (c *Context) AbortWithStatusJSON(code int, jsonObj any)
func (c *Context) Next()

type HandlerFunc func(*Context)
