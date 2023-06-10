package gosrv

import "net"

type serverClientInstance struct {
	client *pClientConn
	server *pServer
}

func (i *serverClientInstance) Start() {
	// ...
}

func (i *serverClientInstance) Init() (err error) {
	return
}

// =====================

func newServerClientInstance(server *pServer, conn net.Conn) *serverClientInstance {
	return &serverClientInstance{
		client: newClientConn(conn),
		server: server,
	}
}
