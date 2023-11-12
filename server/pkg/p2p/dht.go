package p2p

import (
    "context"

    "github.com/libp2p/go-libp2p/core/host"
    "github.com/libp2p/go-libp2p-kad-dht"
)

// SetupDHT sets up a Distributed Hash Table for the peer-to-peer network.
func SetupDHT(ctx context.Context, h host.Host) (*dht.IpfsDHT, error) {
    // Create a new DHT instance. This helps in finding other peers in the network and establishing connections.
    dht, err := dht.New(ctx, h)
    if err != nil {
        return nil, err
    }

    // Bootstrap the DHT. This helps in connecting the DHT to a few well-known peers in the network.
    if err := dht.Bootstrap(ctx); err != nil {
        return nil, err
    }

    return dht, nil
}
