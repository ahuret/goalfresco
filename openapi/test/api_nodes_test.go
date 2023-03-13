/*
Alfresco Content Services REST API

Testing NodesApiService

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

func Test_openapi_NodesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NodesApiService CopyNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.CopyNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.CreateAssociation(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.CreateNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateSecondaryChildAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.CreateSecondaryChildAssociation(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var targetId string

		httpRes, err := apiClient.NodesApi.DeleteAssociation(context.Background(), nodeId, targetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		httpRes, err := apiClient.NodesApi.DeleteNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteSecondaryChildAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var childId string

		httpRes, err := apiClient.NodesApi.DeleteSecondaryChildAssociation(context.Background(), nodeId, childId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.GetNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNodeContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.GetNodeContent(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ListNodeChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.ListNodeChildren(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ListParents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.ListParents(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ListSecondaryChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.ListSecondaryChildren(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ListSourceAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.ListSourceAssociations(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ListTargetAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.ListTargetAssociations(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService LockNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.LockNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService MoveNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.MoveNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UnlockNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.UnlockNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.UpdateNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateNodeContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.NodesApi.UpdateNodeContent(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
