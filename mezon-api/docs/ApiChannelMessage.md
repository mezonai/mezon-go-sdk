# ApiChannelMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClanId** | **string** | The clan this message belong to. | [optional] [default to null]
**ChannelId** | **string** | The channel this message belongs to. | [optional] [default to null]
**MessageId** | **string** | The unique ID of this message. | [optional] [default to null]
**Code** | **int32** | The code representing a message type or category. | [optional] [default to null]
**SenderId** | **string** | Message sender, usually a user ID. | [optional] [default to null]
**Username** | **string** | The username of the message sender, if any. | [optional] [default to null]
**Avatar** | **string** |  | [optional] [default to null]
**Content** | **string** | The content payload. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was created. | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was last updated. | [optional] [default to null]
**ChannelLabel** | **string** | The name of the chat room, or an empty string if this message was not sent through a chat room. | [optional] [default to null]
**ClanLogo** | **string** |  | [optional] [default to null]
**CategoryName** | **string** |  | [optional] [default to null]
**DisplayName** | **string** |  | [optional] [default to null]
**ClanNick** | **string** |  | [optional] [default to null]
**ClanAvatar** | **string** |  | [optional] [default to null]
**Reactions** | **string** |  | [optional] [default to null]
**Mentions** | **string** |  | [optional] [default to null]
**Attachments** | **string** |  | [optional] [default to null]
**References** | **string** |  | [optional] [default to null]
**ReferencedMessage** | **string** |  | [optional] [default to null]
**CreateTimeSeconds** | **int64** |  | [optional] [default to null]
**UpdateTimeSeconds** | **int64** |  | [optional] [default to null]
**Mode** | **int32** |  | [optional] [default to null]
**HideEditted** | **bool** |  | [optional] [default to null]
**IsPublic** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


