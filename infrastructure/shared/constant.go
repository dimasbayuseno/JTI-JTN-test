package shared

type ContextKeyRequestIDType string

func (c ContextKeyRequestIDType) String() string {
	return string(c)
}

const (
	ContextKeyRequestID ContextKeyRequestIDType = "request_id"
)
