/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package sqlv2_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/sql-query-go-sdk/sqlv2"

	"github.com/pact-foundation/pact-go/dsl"
)

var commonHeaders = dsl.MapMatcher{
	"Content-Type": term("application/json; charset=utf-8", `application\/json`),
}

var instance string = "crn:v1:bluemix:public:sql-query:eu-de:a/f5e2ac71094077500e0d4b1ef8b9de0a:2ad2f1b6-6bce-42cd-b7ff-aa397e3afe9b::"

var client *sqlv2.SqlV2

func TestMain(m *testing.M) {
	var exitCode int

	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	exitCode = m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

func TestClientPact_GetUser(t *testing.T) {
	t.Run("tables can be listed", func(t *testing.T) {

		pact.
			AddInteraction().
			Given("Instance has tables").
			UponReceiving("A request to list all tables").
			WithRequest(request{
				Method: "GET",
				Path:   dsl.String("/tables"),
				Query: dsl.MapMatcher{
					"instance_crn": term(instance, "[a-zA-Z]+"),
					"name_pattern": term("pluto", "[a-zA-Z]+"),
					"type":         term("table", "(table|view)"),
				},
			}).
			WillRespondWith(dsl.Response{
				Status:  200,
				Body:    dsl.Match(sqlv2.TableList{}),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			// Construct an instance of the ListTablesOptions model
			listTablesOptionsModel := new(sqlv2.ListTablesOptions)
			listTablesOptionsModel.NamePattern = core.StringPtr("pluto")
			listTablesOptionsModel.Type = core.StringPtr("table")
			listTablesOptionsModel.Headers = map[string]string{}
			// Expect response parsing to fail since we are receiving a text/plain response
			_, _, operationErr := client.ListTables(listTablesOptionsModel)

			if operationErr != nil {
				t.Fatalf("Error on Verify: %v", operationErr)
			}
			return operationErr
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}

// Common test data
var pact dsl.Pact

// Aliases
var term = dsl.Term

type request = dsl.Request

func setup() {
	pact = createPact()

	// Proactively start service to get access to the port
	pact.Setup(true)

	u := fmt.Sprintf("http://localhost:%d", pact.Server.Port)

	var err error
	client, err = sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
		URL:           u,
		Authenticator: &core.NoAuthAuthenticator{},
		InstanceCrn:   core.StringPtr(instance),
	})

	if err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}

}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 os.Getenv("CONSUMER_NAME"),
		Provider:                 os.Getenv("PROVIDER_NAME"),
		LogDir:                   os.Getenv("LOG_DIR"),
		PactDir:                  os.Getenv("PACT_DIR"),
		LogLevel:                 "INFO",
		DisableToolValidityCheck: true,
	}
}
