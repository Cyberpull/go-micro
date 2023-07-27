package gosrv

import (
	"strings"

	"cyberpull.com/gotk/v2/errors"
)

type RequestHandler func(ctx Context) Output

type RequestHandlerCollection interface {
	On(method, channel string, handler RequestHandler) (err error)
}

type pRequestHandlerCollection struct {
	handlers map[string]RequestHandler
}

func (s *pRequestHandlerCollection) key(method, channel string) string {
	method = strings.ToUpper(method)
	return method + "::" + channel
}

func (s *pRequestHandlerCollection) Has(method, channel string) bool {
	key := s.key(method, channel)
	_, ok := s.handlers[key]
	return ok
}

func (s *pRequestHandlerCollection) Get(method, channel string) (handler RequestHandler, err error) {
	key := s.key(method, channel)

	handler, ok := s.handlers[key]

	if !ok {
		err = errors.Newf(`No action found for "%s" -> "%s"`, 400, method, channel)
		return
	}

	return
}

func (s *pRequestHandlerCollection) On(method, channel string, handler RequestHandler) (err error) {
	if s.Has(method, channel) {
		err = errors.Newf(`Action already exists for "%s" -> "%s"`, 500, method, channel)
		return
	}

	key := s.key(method, channel)
	s.handlers[key] = handler

	return
}

func (s *pRequestHandlerCollection) Off(method, channel string) {
	if s.Has(method, channel) {
		key := s.key(method, channel)
		delete(s.handlers, key)
	}
}

// ============================

func newRequestHandlerCollection() *pRequestHandlerCollection {
	return &pRequestHandlerCollection{
		handlers: make(map[string]RequestHandler),
	}
}
