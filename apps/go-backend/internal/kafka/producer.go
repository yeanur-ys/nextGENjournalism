package kafka

type Producer struct{}

func (Producer) Publish(topic string, payload []byte) error {
	_ = topic
	_ = payload
	return nil
}
