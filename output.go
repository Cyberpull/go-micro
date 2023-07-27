package gosrv

type Output interface {
	GetCode() int
	GetContent() any
}

type pOutput struct {
	Code    int
	Content any
}

func (o pOutput) GetCode() int {
	return o.Code
}

func (o pOutput) GetContent() any {
	return o.Content
}

// ======================

func newOutput() *pOutput {
	return &pOutput{}
}
