package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func garbageCollector() {
	for {
		time.Sleep(10 * time.Second)
		// Lock the mutex before accessing the file
		Mu.Lock()

		// Open the file
		file, err := os.OpenFile("resource.txt", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			Mu.Unlock()
			return
		}

		// Count the number of lines in the file
		lineCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineCount++
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
			file.Close()
			Mu.Unlock()
			return
		}

		// If the file has 100 lines, truncate the file
		if lineCount >= 100 {
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
			fmt.Println("Garbage collector truncated the file")
		}

		file.Close()

		// Unlock the mutex after accessing the file
		Mu.Unlock()

	}
}
