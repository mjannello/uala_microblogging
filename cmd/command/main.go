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
	commandeventpublisher "uala/internal/command/eventpublisher/rabbitmq"
	commandrepository "uala/internal/command/repository/postgres"
	commandservice "uala/internal/command/service"
	httprouter "uala/internal/http"
)

func main() {
	db := connectDB()

	eventStore := commandrepository.NewPostgresEventStore(db)

	eventPublisher := commandeventpublisher.NewRabbitMQEventPublisher()

	commandService := commandservice.NewCommandService(eventStore, eventPublisher)

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

func connectDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=mjannello dbname=uala_events_postgres password=uala_db_password sslmode=disable")
	if err != nil {
		log.Fatal("Fail when connecting to db:", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	return db
}

func run(server *http.Server) error {
	log.Printf("Server HTTP started in %s", server.Addr)
	return server.ListenAndServe()
}
