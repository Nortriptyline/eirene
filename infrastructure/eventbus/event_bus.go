package eventbus

import "fmt"

type EventBus struct {
	subscribers map[string][]func(event interface{})
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(event interface{})),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler func(event interface{})) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

func (eb *EventBus) Publish(event interface{}) {
	eventType := fmt.Sprintf("%T", event)
	if handlers, found := eb.subscribers[eventType]; found {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

func (eb *EventBus) PublishMany(events []interface{}) {
	for _, event := range events {
		eb.Publish(event)
	}
}
