# UALA_microblogging

## CQRS and Event Sourcing App

This application, named uala_microblogging, employs the CQRS (Command Query Responsibility Segregation) and Event Sourcing architectural patterns. Its primary goal is to provide a scalable and efficient solution for managing posts, comments, and user interactions. 

With the use of RabbitMQ it includes a manager running on port 15672 to check the status of the queues and messages.

## Purpose

### Command App

The **Command App** handles the commands and write operations of the system. It is responsible for receiving requests to add, update, or delete posts, as well as adding comments. This app generates events corresponding to these actions, which are then propagated to the Event Bus (RabbitMQ) for further processing.

### Query App

The **Query App** is responsible for handling read operations. It provides endpoints to retrieve the feed of posts, both globally and for a specific user. This separation of read and write concerns allows for optimized queries and scaling based on the application's specific needs.

## Events

The application leverages the power of events to capture state changes. The key events include:

- **PostAddedEvent**: Generated when a user adds a new post.

- **PostUpdatedEvent**: Generated when a user updates an existing post.

- **PostDeletedEvent**: Generated when a user deletes a post.

- **CommentAddedEvent**: Generated when a user adds a comment to a post.

Endpoints
---------

Every endpoint accepts the header **user_name** with a string as value with the name of the user. 

### Command App

*   **POST /api/post**: Adds a new post.

*   **PUT /api/post/{id}**: Updates an existing post.

*   **DELETE /api/post/{id}**: Deletes a post.

*   **POST /api/comment/{post\_id}**: Adds a comment to a specific post.


### Query App

*   **GET /api/feed**: Retrieves the feed of posts.

## Components

### Command Component

*   **Controllers**: Handle HTTP requests and invoke services to perform command operations.

*   **Services**: Contain the business logic for command operations, ensuring proper validation and execution.

*   **Event Store (Repository)**: Interact with the PostgreSQL database to store and retrieve events.

*   **Event Publishers**: Responsible for publishing events to the message broker (RabbitMQ) after successful command operations.


### Query Component

*   **Controllers**: Handle HTTP requests for read operations.

*   **Services**: Implement the logic for reading data from the MongoDB database, providing optimized responses and process new events from the broker.

*   **Query Repository (Read-Optimized Repository)**: Interact with the MongoDB database to retrieve data for read operations.

*   **Event Consumers**: Read events from the message broker (RabbitMQ) related to read operations, process them, and update the read-optimized database.

## Scaffolding

```plaintext
.
├── cmd
│   ├── command
│   │   └── main.go
│   ├── query
│   │   └── main.go
├── internal
│   ├── command
│   │   ├── controller
│   │   │   ├── command_controller.go
│   │   │   └── model.go
│   │   ├── eventpublisher
│   │   │   │── rabbitmq_producer
│   │   │   │    └── rabbitmq_event_publisher.go
│   │   │   └─── event_publisher.go
│   │   ├── eventstore
│   │   │   │── postgres
│   │   │   │    └── postgres_event_store.go
│   │   │   └─── event_store.go
│   │   └── service
│   │       └── command_service.go
│   ├── http
│   │   ├── command
│   │   │   └── routes.go
│   │   └── query
│   │       └── routes.go
│   ├── model
│   │   ├── comment
│   │   │   └── comment_added_event.go
│   │   │   └── comment_deleted_event.go
│   │   │   └── comment_updated_event.go
│   │   ├── post
│   │   │   └── post_added_event.go
│   │   │   └── post_deleted_event.go
│   │   │   └── post_updated_event.go
│   │   ├── reaction
│   │   │   └── reaction_added_event.go
│   │   │   └── reaction_deleted_event.go
│   │   └── event.go
│   ├── query
│   │   ├── controller
│   │   │   ├── query_controller.go
│   │   │   └── model.go
│   │   ├── eventconsumer
│   │   │   │── rabbitmq_producer
│   │   │   │    └── rabbitmq_event_consumer.go
│   │   │   └─── event_consumer.go
│   │   ├── repository
│   │   │   │── mongodb
│   │   │   │    └── mongodb_repository.go
│   │   │   └─── query_repository.go
│   │   └── service
│   │       └── query_service.go
├── migrations
│   └── postgres
│       └── 001_create_events_table.sql
├── pkg
│   │── clock
│   │    └── clock.go
│   └── logger
│       └── logger.go
├── docker-compose.yml
├── Dockerfile.command
├── Dockerfile.query
├── go.mod
└── go.sum
```
### `cmd`

