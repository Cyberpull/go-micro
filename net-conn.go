package gosrv

import "net"

type NetConn interface {
	NetIO

	Close() error
}

// =======================

type netConn struct {
	*netIO

	conn net.Conn
}

func (c *netConn) Close() error {
	return c.conn.Close()
}

// =======================

func newNetConn(conn net.Conn) *netConn {
	return &netConn{
		netIO: newNetIO(conn),
		conn:  conn,
	}
}
