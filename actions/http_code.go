package actions

var httpCode = map[string]int{
	// 2xx HTTP status codes
	"ok":      200,
	"created": 201,

	// 3xx HTTP status codes
	"found": 302,

	// 4xx HTTP status codes
	"badRequest": 400,
}
