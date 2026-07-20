package eventbus

import (
	"sync"
)

type Event[T any] struct {
	Topic   string
	Payload T
}

type EventBus[T any] struct {
	listeners    sync.Map
	mu           sync.RWMutex
	errorHandler func(topic string, err error)
}

type EventListener[T any] struct {
	priority int
	listener func(eventData T)
	async    bool
	filter   func(eventData T) bool
}

func NewEventBus[T any]() *EventBus[T] { _ = "STUB: not implemented"; return nil }

func (eb *EventBus[T]) Subscribe(topic string, listener func(eventData T), async bool, priority int, filter func(eventData T) bool) {
	_ = "STUB: not implemented"
	return
}

func (eb *EventBus[T]) Unsubscribe(topic string, listener func(eventData T)) {
	_ = "STUB: not implemented"
	return
}

func (eb *EventBus[T]) Publish(event Event[T]) { _ = "STUB: not implemented"; return }

func (eb *EventBus[T]) publishToListener(listener *EventListener[T], event Event[T]) {
	_ = "STUB: not implemented"
	return
}

func (eb *EventBus[T]) SetErrorHandler(handler func(topic string, err error)) {
	_ = "STUB: not implemented"
	return
}

func (eb *EventBus[T]) ClearListeners() { _ = "STUB: not implemented"; return }

func (eb *EventBus[T]) ClearListenersByTopic(topic string) { _ = "STUB: not implemented"; return }

func (eb *EventBus[T]) GetListenersCount(topic string) int { _ = "STUB: not implemented"; return 0 }

func (eb *EventBus[T]) GetAllListenersCount() int { _ = "STUB: not implemented"; return 0 }

func (eb *EventBus[T]) GetEvents() []string { _ = "STUB: not implemented"; return nil }
