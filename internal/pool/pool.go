package pool

import "sync"

type Pool[T Resetter] struct {
	items   []T
	mu      sync.RWMutex
	factory func() T
}

func NewPool[T Resetter](factory func() T, capacity int) *Pool[T] {
	return &Pool[T]{
		items:   make([]T, capacity), // capacity used for prewarm slice
		factory: factory,
	}
}

func (p *Pool[T]) Get() T {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.items) > 0 {
		item := p.items[0]
		p.items = p.items[1:]
		return item
	}

	return p.factory()
}

func (p *Pool[T]) Put(item T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	item.Reset()
	p.items = append(p.items, item)
}

func (p *Pool[T]) Len() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.items)
}
