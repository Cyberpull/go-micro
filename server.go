package gosrv

import (
	"net"
	"sync"
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
	//
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
