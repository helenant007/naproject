package nsq

import (
	"log"

	utilNsq "github.com/bitly/go-nsq"
	utilRedis "github.com/helenant007/naproject/utils/redis"
)

func InitConsumer() {
	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	config := utilNsq.NewConfig()
	conn, _ := utilNsq.NewConsumer("naproject-hn-visitor-topic", "channel_naproject_hn", config)
	conn.AddHandler(utilNsq.HandlerFunc(func(message *utilNsq.Message) error {
		if string(message.Body) == "todo_count" {
			// increment redis
			_, err := utilRedis.INCR("naproject:helen:visitorcount")
			if err != nil {
				log.Println(err)
				return err
			}
		}
		return nil
	}))
	err := conn.ConnectToNSQLookupd("devel-go.tkpd:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	// wg.Wait()
}
