package conn

import (
	"fmt"
	"github.com/streadway/amqp"
)

func RabbitMQConnection(addr, username, password string) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s",
		username, password, addr,
	)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
