package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
}

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "allen_test",
		Balancer: &kafka.LeastBytes{},
		//CompressionCodec: snappy.NewCompressionCodec(),
	})
	defer w.Close()

	userInfo := new(UserInfo)
	userInfo.Id = 12
	userInfo.Name = "aaa"
	userInfo.Age = 10

	value, err := json.Marshal(userInfo)

	if err != nil {
		return
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(strconv.Itoa(userInfo.Id)),
			Value: value,
		},
	)

	if err != nil {
		log.Println("ERROR    writting message: ", err)
	} else {
		log.Println("INFO    writting message OK")
	}

}
