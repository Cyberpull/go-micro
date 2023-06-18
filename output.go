package gosrv

type Output interface {
	//
}

type pOutput struct {
	Code    int
	Content any
}

// ======================

func newOutput() *pOutput {
	return &pOutput{}
}
