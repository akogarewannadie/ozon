package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akogarewannadie/ozon//internal/config"
	"github.com/akogarewannadie/ozon//internal/db"
	"github.com/akogarewannadie/ozon//internal/graph"
)

func Run() error {
	cfg := config.LoadConfig()

	var database db.Database
	if cfg.UseInMemory {
		database = db.NewInMemoryDB()
	} else {
		postgresDB, err := db.NewPostgresDB(cfg.DatabaseURL)
		if err != nil {
			return err
		}
		database = postgresDB
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{db: database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", os.Getenv("PORT"))
	return http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
