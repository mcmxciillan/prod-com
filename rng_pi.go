package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func getRandomNumber() (uint32, error) {
	// Open the /dev/random file
	file, err := os.Open("/dev/random")
	if err != nil {
		return 0, fmt.Errorf("failed to open /dev/random: %w", err)
	}
	defer file.Close()

	// Create a buffer to hold the random bytes
	var randomNum uint32

	// Read bytes into the buffer
	err = binary.Read(file, binary.LittleEndian, &randomNum)
	if err != nil {
		return 0, fmt.Errorf("failed to read from /dev/random: %w", err)
	}

	return randomNum, nil
}
