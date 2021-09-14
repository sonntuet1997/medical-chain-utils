package messaging

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaOptions struct {
	addr  string
	topic string
}

func ConnectReader(opts KafkaOptions) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{opts.addr},
		Topic:    opts.topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
}

func ConnectWriter(opts KafkaOptions) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Topic:        opts.topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 5 * time.Second,
		MaxAttempts:  4,
		Async:        true,
		BatchSize:    1,
		Brokers:      []string{opts.addr},
	})

}
