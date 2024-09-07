package main

import (
	"github.com/hanle23/pokego"
)

func main() {
	// Initialize the client
	client := pokego.NewClient()

	client.Berry("something123")
}
