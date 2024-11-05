# Producer-Consumer

### Description

This project demonstrates a producer-consumer pattern in Go, where multiple producers generate random numbers and write them to a shared file, while multiple consumers read and remove the numbers from the file. Additionally, a garbage collector ensures the file does not grow indefinitely by truncating it when it reaches a certain number of lines.

### Features

- **Producers**: Generate random numbers and append them to `resource.txt`.
- **Consumers**: Read and remove the first line from `resource.txt`.
- **Garbage Collector**: Truncates `resource.txt` when it exceeds 100 lines.
- **Concurrency**: Utilizes goroutines and a mutex to manage concurrent access to the shared file.

### Usage

To run the project, execute the following command:

```sh
go build .
```

Then execute:

```sh
./main
```

### Requirements

Go 1.23.1 or later

### Files

`main.go`: Entry point of the application.

`producer.go`: Contains the producer logic.

`consumer.go`: Contains the consumer logic.

`garbage_collector.go`: Contains the garbage collector logic.

`shared.go`: Contains shared resources like the mutex.

`resource.txt`: The shared file used by producers and consumers.
