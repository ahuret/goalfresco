/*
Alfresco Content Services REST API

Testing VersionsApiService

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

func Test_openapi_VersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VersionsApiService CreateVersionRendition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		httpRes, err := apiClient.VersionsApi.CreateVersionRendition(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService DeleteVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		httpRes, err := apiClient.VersionsApi.DeleteVersion(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GetVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		resp, httpRes, err := apiClient.VersionsApi.GetVersion(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GetVersionContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		resp, httpRes, err := apiClient.VersionsApi.GetVersionContent(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GetVersionRendition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string
		var renditionId string

		resp, httpRes, err := apiClient.VersionsApi.GetVersionRendition(context.Background(), nodeId, versionId, renditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService GetVersionRenditionContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string
		var renditionId string

		resp, httpRes, err := apiClient.VersionsApi.GetVersionRenditionContent(context.Background(), nodeId, versionId, renditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService ListVersionHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.VersionsApi.ListVersionHistory(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService ListVersionRenditions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		resp, httpRes, err := apiClient.VersionsApi.ListVersionRenditions(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsApiService RevertVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string
		var versionId string

		resp, httpRes, err := apiClient.VersionsApi.RevertVersion(context.Background(), nodeId, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
