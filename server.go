package gosrv

import (
	"net"
	"sync"

	_ "cyberpull.com/gotk/env"
)

type Server interface {
	Listen(errChan ...chan error)
}

// =====================

type pServer struct {
	mutex    sync.Mutex
	listener net.Listener
	opts     ServerOptions
}

func (p *pServer) Listen(errChan ...chan error) {
	var err error

	defer p.Stop()

	if p.listener, err = listen(p.opts); err != nil {
		writeOne(errChan, err)
		return
	}

	writeOne(errChan, nil)

	var conn net.Conn

	for {
		if conn, err = p.listener.Accept(); err != nil {
			break
		}

		go p.handleIncomingConnection(conn)
	}
}

func (p *pServer) Stop() error {
	if p.listener != nil {
		return p.listener.Close()
	}

	return nil
}

// =====================

func NewServer(opts ServerOptions) Server {
	value := &pServer{
		opts: opts,
	}

	return value
}
