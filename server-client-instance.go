package gosrv

import (
	"fmt"
	"net"

	"cyberpull.com/gotk/uuid"
)

type serverClientInstance struct {
	uuid   string
	client *pClientConn
	server *pServer
}

func (i *serverClientInstance) Start() {
	defer func() {
		r := recover()

		if r != nil {
			// Do something
		}
	}()

	mustWriteState(i.client, requestState, true)

	for {
		req, err := i.client.ReadRequest()

		if err != nil {
			continue
		}

		go i.handleRequest(req)
	}
}

func (i *serverClientInstance) handleRequest(req *pRequest) {
	// ctx := newContext(i, req)
}

func (i *serverClientInstance) prepare() (err error) {
	info := i.server.opts.Info
	welcome := fmt.Sprintf("Welcome to %s", info.Name)

	if _, err = i.client.WriteStringLine(welcome); err != nil {
		return
	}

	var cinfo *Info

	if cinfo, err = i.client.ReadInfo(); err != nil {
		return
	}

	i.client.info = cinfo

	_, err = i.client.WriteInfo(info)

	return
}

// =====================

func newServerClientInstance(server *pServer, conn net.Conn) (instance *serverClientInstance, err error) {
	instanceUUID, err := uuid.Generate()

	instance = &serverClientInstance{
		uuid:   instanceUUID,
		client: newClientConn(conn),
		server: server,
	}

	return
}
