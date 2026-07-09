package kafka

type Consumer struct{}

func (Consumer) Start(handler func([]byte) error) error {
	if handler == nil {
		return nil
	}
	return handler(nil)
}
