package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
)

func TestProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		t.Fatalf("producer connect error: %s", err.Error())
	}
	defer func() {
		err := client.Close()
		if err != nil {
			t.Fatalf("producer close error: %v", err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("message from go"),
	}
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		t.Fatalf("producer push message error: %s", err.Error())
	}
	t.Logf("pid: %v | offset: %v\n", pid, offset)
}

func TestAsyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		t.Fatalf("producer connect error: %s", err.Error())
	}
	defer func() {
		err := client.Close()
		if err != nil {
			t.Fatalf("producer close error: %v", err)
		}
	}()

	ch := make(chan struct{})
	go func(producer sarama.AsyncProducer) {
		errors := producer.Errors()
		success := producer.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					t.Logf("producer push message error: %v", err)
				}
			case ret := <-success:
				t.Logf("producer push message success: %v", ret)
				ch <- struct{}{}
				return
			}
		}
	}(client)

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("message from go async"),
	}
	client.Input() <- msg
	<-ch
}
