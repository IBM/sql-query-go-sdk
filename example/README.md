# SQL Query Go SDK Example

## Running example.go

To run the example, run the following commands from this directory:
1. `export IAM_APIKEY=<Your IBM Cloud API key>`
2. `export SQL_QUERY_CRN=<Your SQL Query instance CRN>`
3. `export SQL_QUERY_TARGET=<Taget COS location for query result(e.g. cos://us-geo/<bucket>/<prefix>)>`
4. `go run example.go`

## How-to

### Set up an authenticator
```go
authenticator := &core.IamAuthenticator{
    ApiKey:       os.Getenv("IAM_APIKEY"),
}
```

### Set up a SqlQuery client
```go
sqlService, err := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
	InstanceCrn:   &instance,
	Authenticator: authenticator,
})
```

### Use the SqlQuery client to submit a SQL query
```go
submitResult, _, submitRrr := sqlService.SubmitSqlJob(&sqlv2.SubmitSqlJobOptions{
	Statement: core.StringPtr("SELECT * FROM cos://us-geo/sql/customers.csv"),
	ResultsetTarget: core.StringPtr(target),
})
```
