package main

import (
	"log"

	"github.com/tnm0113/rabbit/internal/app"
	"github.com/tnm0113/rabbit/internal/config"
	"github.com/tnm0113/rabbit/internal/pkg/http"
	"github.com/tnm0113/rabbit/internal/pkg/rabbitmq"
	"github.com/tnm0113/rabbit/internal/pkg/user"
)

func main() {
	// Config
	cnf := config.New()
	//

	// RabbitMQ
	rbt := rabbitmq.New(cnf.RabbitMQ)
	if err := rbt.Connect(); err != nil {
		log.Fatalln(err)
	}
	defer rbt.Shutdown()
	//

	// AMQP setup
	userAMQP := user.NewAMQP(cnf.UserAMQP, rbt)
	if err := userAMQP.Setup(); err != nil {
		log.Fatalln(err)
	}
	//

	// HTTP router
	rtr := http.NewRouter()
	rtr.RegisterUsers(rbt)
	//

	// HTTP server
	srv := http.NewServer(cnf.HTTPAddress, rtr)
	//

	// Run
	if err := app.New(srv).Run(); err != nil {
		log.Fatalln(err)
	}
	//
}
