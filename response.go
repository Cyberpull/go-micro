package gosrv

type Response struct {
	pDataWithCode

	Request *pRequest
}

func newResponse(req *pRequest) *Response {
	return &Response{
		Request: req,
	}
}
