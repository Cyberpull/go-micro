package gosrv

import "sync"

type MCollection[T any] interface {
	Has(method, channel string) bool
	GetAll(method, channel string) (values []T)
	Get(method, channel string) (value T, ok bool)
	Add(method, channel string, value T)
	AddOne(method, channel string, value T) (ok bool)
}

type pMCollection[T any] struct {
	mutex  sync.Mutex
	values map[string][]T
}

func (h *pMCollection[T]) key(method, channel string) string {
	return method + "::" + channel
}

func (h *pMCollection[T]) update(method, channel string) {
	key := h.key(method, channel)

	if _, ok := h.values[key]; !ok {
		h.values[key] = make([]T, 0)
	}
}

func (h *pMCollection[T]) Has(method, channel string) bool {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	key := h.key(method, channel)

	all, ok := h.values[key]

	return ok && len(all) > 0
}

func (h *pMCollection[T]) GetAll(method, channel string) (values []T) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	key := h.key(method, channel)

	values, _ = h.values[key]

	return
}

func (h *pMCollection[T]) Get(method, channel string) (value T, ok bool) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.update(method, channel)

	key := h.key(method, channel)

	var values []T

	if values, ok = h.values[key]; ok {
		value = values[0]
	}

	return
}

func (h *pMCollection[T]) Add(method, channel string, value T) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.update(method, channel)

	key := h.key(method, channel)

	h.values[key] = append(h.values[key], value)
}

func (h *pMCollection[T]) AddOne(method, channel string, value T) (ok bool) {
	if !h.Has(method, channel) {
		h.Add(method, channel, value)
		ok = true
	}

	return
}

func (h *pMCollection[T]) ClearKey(method, channel string) {
	key := h.key(method, channel)
	h.values[key] = make([]T, 0)
}

func (h *pMCollection[T]) Clear() {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.values = make(map[string][]T)
}

// ================================

func newMCollection[T any]() *pMCollection[T] {
	return &pMCollection[T]{
		values: make(map[string][]T),
	}
}
