/*
Alfresco Content Services REST API

Testing AuditApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ahuret/goalfresco/openapi"
)

func Test_openapi_AuditApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditApiService DeleteAuditEntriesForAuditApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string

		httpRes, err := apiClient.AuditApi.DeleteAuditEntriesForAuditApp(context.Background(), auditApplicationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService DeleteAuditEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string
		var auditEntryId string

		httpRes, err := apiClient.AuditApi.DeleteAuditEntry(context.Background(), auditApplicationId, auditEntryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService GetAuditApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string

		resp, httpRes, err := apiClient.AuditApi.GetAuditApp(context.Background(), auditApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService GetAuditEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string
		var auditEntryId string

		resp, httpRes, err := apiClient.AuditApi.GetAuditEntry(context.Background(), auditApplicationId, auditEntryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService ListAuditApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditApi.ListAuditApps(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService ListAuditEntriesForAuditApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string

		resp, httpRes, err := apiClient.AuditApi.ListAuditEntriesForAuditApp(context.Background(), auditApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService ListAuditEntriesForNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.AuditApi.ListAuditEntriesForNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditApiService UpdateAuditApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditApplicationId string

		resp, httpRes, err := apiClient.AuditApi.UpdateAuditApp(context.Background(), auditApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
