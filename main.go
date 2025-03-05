package main

import (
	"context"
	"os"

	"github.com/FastLane-Labs/kingdom-solver/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	log.InitLogger("debug")

	rpcUrl := os.Getenv("OPERATIONS_RELAY_URL")
	if rpcUrl == "" {
		log.Error("OPERATIONS_RELAY_URL is not set")
		os.Exit(1)
	}

	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		log.Error("failed to connect to the operations relay", "error", err)
		os.Exit(1)
	}

	ch := make(chan any)

	sub, err := client.Subscribe(context.Background(), "", ch, "user operation")
	if err != nil {
		log.Error("failed to subscribe to the operations relay", "error", err)
		os.Exit(1)
	}

	for {
		select {
		case d := <-ch:
			log.Info("received user operation", "data", d)
		case err := <-sub.Err():
			log.Error("subscription error", "error", err)
			return
		}
	}
}
