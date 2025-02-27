package tracing

import "github.com/IBM/sarama"

type ConsumerMessageCarrier struct {
	msg *sarama.ConsumerMessage
}

func (c ConsumerMessageCarrier) Set(key string, value string) {
	c.msg.Headers = append(c.msg.Headers, &sarama.RecordHeader{
		Key:   []byte(key),
		Value: []byte(value),
	})
}

func NewConsumerMessageCarrier(msg *sarama.ConsumerMessage) ConsumerMessageCarrier {
	return ConsumerMessageCarrier{msg: msg}
}

func (c ConsumerMessageCarrier) Get(key string) string {
	for _, h := range c.msg.Headers {
		if string(h.Key) == key {
			return string(h.Value)
		}
	}
	return ""
}

func (c ConsumerMessageCarrier) Keys() []string {
	var keys []string
	for _, h := range c.msg.Headers {
		keys = append(keys, string(h.Key))
	}
	return keys
}
