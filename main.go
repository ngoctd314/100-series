package main

import (
	"os"
)

func main() {
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
