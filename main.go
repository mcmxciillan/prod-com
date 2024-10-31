package main

func main() {
	go producer()
	go consumer(1)
	go consumer(2)
	go consumer(3)
	go garbageCollector()

	// Keep the main function running
	select {}
}
