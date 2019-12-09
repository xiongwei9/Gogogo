package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestConsumer(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		t.Fatalf("init consumer error: %s", err.Error())
	}

	defer func() {
		err = consumer.Close()
		if err != nil {
			t.Fatalf("close consumer error")
		}
	}()

	partitionList, err := consumer.Partitions("test")
	if err != nil {
		t.Fatalf("get partitions error: %s", err.Error())
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetOldest)
		if err != nil {
			t.Fatalf("get partition consumer error: %s", err.Error())
		}

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			defer pc.AsyncClose()
			for msg := range pc.Messages() {
				log.Printf("%s --- Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
}
