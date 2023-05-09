package main

import (
	"fmt"
	"os"

	"github.com/mrbryside/pkg/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
