package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //Name of the exchange
		"topic",      //Type
		true,         //Durable
		false,        //auto-delete
		false,        //internal
		false,        //no-wait
		nil,          //arguements
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when unused
		true,  //exclusive
		false, //no-wait
		nil,   //arguments
	)
}
