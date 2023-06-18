package gosrv

import "net"

type NetConn interface {
	NetIO

	ReadInfo() (*Info, error)
	WriteInfo(info *Info) (n int, err error)

	Close() error
}

// =======================

type netConn struct {
	*netIO

	conn net.Conn
}

func (c *netConn) ReadInfo() (info *Info, err error) {
	return getInfo(c)
}

func (c *netConn) WriteInfo(info *Info) (n int, err error) {
	return writeInfo(c, info)
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
