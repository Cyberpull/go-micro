package gosrv

type Client interface {
}

type pClient struct {
	//
}

func NewClient() Client {
	value := &pClient{}
	return value
}
