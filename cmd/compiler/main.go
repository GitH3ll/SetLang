package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panicf("failed to open source file: %s", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var buf []byte
	_, err = file.Read(buf)
	if err != nil {
		log.Panicf("failed to read file: %s", err)
	}
}
