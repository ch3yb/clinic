package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ch3yb/clinic/api/service"
	"github.com/ch3yb/clinic/env"
	"github.com/ch3yb/clinic/graph"
	"github.com/ch3yb/clinic/graph/resolvers"
	"log"
	"net/http"
)

func Start() {
	port := env.Conf.HttpPort

	s := service.New()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		Service: s,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
