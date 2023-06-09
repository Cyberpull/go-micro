package gosrv

type Request interface {
	//
}

type pRequest struct {
	//
}

func newRequest() Request {
	value := &pRequest{}
	return value
}
