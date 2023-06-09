package gosrv

import (
	"bufio"
	"net"
)

type NetIO interface {
	NetReader
	NetWriter
}

// ========================

type netIO struct {
	conn   net.Conn
	reader *bufio.Reader
}

// ========================

func newNetIO(conn net.Conn) *netIO {
	return &netIO{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}
}
