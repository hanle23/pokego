package main

import (
	"fmt"

	"github.com/hanle23/pokego"
)

func main() {
	// Initialize the client
	client := pokego.NewClient()

	result, err := client.Berries("0", "20")
	fmt.Println(err)
	fmt.Println(result)
}
