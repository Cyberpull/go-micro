package gosrv

import (
	"sync"
)

type KCallback[K any] func(i K) string

type KCollection[K any, T any] interface {
	Has(i K) bool
	GetAll(i K) (values []T)
	Get(i K) (value T, ok bool)
	Add(i K, value T)
	AddOne(i K, value T) (ok bool)
}

type pKCollection[K any, T any] struct {
	mutex       sync.Mutex
	keyCallback KCallback[K]
	values      map[string][]T
}

func (h *pKCollection[K, T]) key(i K) string {
	if h.keyCallback != nil {
		return h.keyCallback(i)
	}

	return ""
}

func (h *pKCollection[K, T]) update(i K) {
	key := h.key(i)

	if _, ok := h.values[key]; !ok {
		h.values[key] = make([]T, 0)
	}
}

func (h *pKCollection[K, T]) Has(i K) bool {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	key := h.key(i)

	all, ok := h.values[key]

	return ok && len(all) > 0
}

func (h *pKCollection[K, T]) GetAll(i K) (values []T) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.update(i)

	key := h.key(i)

	return h.values[key]
}

func (h *pKCollection[K, T]) Get(i K) (value T, ok bool) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.update(i)

	key := h.key(i)

	var values []T

	if values, ok = h.values[key]; ok {
		if ok = len(values) > 0; !ok {
			return
		}

		value = values[0]
	}

	return
}

func (h *pKCollection[K, T]) Add(i K, value T) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.update(i)

	key := h.key(i)

	h.values[key] = append(h.values[key], value)
}

func (h *pKCollection[K, T]) AddOne(i K, value T) (ok bool) {
	if !h.Has(i) {
		h.Add(i, value)
		ok = true
	}

	return
}

func (h *pKCollection[K, T]) ClearKey(i K) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	key := h.key(i)

	h.values[key] = make([]T, 0)
}

func (h *pKCollection[K, T]) Clear() {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.values = make(map[string][]T)
}

// ================================

func newKCollection[K any, T any](keyCallback KCallback[K]) *pKCollection[K, T] {
	return &pKCollection[K, T]{
		values:      make(map[string][]T),
		keyCallback: keyCallback,
	}
}
