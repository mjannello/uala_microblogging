package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	commandcontroller "uala/internal/command/controller"
	commandeventpublisher "uala/internal/command/eventpublisher/rabbitmq_producer"
	commandrepository "uala/internal/command/eventstore/postgres"
	commandservice "uala/internal/command/service"
	httprouter "uala/internal/http/command"
	"uala/pkg/clock"
)

func main() {
	db, err := sqlx.Open("postgres", "postgresql://mjannello:uala_db_password@postgres:5432/uala_events_postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Fail when connecting to db:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	eventStore := commandrepository.NewPostgresEventStore(db)

	eventPublisher, err := commandeventpublisher.NewRabbitMQEventPublisher()
	realClock := clock.NewClock()

	commandService := commandservice.NewCommandService(eventStore, eventPublisher, realClock)

	commandController := commandcontroller.NewCommandController(commandService)

	routerHandler := httprouter.NewRouterHandler(commandController)

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
