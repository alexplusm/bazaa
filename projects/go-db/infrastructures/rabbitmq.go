package infrastructures

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

const (
	delayExchangeName = "delay"
	delayExchangeKing = "x-delayed-message"
)

const (
	screenshotQueue = "screenshotQueue"
)

type RabbitMQClient struct {
	Channel *amqp.Channel
}

func initRabbitMQ() (*RabbitMQClient, error) {
	url := os.Getenv("RABBIT_MQ_URL")
	if url == "" {
		return nil, fmt.Errorf("init rabbit mq: url not specified")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("could not establish connection with RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not open RabbitMQ channel: %v", err)
	}

	client := &RabbitMQClient{Channel: ch}

	if err := client.declareDelayExchange(); err != nil {
		return nil, fmt.Errorf("delay exchange: %v", err)
	}

	if err := client.declareQueue(); err != nil {
		return nil, fmt.Errorf("delay exchange: %v", err)
	}

	return client, nil
}

func (c *RabbitMQClient) declareDelayExchange() error {
	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"

	err := c.Channel.ExchangeDeclare(
		delayExchangeName,
		delayExchangeKing,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return fmt.Errorf("declaring exchange: %v", err)
	}
	return nil
}

func (c *RabbitMQClient) declareQueue() error {
	queue, err := c.Channel.QueueDeclare(
		screenshotQueue,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue declaring: %v", err)
	}

	err = c.Channel.QueueBind(
		queue.Name,
		"",
		delayExchangeName,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("queue binding: %v", err)
	}

	return nil
}

func (c *RabbitMQClient) consum() {

	c.Channel.Consume()

	// 	msgs, err := ch.Consume(
	//		q.Name, // queue
	//		"",     // consumer
	//		true,   // auto-ack
	//		false,  // exclusive
	//		false,  // no-local
	//		false,  // no-wait
	//		nil,    // args
	//	)

}
