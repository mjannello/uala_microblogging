package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	httprouter "uala/internal/http/query"
	querycontroller "uala/internal/query/controller"
	"uala/internal/query/eventconsumer/rabbitmq_consumer"
	"uala/internal/query/repository/mongodb"
	"uala/internal/query/service"
)

func main() {
	mongoDBRepo, err := mongodb.NewMongoDBRepository("mongodb://mongodb:27017")
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	queryService := service.NewQueryService(mongoDBRepo)

	go func() {
		eventConsumer, err := rabbitmq_consumer.NewRabbitMQEventConsumer(queryService)
		if err != nil {
			log.Fatal("Error setting up events consumer:", err)
		}

		if err := eventConsumer.StartConsuming(); err != nil {
			log.Fatal("Error starting events consumer:", err)
		}
	}()

	queryController := querycontroller.NewQueryController(queryService)

	routerHandler := httprouter.NewRouterHandler(queryController)

	router := mux.NewRouter()
	routerHandler.ConfigureRoutes(router)

	serverAddr := ":8080"
	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := run(server); err != nil {
			log.Fatal("Fail running app:", err)
		}
	}()

	<-stopChan

}

func run(server *http.Server) error {
	log.Printf("Server HTTP started in %s", server.Addr)
	return server.ListenAndServe()
}
