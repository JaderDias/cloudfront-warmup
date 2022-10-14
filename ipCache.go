package warmup

import (
	"fmt"
	"net"
	"sync"
)

var ipCache sync.Map //  map[string][]net.IP{}

func getUncachedIps(domainName, pointOfPresence string, netLookup NetLookup) []net.IP {
	popDomain := fmt.Sprintf("%s.%s.cloudfront.net", domainName, pointOfPresence)
	ips, err := netLookup.LookupIP(popDomain)
	if err != nil {
		fmt.Printf("error looking up %s: %s\n", popDomain, err)
	}

	return ips
}

func getIPs(domainName, pointOfPresence string, netLookup NetLookup) []net.IP {
	actual, _ := ipCache.LoadOrStore(pointOfPresence, getUncachedIps(domainName, pointOfPresence, netLookup))
	return actual.([]net.IP)
}
