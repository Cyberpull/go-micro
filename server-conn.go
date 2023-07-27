package gosrv

import (
	"net"
	"sync"
)

type ServerConn interface {
	NetConn
}

type pServerConn struct {
	*netConn

	mutex sync.Mutex
	info  *Info
}

func (s *pServerConn) WriteRequest(req *pRequest) (n int, err error) {
	return writeRequest(s, req)
}

// ========================

func newServerConn(conn net.Conn) *pServerConn {
	return &pServerConn{
		netConn: newNetConn(conn),
	}
}
