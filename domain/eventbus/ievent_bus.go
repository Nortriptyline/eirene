package eventbus

type IEventBus interface {
	Subscribe(eventType string, handler func(event interface{}))
	Publish(event interface{})
	PublishMany(events []interface{})
}
