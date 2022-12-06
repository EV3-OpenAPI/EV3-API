<a name="__pageTop"></a>
# ev3api.apis.tags.motor_api.MotorApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**motor_steer_counts_post**](#motor_steer_counts_post) | **post** /motor/steer/counts | 
[**motor_steer_duration_post**](#motor_steer_duration_post) | **post** /motor/steer/duration | 
[**motor_steer_reset_post**](#motor_steer_reset_post) | **post** /motor/steer/reset | 
[**motor_stop_all_post**](#motor_stop_all_post) | **post** /motor/stopAll | 
[**motor_tacho_get**](#motor_tacho_get) | **get** /motor/tacho | 
[**motor_tacho_post**](#motor_tacho_post) | **post** /motor/tacho | 
[**motor_tacho_type_port_get**](#motor_tacho_type_port_get) | **get** /motor/tacho/{type}/{port} | 

# **motor_steer_counts_post**
<a name="motor_steer_counts_post"></a>
> motor_steer_counts_post(any_type)



Steers in the given turn for the given tacho counts and at the specified speed.

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from ev3api.model.steering_unit import SteeringUnit
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example passing only required values which don't have defaults set
    body = dict(
        steering_unit=SteeringUnit(
            left=Motor(
                size="s",
                port="port_example",
            ),
,
        ),
        speed=1,
        turn=1,
        counts=1,
    )
    try:
        api_response = api_instance.motor_steer_counts_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_steer_counts_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**steering_unit** | [**SteeringUnit**]({{complexTypePrefix}}SteeringUnit.md) | [**SteeringUnit**]({{complexTypePrefix}}SteeringUnit.md) |  | [optional] 
**speed** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**turn** | decimal.Decimal, int,  | decimal.Decimal,  | The valid range of turn is -100 (hard left) to +100 (hard right) | [optional] 
**counts** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_steer_counts_post.ApiResponseFor200) | Success
400 | [ApiResponseFor400](#motor_steer_counts_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#motor_steer_counts_post.ApiResponseFor500) | Server error

#### motor_steer_counts_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_counts_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_counts_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_steer_duration_post**
<a name="motor_steer_duration_post"></a>
> motor_steer_duration_post(any_type)



Steers in the given turn for the given duration and at the specified speed

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from ev3api.model.steering_unit import SteeringUnit
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example passing only required values which don't have defaults set
    body = dict(
        steering_unit=SteeringUnit(
            left=Motor(
                size="s",
                port="port_example",
            ),
,
        ),
        speed=1,
        turn=1,
        duration_ms=1,
    )
    try:
        api_response = api_instance.motor_steer_duration_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_steer_duration_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**steering_unit** | [**SteeringUnit**]({{complexTypePrefix}}SteeringUnit.md) | [**SteeringUnit**]({{complexTypePrefix}}SteeringUnit.md) |  | [optional] 
**speed** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**turn** | decimal.Decimal, int,  | decimal.Decimal,  | The valid range of turn is -100 (hard left) to +100 (hard right) | [optional] 
**duration_ms** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_steer_duration_post.ApiResponseFor200) | Success
400 | [ApiResponseFor400](#motor_steer_duration_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#motor_steer_duration_post.ApiResponseFor500) | Server error

#### motor_steer_duration_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_duration_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_duration_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_steer_reset_post**
<a name="motor_steer_reset_post"></a>
> motor_steer_reset_post()



Reset steering unit

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.motor_steer_reset_post()
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_steer_reset_post: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_steer_reset_post.ApiResponseFor200) | Success
400 | [ApiResponseFor400](#motor_steer_reset_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#motor_steer_reset_post.ApiResponseFor500) | Server error

#### motor_steer_reset_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_reset_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_steer_reset_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_stop_all_post**
<a name="motor_stop_all_post"></a>
> motor_stop_all_post()



Stops all connected motors

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.motor_stop_all_post()
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_stop_all_post: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_stop_all_post.ApiResponseFor200) | Success
400 | [ApiResponseFor400](#motor_stop_all_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#motor_stop_all_post.ApiResponseFor500) | Server error

#### motor_stop_all_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_stop_all_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_stop_all_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_tacho_get**
<a name="motor_tacho_get"></a>
> [TachoMotor] motor_tacho_get()



Get information about all connected motors

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from ev3api.model.tacho_motor import TachoMotor
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.motor_tacho_get()
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_tacho_get: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_tacho_get.ApiResponseFor200) | Success

#### motor_tacho_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**TachoMotor**]({{complexTypePrefix}}TachoMotor.md) | [**TachoMotor**]({{complexTypePrefix}}TachoMotor.md) | [**TachoMotor**]({{complexTypePrefix}}TachoMotor.md) |  | 

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_tacho_post**
<a name="motor_tacho_post"></a>
> motor_tacho_post(motor_request)



Set motor parameters

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from ev3api.model.motor_request import MotorRequest
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example passing only required values which don't have defaults set
    body = MotorRequest(
        motors=[
            Motor(
                size="s",
                port="port_example",
            )
        ],
        speed=1,
        speed_pct=1,
        polarity="polarity_example",
        position=1,
        time=1,
        position_setpoint=1,
        ramp_up_setpoint=1,
        ramp_down_setpoint=1,
        hold_pid_kd=1,
        hold_pid_ki=1,
        hold_pid_kp=1,
        speed_pid_kd=1,
        speed_pid_ki=1,
        speed_pid_kp=1,
        command="run-forever",
        stop_action="coast",
    )
    try:
        api_response = api_instance.motor_tacho_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_tacho_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**MotorRequest**](../../models/MotorRequest.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_tacho_post.ApiResponseFor200) | Success
404 | [ApiResponseFor404](#motor_tacho_post.ApiResponseFor404) | Motor not found

#### motor_tacho_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### motor_tacho_post.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **motor_tacho_type_port_get**
<a name="motor_tacho_type_port_get"></a>
> TachoMotor motor_tacho_type_port_get(typeport)



Get information about this motor

### Example

```python
import ev3api
from ev3api.apis.tags import motor_api
from ev3api.model.tacho_motor import TachoMotor
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = motor_api.MotorApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'type': "type_example",
        'port': "port_example",
    }
    try:
        api_response = api_instance.motor_tacho_type_port_get(
            path_params=path_params,
        )
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling MotorApi->motor_tacho_type_port_get: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
type | TypeSchema | | 
port | PortSchema | | 

# TypeSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# PortSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#motor_tacho_type_port_get.ApiResponseFor200) | Success
404 | [ApiResponseFor404](#motor_tacho_type_port_get.ApiResponseFor404) | Motor not found

#### motor_tacho_type_port_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**TachoMotor**](../../models/TachoMotor.md) |  | 


#### motor_tacho_type_port_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

