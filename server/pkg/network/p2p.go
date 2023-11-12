package network

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p/core/host"
    "github.com/libp2p/go-libp2p/core/network"
)

// createHost creates a libp2p host with a random peer ID.
func createHost(ctx context.Context) (host.Host, error) {
    h, err := libp2p.New(ctx)
    if err != nil {
        return nil, err
    }
    return h, nil
}

// handleStream sets a stream handler. This function is called when a peer initiates a connection and starts a stream with this peer.
func handleStream(stream network.Stream) {
    // Create a buffer stream for non blocking read and write.
    fmt.Println("New stream from", stream.Conn().RemotePeer())
    // Implement stream handling logic here
    stream.Close()
}

// RunP2PNetwork starts a P2P network node.
func RunP2PNetwork() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    h, err := createHost(ctx)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Host created. ID:", h.ID())
    h.SetStreamHandler("/myP2PNetwork/1.0.0", handleStream)

    // Wait for a signal to exit
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    fmt.Println("Received signal, shutting down...")
}

