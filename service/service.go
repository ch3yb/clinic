package service

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ch3yb/clinic/env"
	"github.com/ch3yb/clinic/graph"
	"github.com/ch3yb/clinic/graph/resolvers"
	db "github.com/ch3yb/clinic/service/database"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
)

func loggerInit() *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	return logger
}

func Start() {
	port := env.Conf.HttpPort

	database, err := db.StartDatabase()
	if err != nil {
		return
	}

	s = New(database, loggerInit())
	s.TestMethod()
	log.Println("test env: ", env.Conf.HttpPort)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type Server struct {
	db     *sql.DB
	Logger *zap.Logger
}

var s *Server

func New(db *sql.DB, logger *zap.Logger) *Server {
	return &Server{
		db:     db,
		Logger: logger,
	}
}
