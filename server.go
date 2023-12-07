package gosrv

import (
	"net"
	"sync"
)

type BootHandler func() (err error)
type ReadyHandler func() (err error)
type ClientBootHandler func(client NetIO) (err error)
type ClientReadyHandler func(client ClientUpdater) (err error)
type RequestHandlerSubscriber func(collection RequestHandlerCollection)

// =====================

type Server interface {
	Listen(errChan ...chan error)
	OnBoot(handlers ...BootHandler)
	OnReady(handlers ...ReadyHandler)
	OnClientBoot(handlers ...ClientBootHandler)
	OnClientReady(handlers ...ClientReadyHandler)
	RequestHandlers(subscribers ...RequestHandlerSubscriber)
	Stop() error
}

// =====================

type pServer struct {
	mutex       sync.Mutex
	listener    net.Listener
	opts        ServerOptions
	instances   map[string]*serverClientInstance
	collection  *pRequestHandlerCollection
	boot        []BootHandler
	ready       []ReadyHandler
	clientBoot  []ClientBootHandler
	clientReady []ClientReadyHandler
}

func (p *pServer) Listen(errChan ...chan error) {
	var err error

	defer p.Stop()

	if err = validator.Validate(p.opts); err != nil {
		sendOne(errChan, err)
		return
	}

	if err := p.execBoot(); err != nil {
		sendOne(errChan, err)
		return
	}

	if p.listener, err = listen(p.opts); err != nil {
		sendOne(errChan, err)
		return
	}

	if err := p.execReady(); err != nil {
		sendOne(errChan, err)
		return
	}

	sendOne(errChan, nil)

	for {
		var conn net.Conn

		if conn, err = p.listener.Accept(); err != nil {
			break
		}

		go p.handleIncomingConnection(conn)
	}
}

func (p *pServer) OnBoot(handlers ...BootHandler) {
	p.boot = append(p.boot, handlers...)
}

func (p *pServer) OnReady(handlers ...ReadyHandler) {
	p.ready = append(p.ready, handlers...)
}

func (p *pServer) RequestHandlers(subscribers ...RequestHandlerSubscriber) {
	for _, subscriber := range subscribers {
		subscriber(p.collection)
	}
}

func (p *pServer) OnClientBoot(handlers ...ClientBootHandler) {
	p.clientBoot = append(p.clientBoot, handlers...)
}

func (p *pServer) OnClientReady(handlers ...ClientReadyHandler) {
	p.clientReady = append(p.clientReady, handlers...)
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
		opts:        opts,
		instances:   make(map[string]*serverClientInstance),
		boot:        make([]BootHandler, 0),
		ready:       make([]ReadyHandler, 0),
		collection:  newRequestHandlerCollection(),
		clientBoot:  make([]ClientBootHandler, 0),
		clientReady: make([]ClientReadyHandler, 0),
	}

	return value
}
