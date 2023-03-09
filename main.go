package main

import (
	"os"
)

func main() {
}

func readFiles(ch <-chan string) error {
	for path := range ch {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
