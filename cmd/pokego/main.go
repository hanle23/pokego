package main

import (
	"fmt"

	"github.com/hanle23/pokego"
)

func main() {
	// Initialize the client
	client := pokego.NewClient()
	fmt.Println(client.GetCurrentConfig())
	client.Berries("0", "20")
}
