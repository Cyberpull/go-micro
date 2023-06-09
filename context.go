package gosrv

type Context interface {
	//
}

// ======================

type pContext struct {
	server *pServerConn
}

// ======================

func newContext(server *pServerConn) *pContext {
	return &pContext{
		server: server,
	}
}
