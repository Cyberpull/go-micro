package gosrv

import "cyberpull.com/gotk/v2/errors"

type Context interface {
	ParseContent(v any) (err error)
	Update(v any, codes ...int) (n int, err error)
	UpdateAll(v any, codes ...int)
	Output(v any, code ...int) (data Output)
	Success(v any) (data Output)
	Error(v any, code ...int) (data Output)
}

// ======================

type pContext struct {
	server  *pServer
	client  *pClientConn
	request *pRequest
}

func (ctx *pContext) ParseContent(v any) (err error) {
	return ctx.request.ParseContent(v)
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

	go func() {
		for _, instance := range ctx.server.instances {
			writeUpdate(instance.client, data)
		}
	}()
}

func (ctx *pContext) Output(v any, codes ...int) (data Output) {
	code := one(200, codes)

	return &pOutput{
		Code:    code,
		Content: v,
	}
}

func (ctx *pContext) Success(v any) (data Output) {
	return ctx.Output(v, 200)
}

func (ctx *pContext) Error(v any, code ...int) (data Output) {
	err := errors.From(v, code...)
	return ctx.Output(err.Error(), err.Code())
}

// ======================

func newContext(instance *serverClientInstance, request *pRequest) *pContext {
	return &pContext{
		server:  instance.server,
		client:  instance.client,
		request: request,
	}
}
