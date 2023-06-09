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
}

// ========================

func newServerConn(conn net.Conn) *pServerConn {
	return &pServerConn{
		netConn: newNetConn(conn),
	}
}
