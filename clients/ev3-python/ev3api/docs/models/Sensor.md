# ev3api.model.sensor.Sensor

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**type** | str,  | str,  | Sensor information are read only and from type string | [optional] 
**driver_name** | str,  | str,  | Sensor driver_name information is read only and from type string | [optional] 
**port** | str,  | str,  | Sensor portnumber is read only and from type string | [optional] 
**[modes](#modes)** | list, tuple,  | tuple,  | Sensor modes information are read only and in a array | [optional] 
**[commands](#commands)** | list, tuple,  | tuple,  | Sensor commands are in a array | [optional] 
**mode** | str,  | str,  | Sensor mode | [optional] 
**decimals** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**poll_rate_ms** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**units** | str,  | str,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# modes

Sensor modes information are read only and in a array

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  | Sensor modes information are read only and in a array | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str,  | str,  |  | 

# commands

Sensor commands are in a array

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  | Sensor commands are in a array | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str,  | str,  |  | 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