- **`command`**: The entry point for the Command App.
    - **`main.go`**: Initializes and starts the Command App.

- **`query`**: The entry point for the Query App.
    - **`main.go`**: Initializes and starts the Query App.

### `internal`

#### `command`

- **`controller`**: Handles HTTP requests and invokes services.
    - **`command_controller.go`**: Defines controllers for handling post and comment commands.
    - **`model.go`**: Defines data models used in command operations.

- **`eventpublisher`**: Publishes events generated by command operations.
    - **`rabbitmq_producer`**: Implements RabbitMQ as an event publisher.
        - **`rabbitmq_event_publisher.go`**

- **`eventstore`**: Stores events generated by command operations.
    - **`postgres`**: Implements PostgreSQL as the event store.
        - **`postgres_event_store.go`**

- **`service`**: Contains business logic for command operations.
    - **`command_service.go`**: Implements services for handling post and comment commands.

#### `http` Handles HTTP routes

- **`command`**
    - **`routes.go`**: Defines routes for handling post and comment commands.

- **`query`**
    - **`routes.go`**: Defines routes for handling queries.

#### `model`

- **`comment`**: Defines events related to comments.
    - **`comment_added_event.go`**
    - **`comment_deleted_event.go`**
    - **`comment_updated_event.go`**

- **`post`**: Defines events related to posts.
    - **`post_added_event.go`**
    - **`post_deleted_event.go`**
    - **`post_updated_event.go`**

- **`event.go`**: Common event interfaces.

#### `query`

- **`controller`**: Handles HTTP requests for query operations.
    - **`query_controller.go`**: Defines controllers for handling query operations.
    - **`model.go`**: Defines data models used in query operations.

- **`eventconsumer`**: Consumes events for updating the query side.
    - **`rabbitmq_producer`**: Implements RabbitMQ as an event consumer.
        - **`rabbitmq_event_consumer.go`**: Defines the RabbitMQ event consumer.

- **`repository`**: Retrieves data from databases.
    - **`mongodb`**: Implements MongoDB as the query repository.
        - **`mongodb_repository.go`**: Defines the MongoDB query repository.

- **`service`**: Contains business logic for query operations.
    - **`query_service.go`**: Implements services for handling query operations.

### `migrations`

- **`postgres`**: Contains SQL scripts for PostgreSQL migrations.
    - **`001_create_events_table.sql`**: Creates the initial table for storing events.

### `pkg`

## How to Run the Application

To run the application, follow these steps:

1. Clone the code repository:

   ```bash
   git clone <repository_url>
   ```

2. Navigate to the root directory of the cloned repository:

   ```bash
   cd <repository_directory>
   ```

3. Run Docker Compose to start the services:

   ```bash
   docker-compose build
   docker-compose up -d
   ```

This will initiate the command and query applications, along with necessary dependencies like PostgreSQL, RabbitMQ, and MongoDB, as defined in the `docker-compose.yml` file.


### Disclaimer

Please note that there are some missing features and events in this implementation:

* The application lacks proper unit and integration tests and may not cover all edge cases.
* The follow-user events are not implemented at the moment.
* Additionally, events related to reactions are not fully implemented.
* Missing environmental variables for connections, such as credentials and endpoint configurations, are not yet incorporated.

In the pursuit of achieving a functional version, the focus has been primarily on core features, and some aspects had to be prioritized over others.

