package internal

import (
	"github.com/gabrielpsz/dm-challenge/repository"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var amqpConnection *amqp.Connection

func StartQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	amqpConnection = conn
	failOnError(err, "Failed to connect to RabbitMQ")
	// TODO - Verificar fechamento de conexão

	incrementMessageDelivery, err := createQueue(err, "increment", "incremented", "stock")
	decrementMessageDelivery, err := createQueue(err, "decrement", "decremented", "stock")

	go consumeMessages(incrementMessageDelivery, increment)
	consumeMessages(decrementMessageDelivery, decrement)

}

func createQueue(err error, queueName, routingKey, exchangeName string) (<-chan amqp.Delivery, error) {
	ch, err := amqpConnection.Channel()
	queue, err := ch.QueueDeclare(
		queueName, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.QueueBind(
		queue.Name,    // queue name
		routingKey, // routing key
		exchangeName,       // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")
	return msgs, err
}

func consumeMessages(deliveries <-chan amqp.Delivery, deliveryFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		deliveryFunc(d)
	}
}

func increment(d amqp.Delivery) {
	productName := strings.Replace(string(d.Body), "\"", "", -1)
	products, e := repository.GetProductByName(productName)
	fmt.Println("Incremented -> ", productName)
	if (len(products) > 0) {
		product := &products[0]
		product.Quantity += 1
		repository.UpdateProduct(product.ID.Hex(), product)
		if e != nil {
			log.Println("Error: ", e)
		}
	} else {
		fmt.Printf("Product %s not found. Ignoring.", productName)
	}
}

func decrement(d amqp.Delivery) {
	productName := strings.Replace(string(d.Body), "\"", "", -1)
	products, e := repository.GetProductByName(productName)
	fmt.Println("Decremented -> ", productName)
	if (len(products) > 0) {
		product := &products[0]
		if (product.Quantity > 0) {
			product.Quantity -= 1
			repository.UpdateProduct(product.ID.Hex(), product)
		}
		if e != nil {
			log.Println("Error: ", e)
		}
	} else {
		fmt.Printf("Product %s not found. Ignoring.", productName)
	}
}