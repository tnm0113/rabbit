package http

import (
	"net/http"

	"github.com/tnm0113/rabbit/internal/pkg/rabbitmq"
	"github.com/tnm0113/rabbit/internal/pkg/user"
)

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{http.NewServeMux()}
}

func (r *Router) RegisterUsers(rabbitmq *rabbitmq.RabbitMQ) {
	create := user.NewCreate(rabbitmq)

	r.HandleFunc("/users/create", create.Handle)
}
