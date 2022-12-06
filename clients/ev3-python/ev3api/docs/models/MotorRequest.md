# ev3api.model.motor_request.MotorRequest

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[motors](#motors)** | list, tuple,  | tuple,  |  | [optional] 
**speed** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**speed_pct** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**polarity** | str,  | str,  |  | [optional] 
**position** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**time** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**position_setpoint** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**ramp_up_setpoint** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**ramp_down_setpoint** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**hold_PID_kd** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**hold_PID_ki** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**hold_PID_kp** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**speed_PID_kd** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**speed_PID_ki** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**speed_PID_kp** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**command** | str,  | str,  |  | [optional] must be one of ["run-forever", "run-to-abs-pos", "run-to-rel-pos", "run-timed", "run-direct", "stop", "reset", ] 
**stop-action** | str,  | str,  |  | [optional] must be one of ["coast", "brake", "hold", ] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# motors

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Motor**](Motor.md) | [**Motor**](Motor.md) | [**Motor**](Motor.md) |  | 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

