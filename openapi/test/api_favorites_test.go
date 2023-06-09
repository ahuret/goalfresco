/*
Alfresco Content Services REST API

Testing FavoritesApiService

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

func Test_openapi_FavoritesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FavoritesApiService CreateFavorite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.FavoritesApi.CreateFavorite(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService CreateSiteFavorite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.FavoritesApi.CreateSiteFavorite(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService DeleteFavorite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var favoriteId string

		httpRes, err := apiClient.FavoritesApi.DeleteFavorite(context.Background(), personId, favoriteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService DeleteSiteFavorite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var siteId string

		httpRes, err := apiClient.FavoritesApi.DeleteSiteFavorite(context.Background(), personId, siteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService GetFavorite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var favoriteId string

		resp, httpRes, err := apiClient.FavoritesApi.GetFavorite(context.Background(), personId, favoriteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService GetFavoriteSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var siteId string

		resp, httpRes, err := apiClient.FavoritesApi.GetFavoriteSite(context.Background(), personId, siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService ListFavoriteSitesForPerson", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.FavoritesApi.ListFavoriteSitesForPerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoritesApiService ListFavorites", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.FavoritesApi.ListFavorites(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
