package main

import (
	"fmt"

	"github.com/hanle23/pokego"
)

func main() {
	// Initialize the client
	client := pokego.NewClient()

	err := client.Berries("123", "12")
	fmt.Println(err)
}
