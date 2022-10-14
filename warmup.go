package warmup

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func warmup(
	domainName string,
	event events.S3Event,
	pointsOfPresence []string,
	netLookup NetLookup,
	httpClientFactory HttpClientFactory,
) error {
	domain := strings.Split(domainName, ".")
	if len(domain) < 3 {
		return fmt.Errorf("domain name should be in the abcdefgijklm.cloudfront.net format")
	}

	if len(event.Records) == 0 {
		return fmt.Errorf("no records in event")
	}

	for _, record := range event.Records {
		uri := fmt.Sprintf("https://%s/%s", domainName, record.S3.Object.Key)
		log.Printf("uri: %s", uri)
		request(domain[0], uri, pointsOfPresence, netLookup, httpClientFactory)
	}

	return nil
}

func Warmup(domainName string, event events.S3Event) error {
	return warmup(
		domainName,
		event,
		cloudfrontPoPs,
		&MyNetLookup{},
		&MyHttpClientFactory{},
	)
}
