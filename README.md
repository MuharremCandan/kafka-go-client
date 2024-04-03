# Go Kafka Client Example

This is a simple Go Kafka client example that demonstrates how to produce messages to and consume messages from a Kafka cluster running in Docker.

## Requirements

- Docker
- Go (if you want to run the Go client)

## Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/MuharremCandan/kafka-go-client.git
   ```

2. Navigate to the project directory:
   ```bash
   cd kafka-go-client
   ```

3. Start the Kafka cluster using Docker Compose:
   ```bash
   docker-compose up -d
   ```
4. Install dependicies:
    ```bash
    go get "github.com/confluentinc/confluent-kafka-go/kafka"
    ```

## Running the Go Client

### Producer

To run the Go Kafka producer client, follow these steps:

1. Make sure Kafka cluster is up and running (`docker-compose up -d`).
2. Navigate to the `producer` directory:
   ```bash
   cd producer
   ```
3. Build and run the Go program:
   ```bash
   go build
   ./producer
   ```
4. The producer will start a server on http://localhost:8080/hello:name. Give a name to as a parameter. The message will be sent to the Kafka cluster.

### Consumer

To run the Go Kafka consumer client, follow these steps:

1. Make sure Kafka cluster is up and running (`docker-compose up -d`).
2. Navigate to the `consumer` directory:
   ```bash
   cd consumer
   ```
3. Build and run the Go program:
   ```bash
   go build
   ./consumer
   ```
4. The consumer will start listening for messages on the Kafka cluster. Any messages produced to the Kafka cluster will be consumed and printed by the consumer.

## Cleanup

To stop and remove the Docker containers used for the Kafka cluster:
```bash
docker-compose down
```

## Note

- Make sure to update the Kafka broker address in the Go client code (`producer/main.go` and `consumer/main.go`) if you are not using the default configuration provided in the Docker Compose file.
- You may need to install the required Go dependencies if you haven't already:
  ```bash
  go mod tidy
  ```
