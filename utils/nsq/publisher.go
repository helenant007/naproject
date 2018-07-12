package nsq

import (
	"fmt"

	nsq "github.com/bitly/go-nsq"
)

func PublishData(topic string, msg string) {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("devel-go.tkpd:4150", config)
	if err != nil {
		fmt.Println("Unable to publish data")
		return
	}

	err = w.Publish(topic, []byte(msg))
	if err != nil {
		fmt.Println("Could not connect")
		return
	}

	w.Stop()
}
