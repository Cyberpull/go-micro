package gosrv

import (
	"net"
	"sync"
)

type ClientConn interface {
	NetConn
	ClientUpdater

	ReadRequest() (req *Request, err error)
	WriteResponse(data *Response) (n int, err error)
	WriteError(data *Data) (n int, err error)
}

type pClientConn struct {
	*netConn

	mutex sync.Mutex
	info  *Info
}

func (c *pClientConn) ReadRequest() (req *pRequest, err error) {
	return getRequest(c)
}

func (c *pClientConn) WriteResponse(data *Response) (n int, err error) {
	return writeResponse(c, data)
}

func (c *pClientConn) WriteError(data *Data) (n int, err error) {
	return writeError(c, data)
}

// ========================

func newClientConn(conn net.Conn) *pClientConn {
	return &pClientConn{
		netConn: newNetConn(conn),
	}
}
