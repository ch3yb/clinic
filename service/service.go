package service

import (
	"github.com/ch3yb/clinic/env"
	"github.com/ch3yb/clinic/graph"
	"github.com/ch3yb/clinic/graph/resolvers"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Start() {

	r := resolvers.Resolver{}
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &r}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/api"))
	log.Printf("Starting http server on port " + "8080")

	err := http.ListenAndServe(":"+"8080", nil)
	if err != nil {
		log.Fatalf("Error http server on port %v: %v", env.Conf.HttpPort, err)
	}
}
