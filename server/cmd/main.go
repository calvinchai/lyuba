package main

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/discovery"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	// Create a new libp2p Host with options
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), // Listen on all interfaces
		// Add additional options here
	)
	if err != nil {
		log.Fatalf("Failed to create host: %s", err)
	}

	log.Printf("Host created, listening on %s", host.Addrs())
	service := discovery.NewMdnsService(ctx, host, time.Hour, "")
	service.Start()
}
