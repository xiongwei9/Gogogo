package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent
type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.Lock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, slice DataChannelSlice) {
			for _, ch := range slice {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.Unlock()
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func printDataEvent(ch string, data DataEvent) {
	log.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}
func publishTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	}
}

func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publishTo("topic1", "Hi, topic 1")
	go publishTo("topic2", "Welcome to topic 2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
