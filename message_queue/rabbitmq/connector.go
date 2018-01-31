package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Connector struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (connector *Connector) Connection() *amqp.Connection {
	return connector.conn
}

func (connector *Connector) Channel() *amqp.Channel {
	return connector.ch
}

func (connector *Connector) Connect(connectionStr string) error {
	var err error
	connector.conn, err = amqp.Dial(connectionStr)
	if err != nil {
		return err
	}

	connector.ch, err = connector.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (connector *Connector) Close() error {
	var err error
	err = connector.conn.Close()
	if err != nil {
		return err
	}

	err = connector.ch.Close()
	if err != nil {
		return err
	}

	return nil
}
