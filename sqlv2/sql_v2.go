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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.25.0-2b3f843a-20210115-164628
 */

// Package sqlv2 : Operations and models for the SqlV2 service
package sqlv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.com/IBM/sql-query-go-sdk/common"
)

// SqlV2 : SQL Query is a stateless service for analyzing rectangular data stored in IBM Cloud Object Store using ANSI
// SQL.
//
// Version: 2.0
type SqlV2 struct {
	Service *core.BaseService

	// The cloud resource name (CRN) of the SQL query service instance.
	InstanceCrn *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.sql-query.cloud.ibm.com/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "sql"

// SqlV2Options : Service options
type SqlV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// The cloud resource name (CRN) of the SQL query service instance.
	InstanceCrn *string `validate:"required"`
}

// NewSqlV2UsingExternalConfig : constructs an instance of SqlV2 with passed in options and external configuration.
func NewSqlV2UsingExternalConfig(options *SqlV2Options) (sql *SqlV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	sql, err = NewSqlV2(options)
	if err != nil {
		return
	}

	err = sql.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = sql.Service.SetServiceURL(options.URL)
	}
	return
}

// NewSqlV2 : constructs an instance of SqlV2 with passed in options.
func NewSqlV2(options *SqlV2Options) (service *SqlV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &SqlV2{
		Service:     baseService,
		InstanceCrn: options.InstanceCrn,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "sql" suitable for processing requests.
func (sql *SqlV2) Clone() *SqlV2 {
	if core.IsNil(sql) {
		return nil
	}
	clone := *sql
	clone.Service = sql.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (sql *SqlV2) SetServiceURL(url string) error {
	return sql.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (sql *SqlV2) GetServiceURL() string {
	return sql.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (sql *SqlV2) SetDefaultHeaders(headers http.Header) {
	sql.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (sql *SqlV2) SetEnableGzipCompression(enableGzip bool) {
	sql.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (sql *SqlV2) GetEnableGzipCompression() bool {
	return sql.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (sql *SqlV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	sql.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (sql *SqlV2) DisableRetries() {
	sql.Service.DisableRetries()
}

// ListTables : List catalog tables
// Retrieve a list of the first 100 tables that are defined for the given instance in the catalog.
func (sql *SqlV2) ListTables(listTablesOptions *ListTablesOptions) (result *TableList, response *core.DetailedResponse, err error) {
	return sql.ListTablesWithContext(context.Background(), listTablesOptions)
}

// ListTablesWithContext is an alternate form of the ListTables method which supports a Context parameter
func (sql *SqlV2) ListTablesWithContext(ctx context.Context, listTablesOptions *ListTablesOptions) (result *TableList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTablesOptions, "listTablesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sql.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sql.Service.Options.URL, `/tables`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sql", "V2", "ListTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("instance_crn", fmt.Sprint(*sql.InstanceCrn))
	if listTablesOptions.NamePattern != nil {
		builder.AddQuery("name_pattern", fmt.Sprint(*listTablesOptions.NamePattern))
	}
	if listTablesOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listTablesOptions.Type))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sql.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetTable : Get information about a specific catalog table
// Returns information about the specified catalog table.
func (sql *SqlV2) GetTable(getTableOptions *GetTableOptions) (result *TableInformation, response *core.DetailedResponse, err error) {
	return sql.GetTableWithContext(context.Background(), getTableOptions)
}

// GetTableWithContext is an alternate form of the GetTable method which supports a Context parameter
func (sql *SqlV2) GetTableWithContext(ctx context.Context, getTableOptions *GetTableOptions) (result *TableInformation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTableOptions, "getTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTableOptions, "getTableOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"table_name": *getTableOptions.TableName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sql.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sql.Service.Options.URL, `/tables/{table_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sql", "V2", "GetTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("instance_crn", fmt.Sprint(*sql.InstanceCrn))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sql.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableInformation)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SubmitSqlJob : Run an SQL job
// Runs an SQL job using the *IBM Cloud SQL Query* service and stores the result in a new CSV data set in *IBM Cloud
// Object Storage*. The `FROM` clause references rectangular data that is stored in a Parquet, CSV, or JSON file in *IBM
// Cloud Object Storage*. Click <a
// href="https://console.bluemix.net/docs/services/sql-query/sql-query.html#overview">here</a> for more information.
func (sql *SqlV2) SubmitSqlJob(submitSqlJobOptions *SubmitSqlJobOptions) (result *SqlJobInfoShort, response *core.DetailedResponse, err error) {
	return sql.SubmitSqlJobWithContext(context.Background(), submitSqlJobOptions)
}

// SubmitSqlJobWithContext is an alternate form of the SubmitSqlJob method which supports a Context parameter
func (sql *SqlV2) SubmitSqlJobWithContext(ctx context.Context, submitSqlJobOptions *SubmitSqlJobOptions) (result *SqlJobInfoShort, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(submitSqlJobOptions, "submitSqlJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(submitSqlJobOptions, "submitSqlJobOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sql.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sql.Service.Options.URL, `/sql_jobs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range submitSqlJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sql", "V2", "SubmitSqlJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("instance_crn", fmt.Sprint(*sql.InstanceCrn))

	body := make(map[string]interface{})
	if submitSqlJobOptions.Statement != nil {
		body["statement"] = submitSqlJobOptions.Statement
	}
	if submitSqlJobOptions.ResultsetTarget != nil {
		body["resultset_target"] = submitSqlJobOptions.ResultsetTarget
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sql.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSqlJobInfoShort)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListSqlJobs : Get information about recent SQL jobs
// Returns information about recently submitted SQL jobs.
func (sql *SqlV2) ListSqlJobs(listSqlJobsOptions *ListSqlJobsOptions) (result *SqlJobInfoList, response *core.DetailedResponse, err error) {
	return sql.ListSqlJobsWithContext(context.Background(), listSqlJobsOptions)
}

// ListSqlJobsWithContext is an alternate form of the ListSqlJobs method which supports a Context parameter
func (sql *SqlV2) ListSqlJobsWithContext(ctx context.Context, listSqlJobsOptions *ListSqlJobsOptions) (result *SqlJobInfoList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSqlJobsOptions, "listSqlJobsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sql.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sql.Service.Options.URL, `/sql_jobs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSqlJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sql", "V2", "ListSqlJobs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("instance_crn", fmt.Sprint(*sql.InstanceCrn))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sql.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSqlJobInfoList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSqlJob : Get information about a specific SQL job
// Returns information about the specified SQL job.
func (sql *SqlV2) GetSqlJob(getSqlJobOptions *GetSqlJobOptions) (result *SqlJobInfoFull, response *core.DetailedResponse, err error) {
	return sql.GetSqlJobWithContext(context.Background(), getSqlJobOptions)
}

// GetSqlJobWithContext is an alternate form of the GetSqlJob method which supports a Context parameter
func (sql *SqlV2) GetSqlJobWithContext(ctx context.Context, getSqlJobOptions *GetSqlJobOptions) (result *SqlJobInfoFull, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSqlJobOptions, "getSqlJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSqlJobOptions, "getSqlJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_id": *getSqlJobOptions.JobID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sql.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sql.Service.Options.URL, `/sql_jobs/{job_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSqlJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("sql", "V2", "GetSqlJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("instance_crn", fmt.Sprint(*sql.InstanceCrn))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sql.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSqlJobInfoFull)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSqlJobOptions : The GetSqlJob options.
type GetSqlJobOptions struct {
	// ID of the SQL job for which information is to be retrieved. This ID is returned when an SQL job is submitted, and
	// when information about recently submitted SQL jobs is requested.
	JobID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSqlJobOptions : Instantiate GetSqlJobOptions
func (*SqlV2) NewGetSqlJobOptions(jobID string) *GetSqlJobOptions {
	return &GetSqlJobOptions{
		JobID: core.StringPtr(jobID),
	}
}

// SetJobID : Allow user to set JobID
func (options *GetSqlJobOptions) SetJobID(jobID string) *GetSqlJobOptions {
	options.JobID = core.StringPtr(jobID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSqlJobOptions) SetHeaders(param map[string]string) *GetSqlJobOptions {
	options.Headers = param
	return options
}

// GetTableOptions : The GetTable options.
type GetTableOptions struct {
	// Name of the catalog table for which information is to be retrieved. Table names are case-insensitive and must only
	// contain alphabetic and numeral characters, and underscore (_).
	TableName *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTableOptions : Instantiate GetTableOptions
func (*SqlV2) NewGetTableOptions(tableName string) *GetTableOptions {
	return &GetTableOptions{
		TableName: core.StringPtr(tableName),
	}
}

// SetTableName : Allow user to set TableName
func (options *GetTableOptions) SetTableName(tableName string) *GetTableOptions {
	options.TableName = core.StringPtr(tableName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTableOptions) SetHeaders(param map[string]string) *GetTableOptions {
	options.Headers = param
	return options
}

// ListSqlJobsOptions : The ListSqlJobs options.
type ListSqlJobsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSqlJobsOptions : Instantiate ListSqlJobsOptions
func (*SqlV2) NewListSqlJobsOptions() *ListSqlJobsOptions {
	return &ListSqlJobsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListSqlJobsOptions) SetHeaders(param map[string]string) *ListSqlJobsOptions {
	options.Headers = param
	return options
}

// ListTablesOptions : The ListTables options.
type ListTablesOptions struct {
	// A table name pattern for filtering the tables that should be listed. The pattern follows Hive syntax conventions and
	// can include asterisks as wildcards and vertical bars to separate alternatives.
	NamePattern *string

	// A table type for filtering the tables that should be listed, can be "table" or "view".
	Type *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTablesOptions : Instantiate ListTablesOptions
func (*SqlV2) NewListTablesOptions() *ListTablesOptions {
	return &ListTablesOptions{}
}

// SetNamePattern : Allow user to set NamePattern
func (options *ListTablesOptions) SetNamePattern(namePattern string) *ListTablesOptions {
	options.NamePattern = core.StringPtr(namePattern)
	return options
}

// SetType : Allow user to set Type
func (options *ListTablesOptions) SetType(typeVar string) *ListTablesOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTablesOptions) SetHeaders(param map[string]string) *ListTablesOptions {
	options.Headers = param
	return options
}

// SubmitSqlJobOptions : The SubmitSqlJob options.
type SubmitSqlJobOptions struct {
	// This is the SQL query to be submitted. The table names specified in its FROM clause must correspond to files in one
	// or more *IBM Cloud Object Storage* instances. The INTO clause of the query indicates the endpoint, bucket, and
	// (optionally) subfolder in *IBM Cloud Object Storage* in which the query result is to be stored. Within the specified
	// target, each result is stored in a separate subfolder with a name that indicates the job ID. The job ID of a query
	// is returned by the GET endpoint.
	Statement *string `validate:"required"`

	// This field provides an alternative way to specify the target URI for a query. It is supported to preserve backward
	// compatibility and will be removed in a future API version. Use the INTO clause of the SQL query to specify the
	// target URI instead.
	ResultsetTarget *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubmitSqlJobOptions : Instantiate SubmitSqlJobOptions
func (*SqlV2) NewSubmitSqlJobOptions(statement string) *SubmitSqlJobOptions {
	return &SubmitSqlJobOptions{
		Statement: core.StringPtr(statement),
	}
}

// SetStatement : Allow user to set Statement
func (options *SubmitSqlJobOptions) SetStatement(statement string) *SubmitSqlJobOptions {
	options.Statement = core.StringPtr(statement)
	return options
}

// SetResultsetTarget : Allow user to set ResultsetTarget
func (options *SubmitSqlJobOptions) SetResultsetTarget(resultsetTarget string) *SubmitSqlJobOptions {
	options.ResultsetTarget = core.StringPtr(resultsetTarget)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SubmitSqlJobOptions) SetHeaders(param map[string]string) *SubmitSqlJobOptions {
	options.Headers = param
	return options
}

// ColumnInformation : Information about a table column.
type ColumnInformation struct {
	// The name of the column.
	Name *string `json:"name" validate:"required"`

	// The data type of the column.
	Type *string `json:"type" validate:"required"`

	// Whether the column may contain NULL values.
	Nullable *bool `json:"nullable,omitempty"`
}

// UnmarshalColumnInformation unmarshals an instance of ColumnInformation from the specified map of raw messages.
func UnmarshalColumnInformation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ColumnInformation)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nullable", &obj.Nullable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SqlJobInfoFull : Full information about an SQL job, including output or error information.
type SqlJobInfoFull struct {
	// Identifier for an SQL job.
	JobID *string `json:"job_id" validate:"required"`

	// Execution status of an SQL job.
	Status *string `json:"status" validate:"required"`

	// ID of the user who submitted an SQL job.
	UserID *string `json:"user_id" validate:"required"`

	// Timestamp indicating when an SQL job was accepted by the service.
	SubmitTime *strfmt.DateTime `json:"submit_time" validate:"required"`

	// The SQL query that the job processes.
	Statement *string `json:"statement" validate:"required"`

	// The service plan id of the instance.
	PlanID *string `json:"plan_id,omitempty"`

	// Format of the query result.
	ResultsetFormat *string `json:"resultset_format,omitempty"`

	// A URI that indicates where the query result is stored. This URI can be used as input for another SQL query. The
	// result comprises all objects that have a name with this URI as its prefix.
	ResultsetLocation *string `json:"resultset_location,omitempty"`

	// Timestamp indicating when the job finished processing.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// Number of rows returned.
	RowsReturned *float64 `json:"rows_returned,omitempty"`

	// Number of rows read.
	RowsRead *float64 `json:"rows_read,omitempty"`

	// Number of bytes read.
	BytesRead *float64 `json:"bytes_read,omitempty"`

	// Number of objects skipped using index management.(Optional).
	ObjectsSkipped *float64 `json:"objects_skipped,omitempty"`

	// Number of objects qualified using index management.(Optional).
	ObjectsQualified *float64 `json:"objects_qualified,omitempty"`

	// An error that was encountered while processing the job.
	Error *string `json:"error,omitempty"`

	// Detailed information about the error.
	ErrorMessage *string `json:"error_message,omitempty"`

	// Suggests possible optimizations for a query.
	Hints []string `json:"hints,omitempty"`
}

// Constants associated with the SqlJobInfoFull.Status property.
// Execution status of an SQL job.
const (
	SqlJobInfoFull_Status_Completed = "completed"
	SqlJobInfoFull_Status_Failed    = "failed"
	SqlJobInfoFull_Status_Queued    = "queued"
	SqlJobInfoFull_Status_Running   = "running"
)

// UnmarshalSqlJobInfoFull unmarshals an instance of SqlJobInfoFull from the specified map of raw messages.
func UnmarshalSqlJobInfoFull(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SqlJobInfoFull)
	err = core.UnmarshalPrimitive(m, "job_id", &obj.JobID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "submit_time", &obj.SubmitTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statement", &obj.Statement)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resultset_format", &obj.ResultsetFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resultset_location", &obj.ResultsetLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rows_returned", &obj.RowsReturned)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rows_read", &obj.RowsRead)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bytes_read", &obj.BytesRead)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "objects_skipped", &obj.ObjectsSkipped)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "objects_qualified", &obj.ObjectsQualified)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_message", &obj.ErrorMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hints", &obj.Hints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SqlJobInfoList : List of information about SQL jobs.
type SqlJobInfoList struct {
	// The SQL jobs.
	Jobs []SqlJobInfoShort `json:"jobs" validate:"required"`
}

// UnmarshalSqlJobInfoList unmarshals an instance of SqlJobInfoList from the specified map of raw messages.
func UnmarshalSqlJobInfoList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SqlJobInfoList)
	err = core.UnmarshalModel(m, "jobs", &obj.Jobs, UnmarshalSqlJobInfoShort)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SqlJobInfoShort : Abridged information about an SQL job, including its identifier and processing status.
type SqlJobInfoShort struct {
	// Identifier for an SQL job.
	JobID *string `json:"job_id" validate:"required"`

	// Execution status of an SQL job.
	Status *string `json:"status" validate:"required"`

	// ID of the user who submitted an SQL job.
	UserID *string `json:"user_id,omitempty"`

	// Timestamp indicating when an SQL job was accepted by the service.
	SubmitTime *strfmt.DateTime `json:"submit_time,omitempty"`

	// Boolean indicating when an SQL job has an improvement hint.
	HasHints *bool `json:"has_hints,omitempty"`
}

// Constants associated with the SqlJobInfoShort.Status property.
// Execution status of an SQL job.
const (
	SqlJobInfoShort_Status_Completed = "completed"
	SqlJobInfoShort_Status_Failed    = "failed"
	SqlJobInfoShort_Status_Queued    = "queued"
	SqlJobInfoShort_Status_Running   = "running"
)

// UnmarshalSqlJobInfoShort unmarshals an instance of SqlJobInfoShort from the specified map of raw messages.
func UnmarshalSqlJobInfoShort(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SqlJobInfoShort)
	err = core.UnmarshalPrimitive(m, "job_id", &obj.JobID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "submit_time", &obj.SubmitTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "has_hints", &obj.HasHints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableInformation : Detailed information about a catalog table.
type TableInformation struct {
	// The name of a catalog table.
	Name *string `json:"name" validate:"required"`

	// The type of a catalog table (for example, "TABLE" or "VIEW").
	Type *string `json:"type" validate:"required"`

	// The columns of the table.
	Columns []ColumnInformation `json:"columns" validate:"required"`
}

// UnmarshalTableInformation unmarshals an instance of TableInformation from the specified map of raw messages.
func UnmarshalTableInformation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableInformation)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "columns", &obj.Columns, UnmarshalColumnInformation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableList : List of catalog tables.
type TableList struct {
	// The table names.
	Tables []string `json:"tables" validate:"required"`

	// Metadata about the returned tables.
	TablesMetadata []TableMetadata `json:"tables_metadata" validate:"required"`
}

// UnmarshalTableList unmarshals an instance of TableList from the specified map of raw messages.
func UnmarshalTableList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableList)
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tables_metadata", &obj.TablesMetadata, UnmarshalTableMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableMetadata : Short metadata about a catalog table.
type TableMetadata struct {
	// The name of a catalog table.
	Name *string `json:"name" validate:"required"`

	// The type of a catalog table (for example, "TABLE" or "VIEW").
	Type *string `json:"type" validate:"required"`
}

// UnmarshalTableMetadata unmarshals an instance of TableMetadata from the specified map of raw messages.
func UnmarshalTableMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableMetadata)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
