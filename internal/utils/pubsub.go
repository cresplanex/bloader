package utils

import (
	"sync"
)

// Broadcaster is a type that broadcasts a value to multiple subscribers.
type Broadcaster[T any] struct {
	subscribers map[chan T]struct{}
	mutex       sync.RWMutex
}

// NewBroadcaster creates a new Broadcaster.
func NewBroadcaster[T any]() *Broadcaster[T] {
	return &Broadcaster[T]{
		subscribers: make(map[chan T]struct{}),
	}
}

// Close closes the Broadcaster.
func (b *Broadcaster[T]) Close() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for ch := range b.subscribers {
		delete(b.subscribers, ch)
		close(ch)
	}
}

// Subscribe subscribes to the Broadcaster.
func (b *Broadcaster[T]) Subscribe() <-chan T {
	ch := make(chan T)
	b.mutex.Lock()
	b.subscribers[ch] = struct{}{}
	b.mutex.Unlock()
	return ch
}

// Unsubscribe unsubscribes from the Broadcaster.
func (b *Broadcaster[T]) Unsubscribe(ch chan T) {
	b.mutex.Lock()
	delete(b.subscribers, ch)
	close(ch)
	b.mutex.Unlock()
}

// Broadcast broadcasts a value to all subscribers.
func (b *Broadcaster[T]) Broadcast(value T) <-chan struct{} {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	doneAny := make(chan struct{}, len(b.subscribers))
	done := make(chan struct{})

	go func() {
		defer close(done)
		for range b.subscribers {
			<-doneAny
		}
	}()

	for ch := range b.subscribers {
		go func(ch chan T) {
			ch <- value
			doneAny <- struct{}{}
		}(ch)
	}

	return done
}
