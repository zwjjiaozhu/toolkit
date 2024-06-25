package requests

// Response 统一HTTP 返回结果
type Response struct {
	Header     map[string]string `json:"header"`
	Body       []byte            `json:"body"`
	StatusCode int               `json:"statusCode"`
}

// JzWebResponse JzWeb的返回结果
//type JzWebResponse struct {
//	Code int    `json:"code,omitempty"`
//	Msg  string `json:"msg,omitempty"`
//	Data any    `json:"data,omitempty"`
//}
