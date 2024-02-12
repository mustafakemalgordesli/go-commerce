package rabbitmq

import (
	"fmt"

	"github.com/mustafakemalgordesli/go-commerce/config"
	"github.com/streadway/amqp"
)

var (
	Conn   *amqp.Connection
	ConErr error
)

type RabbitMq struct {
	*amqp.Connection
}

func Setup() error {
	configs := config.GetConfig()

	connection, err := amqp.Dial(configs.RabbitMq.Connection)

	if err != nil {
		ConErr = err
		return err
	}

	channel, err := connection.Channel()

	if err != nil {
		ConErr = err
		return err
	}

	queue, err := channel.QueueDeclare(
		configs.RabbitMq.MailVerifiedQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(queue)

	if err != nil {
		ConErr = err
		return err
	}

	Conn = connection

	return nil
}

func GetConn() *amqp.Connection {
	return Conn
}

func GetConnErr() error {
	return ConErr
}
