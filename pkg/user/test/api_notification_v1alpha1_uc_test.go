/*
Halo

Testing NotificationV1alpha1UcAPIService

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

func Test_user_sdk_NotificationV1alpha1UcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationV1alpha1UcAPIService DeleteSpecifiedNotification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var name string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.DeleteSpecifiedNotification(context.Background(), username, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationV1alpha1UcAPIService ListUserNotificationPreferences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.ListUserNotificationPreferences(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationV1alpha1UcAPIService ListUserNotifications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.ListUserNotifications(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationV1alpha1UcAPIService MarkNotificationAsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var name string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.MarkNotificationAsRead(context.Background(), username, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationV1alpha1UcAPIService MarkNotificationsAsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.MarkNotificationsAsRead(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationV1alpha1UcAPIService SaveUserNotificationPreferences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.NotificationV1alpha1UcAPI.SaveUserNotificationPreferences(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
