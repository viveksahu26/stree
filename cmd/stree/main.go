package main

import (
	"log"

	"github.com/viveksahu26/stree/cmd/stree/cli"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		// we don't call os.Exit as Fatalf does both PrintF and os.Exit(1)
		log.Fatalf("error during command execution: %v", err)
	}
}
