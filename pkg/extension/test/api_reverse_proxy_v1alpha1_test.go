/*
Halo

Testing ReverseProxyV1alpha1APIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package extension_sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func Test_extension_sdk_ReverseProxyV1alpha1APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReverseProxyV1alpha1APIService CreateReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReverseProxyV1alpha1API.CreateReverseProxy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReverseProxyV1alpha1APIService DeleteReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.ReverseProxyV1alpha1API.DeleteReverseProxy(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReverseProxyV1alpha1APIService GetReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ReverseProxyV1alpha1API.GetReverseProxy(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReverseProxyV1alpha1APIService ListReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReverseProxyV1alpha1API.ListReverseProxy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReverseProxyV1alpha1APIService PatchReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ReverseProxyV1alpha1API.PatchReverseProxy(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReverseProxyV1alpha1APIService UpdateReverseProxy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ReverseProxyV1alpha1API.UpdateReverseProxy(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
