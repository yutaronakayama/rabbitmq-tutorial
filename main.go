package main

import (
	"fmt"

	"github.com/TutorialEdge/rabbitmq-crash-course/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("hi")
	if err != nil {
		return err
	}

	app.Rmq.Consume()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Failed to run Setting Up our Application")
		fmt.Println(err)
	}
}
