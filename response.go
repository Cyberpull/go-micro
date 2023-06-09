package gosrv

type Response interface {
	//
}

type pResponse struct {
	//
}

func newResponse() Response {
	value := &pResponse{}
	return value
}
