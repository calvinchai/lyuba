package config

import (
	"os"
	"net"
)

type Config struct {
	APIKey             string
	MastodonKey        string
	BootstrapNode      []string
	BootstrapDomain    string
	DomainApiKey       string
	ProtocolID         string
	StreamPort         int
	WebPort            int
	P2PPort            int
	PgPort             int
	PgBouncerPort      int
	StorjSatellitePort int
	StorjBucket        string
	StorjNodePort      int
	RendezvousString   string
}

var AppConfig *Config

func init() {
	AppConfig = &Config{
		APIKey:          os.Getenv("MY_APP_API_KEY"),
		MastodonKey:     os.Getenv("MY_APP_MASTODON_KEY"),
		BootstrapNode:   []string{"QmZ4YXJ2YXJvYmJpZ2dlcjpzZWNyZXQ6c2VjcmV0QG1hc3RvZG9uLmNvbQ=="},
		BootstrapDomain: "lyuba.social",
		// rest are all default
		ProtocolID:         "lyubanode/1.0",
		StreamPort:         4000,
		WebPort:            3000,
		P2PPort:            4001,
		PgPort:             5432,
		PgBouncerPort:      6432,
		StorjSatellitePort: 7777,
		StorjBucket:        "lyubanode",
		StorjNodePort:      4444,
		RendezvousString:  "lyubanode",
	}
	bootstrapNode, _ := net.LookupIP(AppConfig.BootstrapDomain)
	if len(bootstrapNode) > 0 {
		AppConfig.BootstrapNode = append(AppConfig.BootstrapNode, bootstrapNode[0].String())
	}
}
	

