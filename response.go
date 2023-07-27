package gosrv

type Response interface {
	DataWithCode

	GetRequest() Request
}

type pResponse struct {
	pDataWithCode

	Request *pRequest
}

func (r *pResponse) GetRequest() Request {
	return r.Request
}

// ==================================

func newResponse(req *pRequest) *pResponse {
	return &pResponse{
		Request: req,
	}
}
