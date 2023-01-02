package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dusk-network/dusk-protobuf/autogen/go/node"
	"github.com/dusk-network/dusk-wallet-cli/prompt"
)

func main() {
	conf := initConfig()

	// Establish a gRPC connection with the node.
	client := newNodeClient()

	if err := client.Connect(conf.RPC); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Defer the client.Close() call using the WaitGroup
	defer func() {
		client.Close()
		wg.Done()
	}()

	// Inquire node about its wallet state, so we know which menu to open.
	resp, err := client.c.GetWalletStatus(context.Background(), &node.EmptyRequest{})
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	wg.Wait() // Wait for the deferred tasks to complete before exiting
	}
}
