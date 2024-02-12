package events

import (
	"fmt"
	"log"

	"github.com/mustafakemalgordesli/go-commerce/config"
	"github.com/mustafakemalgordesli/go-commerce/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func MailVerificationEvent() {
	configs := config.GetConfig()

	conn := rabbitmq.GetConn()

	channel, _ := conn.Channel()

	for {
		msgs, err := channel.Consume(
			configs.RabbitMq.MailVerifiedQueue,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalf("Rabbitmq consume error: %v", err)
		} else {
			forever := make(chan bool)
			go func() {
				for msg := range msgs {
					fmt.Printf("Received Message: %s\n", msg.Body)
				}
			}()

			fmt.Println("Waiting for messages...")
			<-forever
		}
	}
}

func PublishMail(mail string) {
	configs := config.GetConfig()

	conn := rabbitmq.GetConn()

	channel, err := conn.Channel()

	if err != nil {
		log.Fatalf("Rabbitmq send mail err %v", err)
	}

	channel.Publish(
		"",
		configs.RabbitMq.MailVerifiedQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mail),
		},
	)

}
