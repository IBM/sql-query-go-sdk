package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/sql-query-go-sdk/sqlv2"
)

func main() {
	// Set environment varable IAM_APIKEY to your IBM Cloud IAM ApiKey 
	iamApikey := os.Getenv("IAM_APIKEY")
	// Set environment varable SQL_QUERY_CRN to your SqlQuery instanceCrn
	instance := os.Getenv("SQL_QUERY_CRN")
	// Set environment varable SQL_QUERY_TARGET to the target COS location for query result,
	// example:  cos://us-geo/<bucket>/<prefix>
	target := os.Getenv("SQL_QUERY_TARGET")

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: iamApikey,
		// Comment the following line to use default. Use an alternative iam token server
		// URL: "https://iam.test.cloud.ibm.com/identity/token",
	}

	sqlService, err := sqlv2.NewSqlV2(&sqlv2.SqlV2Options{
		InstanceCrn:   &instance,
		Authenticator: authenticator,
	})
	if err != nil {
		fmt.Println("Failed to instantiate service:", err)
		return
	}

	// Comment the following lines to use default. Use an alternative SqlQuery server
	// err = sqlService.SetServiceURL("https://api.sql-query.test.cloud.ibm.com/v2")
	// if err != nil {
	// 	fmt.Println("Failed to set new service URL:", err)
	// 	return
	// }

	// Submit a sql job
	submitResult, _, submitRrr := sqlService.SubmitSqlJob(&sqlv2.SubmitSqlJobOptions{
		Statement: core.StringPtr("SELECT * FROM cos://us-geo/sql/customers.csv"),
		ResultsetTarget: core.StringPtr(target),
	})
	if submitRrr != nil {
		fmt.Println("Failed to submit a job:", err)
		return
	}
	submitDetails, err := json.Marshal(submitResult)
	if err != nil {
		fmt.Println("Failed to marshal job into json:", err)
		return
	}
	fmt.Println("A job has been successfully submitted:\n", string(submitDetails))
	jobID := submitResult.JobID

	// Get the info of a single sql job
	getJobResult, _, getJobErr := sqlService.GetSqlJob(&sqlv2.GetSqlJobOptions{
		JobID: jobID,
	})
	if getJobErr != nil {
		fmt.Println("Failed to get job info:", getJobErr)
		return
	}
	getJobDetails, err := json.Marshal(getJobResult)
	if err != nil {
		fmt.Println("Failed to marshal job into json:", err)
		return
	}
	fmt.Println("\nThe job info has been successfully retrieved:\n", string(getJobDetails))

	// Get the list of jobs that exist on the instance
	listJobsResult, _, listJobsErr := sqlService.ListSqlJobs(&sqlv2.ListSqlJobsOptions{})
	if listJobsErr != nil {
		fmt.Println("Failed to list jobs:", listJobsErr)
		return
	}
	fmt.Println("\nThe job list has been successfully retrieved:")
	for i, job := range listJobsResult.Jobs {
		details, err := json.Marshal(job)
		if err != nil {
			fmt.Println("Failed to marshal job into json:", err)
			return
		}
		fmt.Println(i, string(details))
	}

	// Get the list of tables that exist on the instance
	listTablesResult, _, listTablesErr := sqlService.ListTables(&sqlv2.ListTablesOptions{
		NamePattern: core.StringPtr("*"),
		Type: core.StringPtr("table"),
	})
	if listTablesErr != nil {
		fmt.Println("Failed to list tables:", listTablesErr)
		return
	}
	fmt.Println("\nThe table list has been successfully retrieved:")
	tableName := ""
	for i, table := range listTablesResult.TablesMetadata {
		tableName = *table.Name
		details, err := json.Marshal(table)
		if err != nil {
			fmt.Println("Failed to marshal table into json:", err)
			return
		}
		fmt.Println(i, string(details))
	}

	// Get the detailed info of a single table
	if tableName != "" {
		getTableResult, _, getTableErr := sqlService.GetTable(&sqlv2.GetTableOptions{
			TableName: core.StringPtr(tableName),
		})
		if getTableErr != nil {
			fmt.Println("Failed to get table info:", getTableErr)
			return
		}
		getTableDetails, err := json.Marshal(getTableResult)
		if err != nil {
			fmt.Println("Failed to marshal table into json:", err)
			return
		}
		fmt.Println("\nThe table info has been successfully retrieved:\n", string(getTableDetails))
	}
}
