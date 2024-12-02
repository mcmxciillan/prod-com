package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func producer() {
	for {
		// Generate a random number
		randomNumber := rand.Intn(100) // Generates a random number between 0 and 99

		// Lock the mutex before accessing the file
		Mu.Lock()

		// Open or create the file
		file, err := os.OpenFile("./resource.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// Wait for a random small amount of time to simulate work
		waitTime := time.Duration(rand.Intn(5)+1) * time.Second
		fmt.Printf("Producer is working for %s\n", waitTime)
		time.Sleep(waitTime)

		if err != nil {
			fmt.Println("Error opening or creating the file:", err)
			Mu.Unlock()
			return
		}

		// Write the random number to the file
		_, err = file.WriteString(fmt.Sprintf("%d\n", randomNumber))
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			file.Close()
			Mu.Unlock()
			return
		}

		file.Close()
		fmt.Println("Producer wrote:", randomNumber)

		// Unlock the mutex after accessing the file
		Mu.Unlock()

		time.Sleep(2 * time.Second)
	}
}
