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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sql-query-go-sdk/sqlv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SqlV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		instanceCrn := "testString"
		It(`Instantiate service client`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				InstanceCrn: core.StringPtr(instanceCrn),
			})
			Expect(sqlService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				URL: "{BAD_URL_STRING",
				InstanceCrn: core.StringPtr(instanceCrn),
			})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				URL: "https://sqlv2/api",
				InstanceCrn: core.StringPtr(instanceCrn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		instanceCrn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_URL": "https://sqlv2/api",
				"SQL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					URL: "https://testService/api",
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sqlService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				err := sqlService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sqlService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_URL": "https://sqlv2/api",
				"SQL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
				InstanceCrn: core.StringPtr(instanceCrn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(sqlService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
				URL: "{BAD_URL_STRING",
				InstanceCrn: core.StringPtr(instanceCrn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(sqlService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = sqlv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListTables(listTablesOptions *ListTablesOptions) - Operation response error`, func() {
		instanceCrn := "testString"
		listTablesPath := "/tables"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name_pattern"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTables with error: Operation response processing error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(sqlv2.ListTablesOptions)
				listTablesOptionsModel.NamePattern = core.StringPtr("testString")
				listTablesOptionsModel.Type = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sqlService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sqlService.EnableRetries(0, 0)
				result, response, operationErr = sqlService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTables(listTablesOptions *ListTablesOptions)`, func() {
		instanceCrn := "testString"
		listTablesPath := "/tables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name_pattern"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tables": ["Tables"], "tables_metadata": [{"name": "customer_address", "type": "table"}]}`)
				}))
			})
			It(`Invoke ListTables successfully with retries`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				sqlService.EnableRetries(0, 0)

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(sqlv2.ListTablesOptions)
				listTablesOptionsModel.NamePattern = core.StringPtr("testString")
				listTablesOptionsModel.Type = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sqlService.ListTablesWithContext(ctx, listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sqlService.DisableRetries()
				result, response, operationErr := sqlService.ListTables(listTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sqlService.ListTablesWithContext(ctx, listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name_pattern"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tables": ["Tables"], "tables_metadata": [{"name": "customer_address", "type": "table"}]}`)
				}))
			})
			It(`Invoke ListTables successfully`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sqlService.ListTables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(sqlv2.ListTablesOptions)
				listTablesOptionsModel.NamePattern = core.StringPtr("testString")
				listTablesOptionsModel.Type = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sqlService.ListTables(listTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTables with error: Operation request error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(sqlv2.ListTablesOptions)
				listTablesOptionsModel.NamePattern = core.StringPtr("testString")
				listTablesOptionsModel.Type = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sqlService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sqlService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTable(getTableOptions *GetTableOptions) - Operation response error`, func() {
		instanceCrn := "testString"
		getTablePath := "/tables/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTable with error: Operation response processing error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(sqlv2.GetTableOptions)
				getTableOptionsModel.TableName = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sqlService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sqlService.EnableRetries(0, 0)
				result, response, operationErr = sqlService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTable(getTableOptions *GetTableOptions)`, func() {
		instanceCrn := "testString"
		getTablePath := "/tables/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "customer_address", "type": "table", "columns": [{"name": "Name", "type": "string", "nullable": true}]}`)
				}))
			})
			It(`Invoke GetTable successfully with retries`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				sqlService.EnableRetries(0, 0)

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(sqlv2.GetTableOptions)
				getTableOptionsModel.TableName = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sqlService.GetTableWithContext(ctx, getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sqlService.DisableRetries()
				result, response, operationErr := sqlService.GetTable(getTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sqlService.GetTableWithContext(ctx, getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "customer_address", "type": "table", "columns": [{"name": "Name", "type": "string", "nullable": true}]}`)
				}))
			})
			It(`Invoke GetTable successfully`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sqlService.GetTable(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(sqlv2.GetTableOptions)
				getTableOptionsModel.TableName = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sqlService.GetTable(getTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTable with error: Operation validation and request error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(sqlv2.GetTableOptions)
				getTableOptionsModel.TableName = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sqlService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sqlService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTableOptions model with no property values
				getTableOptionsModelNew := new(sqlv2.GetTableOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sqlService.GetTable(getTableOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		instanceCrn := "testString"
		It(`Instantiate service client`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				InstanceCrn: core.StringPtr(instanceCrn),
			})
			Expect(sqlService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				URL: "{BAD_URL_STRING",
				InstanceCrn: core.StringPtr(instanceCrn),
			})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				URL: "https://sqlv2/api",
				InstanceCrn: core.StringPtr(instanceCrn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{})
			Expect(sqlService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		instanceCrn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_URL": "https://sqlv2/api",
				"SQL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					URL: "https://testService/api",
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sqlService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				err := sqlService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sqlService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sqlService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sqlService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sqlService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sqlService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_URL": "https://sqlv2/api",
				"SQL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
				InstanceCrn: core.StringPtr(instanceCrn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(sqlService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SQL_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			sqlService, serviceErr := sqlv2.NewSqlV2UsingExternalConfig(&sqlv2.SqlV2Options{
				URL: "{BAD_URL_STRING",
				InstanceCrn: core.StringPtr(instanceCrn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(sqlService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = sqlv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`SubmitSqlJob(submitSqlJobOptions *SubmitSqlJobOptions) - Operation response error`, func() {
		instanceCrn := "testString"
		submitSqlJobPath := "/sql_jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(submitSqlJobPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SubmitSqlJob with error: Operation response processing error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the SubmitSqlJobOptions model
				submitSqlJobOptionsModel := new(sqlv2.SubmitSqlJobOptions)
				submitSqlJobOptionsModel.Statement = core.StringPtr("testString")
				submitSqlJobOptionsModel.ResultsetTarget = core.StringPtr("testString")
				submitSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sqlService.SubmitSqlJob(submitSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sqlService.EnableRetries(0, 0)
				result, response, operationErr = sqlService.SubmitSqlJob(submitSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SubmitSqlJob(submitSqlJobOptions *SubmitSqlJobOptions)`, func() {
		instanceCrn := "testString"
		submitSqlJobPath := "/sql_jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(submitSqlJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "has_hints": true}`)
				}))
			})
			It(`Invoke SubmitSqlJob successfully with retries`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				sqlService.EnableRetries(0, 0)

				// Construct an instance of the SubmitSqlJobOptions model
				submitSqlJobOptionsModel := new(sqlv2.SubmitSqlJobOptions)
				submitSqlJobOptionsModel.Statement = core.StringPtr("testString")
				submitSqlJobOptionsModel.ResultsetTarget = core.StringPtr("testString")
				submitSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sqlService.SubmitSqlJobWithContext(ctx, submitSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sqlService.DisableRetries()
				result, response, operationErr := sqlService.SubmitSqlJob(submitSqlJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sqlService.SubmitSqlJobWithContext(ctx, submitSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(submitSqlJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "has_hints": true}`)
				}))
			})
			It(`Invoke SubmitSqlJob successfully`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sqlService.SubmitSqlJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubmitSqlJobOptions model
				submitSqlJobOptionsModel := new(sqlv2.SubmitSqlJobOptions)
				submitSqlJobOptionsModel.Statement = core.StringPtr("testString")
				submitSqlJobOptionsModel.ResultsetTarget = core.StringPtr("testString")
				submitSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sqlService.SubmitSqlJob(submitSqlJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SubmitSqlJob with error: Operation validation and request error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the SubmitSqlJobOptions model
				submitSqlJobOptionsModel := new(sqlv2.SubmitSqlJobOptions)
				submitSqlJobOptionsModel.Statement = core.StringPtr("testString")
				submitSqlJobOptionsModel.ResultsetTarget = core.StringPtr("testString")
				submitSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sqlService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sqlService.SubmitSqlJob(submitSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SubmitSqlJobOptions model with no property values
				submitSqlJobOptionsModelNew := new(sqlv2.SubmitSqlJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sqlService.SubmitSqlJob(submitSqlJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSqlJobs(listSqlJobsOptions *ListSqlJobsOptions) - Operation response error`, func() {
		instanceCrn := "testString"
		listSqlJobsPath := "/sql_jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSqlJobsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSqlJobs with error: Operation response processing error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the ListSqlJobsOptions model
				listSqlJobsOptionsModel := new(sqlv2.ListSqlJobsOptions)
				listSqlJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sqlService.ListSqlJobs(listSqlJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sqlService.EnableRetries(0, 0)
				result, response, operationErr = sqlService.ListSqlJobs(listSqlJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListSqlJobs(listSqlJobsOptions *ListSqlJobsOptions)`, func() {
		instanceCrn := "testString"
		listSqlJobsPath := "/sql_jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSqlJobsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"jobs": [{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "has_hints": true}]}`)
				}))
			})
			It(`Invoke ListSqlJobs successfully with retries`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				sqlService.EnableRetries(0, 0)

				// Construct an instance of the ListSqlJobsOptions model
				listSqlJobsOptionsModel := new(sqlv2.ListSqlJobsOptions)
				listSqlJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sqlService.ListSqlJobsWithContext(ctx, listSqlJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sqlService.DisableRetries()
				result, response, operationErr := sqlService.ListSqlJobs(listSqlJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sqlService.ListSqlJobsWithContext(ctx, listSqlJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSqlJobsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"jobs": [{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "has_hints": true}]}`)
				}))
			})
			It(`Invoke ListSqlJobs successfully`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sqlService.ListSqlJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSqlJobsOptions model
				listSqlJobsOptionsModel := new(sqlv2.ListSqlJobsOptions)
				listSqlJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sqlService.ListSqlJobs(listSqlJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSqlJobs with error: Operation request error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the ListSqlJobsOptions model
				listSqlJobsOptionsModel := new(sqlv2.ListSqlJobsOptions)
				listSqlJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sqlService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sqlService.ListSqlJobs(listSqlJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSqlJob(getSqlJobOptions *GetSqlJobOptions) - Operation response error`, func() {
		instanceCrn := "testString"
		getSqlJobPath := "/sql_jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSqlJobPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSqlJob with error: Operation response processing error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the GetSqlJobOptions model
				getSqlJobOptionsModel := new(sqlv2.GetSqlJobOptions)
				getSqlJobOptionsModel.JobID = core.StringPtr("testString")
				getSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sqlService.GetSqlJob(getSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sqlService.EnableRetries(0, 0)
				result, response, operationErr = sqlService.GetSqlJob(getSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSqlJob(getSqlJobOptions *GetSqlJobOptions)`, func() {
		instanceCrn := "testString"
		getSqlJobPath := "/sql_jobs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSqlJobPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "statement": "Statement", "plan_id": "PlanID", "resultset_format": "csv", "resultset_location": "cos://s3.dal.us.cloud-object-storage.appdomain.cloud/target-bucket/q1-results/jobid=411323a4-04de-440a-9e41-011d31052f54", "end_time": "2019-01-01T12:00:00", "rows_returned": 12, "rows_read": 8, "bytes_read": 9, "objects_skipped": 14, "objects_qualified": 16, "error": "Error", "error_message": "ErrorMessage", "hints": ["Hints"]}`)
				}))
			})
			It(`Invoke GetSqlJob successfully with retries`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())
				sqlService.EnableRetries(0, 0)

				// Construct an instance of the GetSqlJobOptions model
				getSqlJobOptionsModel := new(sqlv2.GetSqlJobOptions)
				getSqlJobOptionsModel.JobID = core.StringPtr("testString")
				getSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sqlService.GetSqlJobWithContext(ctx, getSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sqlService.DisableRetries()
				result, response, operationErr := sqlService.GetSqlJob(getSqlJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sqlService.GetSqlJobWithContext(ctx, getSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSqlJobPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["instance_crn"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"job_id": "637a0b7d-069d-453f-a418-a35d4db3ea64", "status": "queued", "user_id": "test_user@my.org", "submit_time": "2019-01-01T12:00:00", "statement": "Statement", "plan_id": "PlanID", "resultset_format": "csv", "resultset_location": "cos://s3.dal.us.cloud-object-storage.appdomain.cloud/target-bucket/q1-results/jobid=411323a4-04de-440a-9e41-011d31052f54", "end_time": "2019-01-01T12:00:00", "rows_returned": 12, "rows_read": 8, "bytes_read": 9, "objects_skipped": 14, "objects_qualified": 16, "error": "Error", "error_message": "ErrorMessage", "hints": ["Hints"]}`)
				}))
			})
			It(`Invoke GetSqlJob successfully`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sqlService.GetSqlJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSqlJobOptions model
				getSqlJobOptionsModel := new(sqlv2.GetSqlJobOptions)
				getSqlJobOptionsModel.JobID = core.StringPtr("testString")
				getSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sqlService.GetSqlJob(getSqlJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSqlJob with error: Operation validation and request error`, func() {
				sqlService, serviceErr := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					InstanceCrn: core.StringPtr(instanceCrn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(sqlService).ToNot(BeNil())

				// Construct an instance of the GetSqlJobOptions model
				getSqlJobOptionsModel := new(sqlv2.GetSqlJobOptions)
				getSqlJobOptionsModel.JobID = core.StringPtr("testString")
				getSqlJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sqlService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sqlService.GetSqlJob(getSqlJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSqlJobOptions model with no property values
				getSqlJobOptionsModelNew := new(sqlv2.GetSqlJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sqlService.GetSqlJob(getSqlJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			instanceCrn := "testString"
			sqlService, _ := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
				URL:           "http://sqlv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				InstanceCrn: core.StringPtr(instanceCrn),
			})
			It(`Invoke NewGetSqlJobOptions successfully`, func() {
				// Construct an instance of the GetSqlJobOptions model
				jobID := "testString"
				getSqlJobOptionsModel := sqlService.NewGetSqlJobOptions(jobID)
				getSqlJobOptionsModel.SetJobID("testString")
				getSqlJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSqlJobOptionsModel).ToNot(BeNil())
				Expect(getSqlJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(getSqlJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTableOptions successfully`, func() {
				// Construct an instance of the GetTableOptions model
				tableName := "testString"
				getTableOptionsModel := sqlService.NewGetTableOptions(tableName)
				getTableOptionsModel.SetTableName("testString")
				getTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTableOptionsModel).ToNot(BeNil())
				Expect(getTableOptionsModel.TableName).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSqlJobsOptions successfully`, func() {
				// Construct an instance of the ListSqlJobsOptions model
				listSqlJobsOptionsModel := sqlService.NewListSqlJobsOptions()
				listSqlJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSqlJobsOptionsModel).ToNot(BeNil())
				Expect(listSqlJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTablesOptions successfully`, func() {
				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := sqlService.NewListTablesOptions()
				listTablesOptionsModel.SetNamePattern("testString")
				listTablesOptionsModel.SetType("testString")
				listTablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTablesOptionsModel).ToNot(BeNil())
				Expect(listTablesOptionsModel.NamePattern).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSubmitSqlJobOptions successfully`, func() {
				// Construct an instance of the SubmitSqlJobOptions model
				submitSqlJobOptionsStatement := "testString"
				submitSqlJobOptionsModel := sqlService.NewSubmitSqlJobOptions(submitSqlJobOptionsStatement)
				submitSqlJobOptionsModel.SetStatement("testString")
				submitSqlJobOptionsModel.SetResultsetTarget("testString")
				submitSqlJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(submitSqlJobOptionsModel).ToNot(BeNil())
				Expect(submitSqlJobOptionsModel.Statement).To(Equal(core.StringPtr("testString")))
				Expect(submitSqlJobOptionsModel.ResultsetTarget).To(Equal(core.StringPtr("testString")))
				Expect(submitSqlJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
