package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func consumer(id int) {
	for {

		time.Sleep(3 * time.Second)

		// Lock the mutex before accessing the file
		Mu.Lock()

		// Open the file
		file, err := os.OpenFile("./resource.txt", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			Mu.Unlock()
			return
		}

		// Read the first line
		scanner := bufio.NewScanner(file)
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		} else {
			// If the file is empty, unlock the mutex and wait for a random amount of time
			file.Close()
			Mu.Unlock()
			waitTime := time.Duration(rand.Intn(5)+1) * time.Second
			fmt.Printf("Resource is empty. Consumer %d is waiting for %s\n", id, waitTime)
			time.Sleep(waitTime)
			continue
		}

		// Read the rest of the file
		var remainingLines []string
		for scanner.Scan() {
			remainingLines = append(remainingLines, scanner.Text())
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
			file.Close()
			Mu.Unlock()
			return
		}

		// Truncate the file and write back the remaining lines
		err = file.Truncate(0)
		if err != nil {
			fmt.Println("Error truncating the file:", err)
			file.Close()
			Mu.Unlock()
			return
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Println("Error seeking to the beginning of the file:", err)
			file.Close()
			Mu.Unlock()
			return
		}
		writer := bufio.NewWriter(file)
		for _, line := range remainingLines {
			_, err = writer.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to the file:", err)
				file.Close()
				Mu.Unlock()
				return
			}
		}
		writer.Flush()

		file.Close()

		// Print the first line
		fmt.Printf("Consumer %d read and removed: %s\n", id, firstLine)

		// Unlock the mutex after accessing the file
		Mu.Unlock()
	}
}
