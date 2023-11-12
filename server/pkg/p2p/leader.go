package p2p
import (
	"net"
	"lyubanode/config"
)

func update_dns() {
	BootstrapDomain := config.AppConfig.BootstrapDomain
	// check if nul
	if BootstrapDomain == "" {
		return 
	}

	// resolve and check if match
	// if not match, update
	ips, _ := net.LookupIP(BootstrapDomain)
	// Print the resolved IP addresses
    for _, ip := range ips {
	print(ip.String())
    }
}

func retire(){

}


