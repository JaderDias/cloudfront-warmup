package warmup

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func warmup(
	domainName string,
	event events.S3Event,
	pointsOfPresence []string,
	netLookup NetLookup,
	httpClientFactory HttpClientFactory,
) error {
	if len(event.Records) == 0 {
		return fmt.Errorf("no records in event")
	}

	for _, record := range event.Records {
		uri := fmt.Sprintf("https://%s/%s", domainName, record.S3.Object.Key)
		log.Printf("uri: %s", uri)
		request(domainName, uri, pointsOfPresence, netLookup, httpClientFactory)
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
