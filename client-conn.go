package gosrv

import (
	"net"
	"sync"
)

type ClientConn interface {
	NetConn
}

type pClientConn struct {
	*netConn

	mutex sync.Mutex
}

// ========================

func newClientConn(conn net.Conn) *pClientConn {
	return &pClientConn{
		netConn: newNetConn(conn),
	}
}
