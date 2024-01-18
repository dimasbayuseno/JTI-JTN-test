package response

type GeneralResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
