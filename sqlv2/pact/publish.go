package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IBM/sql-query-go-sdk/common"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func main() {
	// Publish the Pacts...
	p := dsl.Publisher{}

	fmt.Println("Publishing Pact files to broker", os.Getenv("PACT_DIR"), os.Getenv("PACT_BROKER_URL"))
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/go-sdk-sql-query-sql-service-api.json", os.Getenv("PACT_DIR")))},
		ConsumerVersion: common.Version,
		Tags:            []string{"main"},
		BrokerToken:     os.Getenv("PACT_TOKEN"),
	})

	if err != nil {
		log.Println("ERROR: ", err)
		os.Exit(1)
	}
}
