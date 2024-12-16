# ApiNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the Notification. | [optional] [default to null]
**Subject** | **string** | Subject of the notification. | [optional] [default to null]
**Content** | **string** | Content of the notification in JSON. | [optional] [default to null]
**Code** | **int32** | Category code for this notification. | [optional] [default to null]
**SenderId** | **string** | ID of the sender, if a user. Otherwise &#39;null&#39;. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The UNIX time (for gRPC clients) or ISO string (for REST clients) when the notification was created. | [optional] [default to null]
**Persistent** | **bool** | True if this notification was persisted to the database. | [optional] [default to null]
**ClanId** | **string** |  | [optional] [default to null]
**ChannelId** | **string** |  | [optional] [default to null]
**ChannelType** | **int32** |  | [optional] [default to null]
**AvatarUrl** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


