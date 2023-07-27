package gosrv

import (
	"fmt"
	"net"

	"cyberpull.com/gotk/v2"
	"cyberpull.com/gotk/v2/errors"
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

	// mustWriteState(i.client, requestState, true)

	for {
		req, err := i.client.ReadRequest()

		if err != nil {
			continue
		}

		go i.handleRequest(req)
	}
}

func (i *serverClientInstance) handleRequest(req *pRequest) {
	var err error
	var handler RequestHandler

	defer func() {
		if r := recover(); r != nil {
			err = errors.From(r)
		}

		if err != nil {
			writeErrorResponse(i.client, req, err)
		}
	}()

	method, channel := req.Method, req.Channel

	if handler, err = i.server.collection.Get(method, channel); err != nil {
		// writeErrorResponse(i.client, req, err)
		return
	}

	ctx := newContext(i, req)

	output := handler(ctx)

	writeOutputResponse(i.client, req, output)
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
	instanceUUID, err := gotk.UUID()

	instance = &serverClientInstance{
		uuid:   instanceUUID,
		client: newClientConn(conn),
		server: server,
	}

	return
}
