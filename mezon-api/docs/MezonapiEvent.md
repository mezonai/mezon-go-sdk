# MezonapiEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An event name, type, category, or identifier. | [optional] [default to null]
**Properties** | **map[string]string** | Arbitrary event property values. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The time when the event was triggered. | [optional] [default to null]
**External** | **bool** | True if the event came directly from a client call, false otherwise. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


