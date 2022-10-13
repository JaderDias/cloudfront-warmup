package warmup

import (
	"fmt"
	"log"
	"runtime"

	"github.com/JaderDias/limiter"
)

func request(
	domainName,
	uri string,
	pointsOfPresence []string,
	netLookup NetLookup,
	httpClientFactory HttpClientFactory,
) {
	numberOfWorkers := runtime.NumCPU() * 4 // optimal number for my use case
	limiter.BoundedConcurrency(
		numberOfWorkers,
		len(pointsOfPresence),
		func(i int) {
			pointOfPresence := pointsOfPresence[i]
			log.Printf("Requesting %s from %s", uri, pointOfPresence)
			ips, err := getIPs(domainName, pointOfPresence, netLookup)
			if err != nil {
				fmt.Printf("error looking up %s: %v", pointOfPresence, err)
			}

			for _, ip := range ips {
				if ip.To4() == nil {
					// AWS Lambda doesn't support IPv6 outgoing connections
					continue
				}

				client := httpClientFactory.Get(ip)
				response, err := client.Get(uri)
				if err != nil {
					fmt.Printf("request for %s ip %s failed: %s", pointOfPresence, ip, err)
				}

				log.Printf(
					"response for %s status %s x-cache %s x-amz-cf-pop %s",
					ip,
					response.Status,
					response.Header.Get("x-cache"),
					response.Header.Get("x-amz-cf-pop"),
				)
			}
		},
	)
}