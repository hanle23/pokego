package main

import (
	"fmt"

	"github.com/hanle23/pokego"
)

func main() {
	// Initialize the client
	client := pokego.NewClient(pokego.WithUseCache(false))
	fmt.Println(client.GetCurrentConfig())
}
