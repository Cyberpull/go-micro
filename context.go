package gosrv

type Context interface {
	Server() NetConn
	Client() NetConn
	Request() Request
}

// ======================

type pContext struct {
	server  *pServerConn
	client  *pClientConn
	request *pRequest
}

func (ctx pContext) Server() *pServerConn {
	return ctx.server
}

func (ctx pContext) Client() *pClientConn {
	return ctx.client
}

func (ctx pContext) Request() *pRequest {
	return ctx.request
}

// ======================

func newContext(server *pServerConn, client *pClientConn, request *pRequest) *pContext {
	return &pContext{
		server:  server,
		client:  client,
		request: request,
	}
}
