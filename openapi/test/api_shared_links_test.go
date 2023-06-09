/*
Alfresco Content Services REST API

Testing SharedLinksApiService

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

func Test_openapi_SharedLinksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SharedLinksApiService CreateSharedLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SharedLinksApi.CreateSharedLink(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService DeleteSharedLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string

		httpRes, err := apiClient.SharedLinksApi.DeleteSharedLink(context.Background(), sharedId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService EmailSharedLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string

		httpRes, err := apiClient.SharedLinksApi.EmailSharedLink(context.Background(), sharedId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService GetSharedLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string

		resp, httpRes, err := apiClient.SharedLinksApi.GetSharedLink(context.Background(), sharedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService GetSharedLinkContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string

		resp, httpRes, err := apiClient.SharedLinksApi.GetSharedLinkContent(context.Background(), sharedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService GetSharedLinkRendition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string
		var renditionId string

		resp, httpRes, err := apiClient.SharedLinksApi.GetSharedLinkRendition(context.Background(), sharedId, renditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService GetSharedLinkRenditionContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string
		var renditionId string

		resp, httpRes, err := apiClient.SharedLinksApi.GetSharedLinkRenditionContent(context.Background(), sharedId, renditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService ListSharedLinkRenditions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedId string

		resp, httpRes, err := apiClient.SharedLinksApi.ListSharedLinkRenditions(context.Background(), sharedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedLinksApiService ListSharedLinks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SharedLinksApi.ListSharedLinks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
