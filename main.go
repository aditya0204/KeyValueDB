package main

import (
	"adiDB/server"
	"fmt"
)

func main() {
	s := server.NewServer()
	for {
		select {
		case <-s.QuitChannel:
			fmt.Println("Closing the database.")
			return
		}
	}

}
