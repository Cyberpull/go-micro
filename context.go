package gosrv

type Context interface {
	Update(v any, codes ...int) (n int, err error)
	UpdateAll(v any, codes ...int)
	Output(v any, code ...int) (data Output)
	Error(v any, code ...int) (data Output)
}

// ======================

type pContext struct {
	server  *pServer
	client  *pClientConn
	request *pRequest
}

func (ctx *pContext) Update(v any, codes ...int) (n int, err error) {
	data, err := newUpdate(ctx.request.Method, ctx.request.Channel, v, codes...)

	if err != nil {
		return
	}

	return writeUpdate(ctx.client, data)
}

func (ctx *pContext) UpdateAll(v any, codes ...int) {
	data, err := newUpdate(ctx.request.Method, ctx.request.Channel, v, codes...)

	if err != nil {
		return
	}

	for _, instance := range ctx.server.instances {
		writeUpdate(instance.client, data)
	}
}

func (ctx *pContext) Output(v any, codes ...int) (data *pOutput) {
	code := one(200, codes)

	return &pOutput{
		Code:    code,
		Content: v,
	}
}

func (ctx *pContext) Error(v any, code ...int) (data *pOutput) {
	return ctx.Output(v, code...)
}

// ======================

func newContext(instance *serverClientInstance, request *pRequest) *pContext {
	return &pContext{
		server:  instance.server,
		client:  instance.client,
		request: request,
	}
}
