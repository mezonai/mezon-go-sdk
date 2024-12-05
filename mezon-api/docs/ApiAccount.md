# ApiAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [***ApiUser**](apiUser.md) | The user object. | [optional] [default to null]
**Wallet** | **string** | The user&#39;s wallet data. | [optional] [default to null]
**Email** | **string** | The email address of the user. | [optional] [default to null]
**Devices** | [**[]ApiAccountDevice**](apiAccountDevice.md) | The devices which belong to the user&#39;s account. | [optional] [default to null]
**CustomId** | **string** | The custom id in the user&#39;s account. | [optional] [default to null]
**VerifyTime** | [**time.Time**](time.Time.md) | The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user&#39;s email was verified. | [optional] [default to null]
**DisableTime** | [**time.Time**](time.Time.md) | The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user&#39;s account was disabled/banned. | [optional] [default to null]
**Logo** | **string** |  | [optional] [default to null]
**SplashScreen** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


