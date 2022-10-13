package warmup

import (
	"fmt"
	"net"
)

var ipCache = map[string][]net.IP{}

func getIPs(domainName, pop string, netLookup NetLookup) ([]net.IP, error) {
	if ips, ok := ipCache[pop]; ok {
		return ips, nil
	}

	popDomain := fmt.Sprintf("%s.%s.cloudfront.net", domainName, pop)
	ips, err := netLookup.LookupIP(popDomain)
	if err != nil {
		return nil, fmt.Errorf("error looking up %s: %s", popDomain, err)
	}

	ipCache[pop] = ips
	return ips, nil
}
