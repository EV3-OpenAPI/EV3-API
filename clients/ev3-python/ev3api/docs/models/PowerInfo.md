# ev3api.model.power_info.PowerInfo

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**voltage** | decimal.Decimal, int, float,  | decimal.Decimal,  | Voltage returns voltage measured from the power supply in volts. | [optional] 
**current** | decimal.Decimal, int, float,  | decimal.Decimal,  | Current returns the current drawn from the power supply in milliamps. | [optional] 
**voltage_max** | decimal.Decimal, int, float,  | decimal.Decimal,  | VoltageMax returns the maximum design voltage for the power supply in volts. | [optional] 
**voltage_min** | decimal.Decimal, int, float,  | decimal.Decimal,  | VoltageMin returns the minimum design voltage for the power supply in volts. | [optional] 
**technology** | str,  | str,  | Technology returns the battery technology of the power supply. | [optional] 
**type** | str,  | str,  | Type returns the battery type of the power supply. | [optional] 
**[u_event](#u_event)** | dict, frozendict.frozendict,  | frozendict.frozendict,  | Uevent returns the current uevent state for the power supply. | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# u_event

Uevent returns the current uevent state for the power supply.

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  | Uevent returns the current uevent state for the power supply. | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**any_string_name** | str,  | str,  | any string name can be used but the value must be the correct type | [optional] 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

