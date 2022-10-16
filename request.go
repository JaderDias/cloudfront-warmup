package warmup

import (
	"fmt"
	"log"
	"runtime"

	"github.com/JaderDias/limiter"
)

func request(
	subdomain,
	uri string,
	pointsOfPresence []string,
	netLookup NetLookup,
	httpClientFactory HttpClientFactory,
) {
	numberOfWorkers := runtime.NumCPU() * 4 // optimal number for my use case
	log.Printf("number of workers: %d", numberOfWorkers)
	limiter.BoundedConcurrency(
		numberOfWorkers,
		len(pointsOfPresence),
		func(i int) {
			pointOfPresence := pointsOfPresence[i]
			ips := getIPs(subdomain, pointOfPresence, netLookup)
			for _, ip := range ips {
				if ip.To4() == nil {
					// AWS Lambda doesn't support IPv6 outgoing connections
					continue
				}

				client := httpClientFactory.Get(ip)
				response, err := client.Get(uri)
				if err != nil {
					fmt.Printf("request for %s ip %s failed: %s\n", pointOfPresence, ip, err)
					continue
				}

				if response.Body != nil {
					response.Body.Close()
				}

				log.Printf(
					"x-amz-cf-pop %s ip %s response status %s x-cache %s ",
					response.Header.Get("x-amz-cf-pop"),
					ip,
					response.Status,
					response.Header.Get("x-cache"),
				)
			}
		},
	)
}
