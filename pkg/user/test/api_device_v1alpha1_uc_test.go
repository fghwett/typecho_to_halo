/*
Halo

Testing DeviceV1alpha1UcAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package user_sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func Test_user_sdk_DeviceV1alpha1UcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceV1alpha1UcAPIService ListDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeviceV1alpha1UcAPI.ListDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceV1alpha1UcAPIService RevokeDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceV1alpha1UcAPI.RevokeDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
