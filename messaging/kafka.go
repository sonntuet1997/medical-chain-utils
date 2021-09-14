package messaging

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaOptions struct {
	Addr  string
	Topic string
}

func ConnectReader(opts KafkaOptions) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{opts.Addr},
		Topic:    opts.Topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
}

func ConnectWriter(opts KafkaOptions) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Topic:        opts.Topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 5 * time.Second,
		MaxAttempts:  4,
		Async:        true,
		BatchSize:    1,
		Brokers:      []string{opts.Addr},
	})

}