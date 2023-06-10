package gosrv

import "net"

func (p *pServer) handleIncomingConnection(conn net.Conn) {
	instance := newServerClientInstance(p, conn)

	if err := instance.Init(); err != nil {
		return
	}

	instance.Start()
}
