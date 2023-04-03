package qn

type Response interface {
	_unImplement()
}

type SuccessResponse struct {
	Response
	Payload interface{}
}

type ErrorResponse struct {
	Response
	Message string
	Code    uint
	Err     error
}
