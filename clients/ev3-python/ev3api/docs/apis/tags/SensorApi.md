<a name="__pageTop"></a>
# ev3api.apis.tags.sensor_api.SensorApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sensor_get**](#sensor_get) | **get** /sensor | 
[**sensor_type_get**](#sensor_type_get) | **get** /sensor/{type} | 
[**sensor_type_put**](#sensor_type_put) | **put** /sensor/{type} | 
[**sensor_type_text_values_get**](#sensor_type_text_values_get) | **get** /sensor/{type}/text_values | 
[**sensor_type_values_get**](#sensor_type_values_get) | **get** /sensor/{type}/values | 

# **sensor_get**
<a name="sensor_get"></a>
> [Sensor] sensor_get()



Get information about the sensor

### Example

```python
import ev3api
from ev3api.apis.tags import sensor_api
from ev3api.model.sensor import Sensor
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sensor_api.SensorApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.sensor_get()
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling SensorApi->sensor_get: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sensor_get.ApiResponseFor200) | List of all connected sensors

#### sensor_get.ApiResponseFor200
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
[**Sensor**]({{complexTypePrefix}}Sensor.md) | [**Sensor**]({{complexTypePrefix}}Sensor.md) | [**Sensor**]({{complexTypePrefix}}Sensor.md) |  | 

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sensor_type_get**
<a name="sensor_type_get"></a>
> Sensor sensor_type_get(type)



Get details about a specific sensor

### Example

```python
import ev3api
from ev3api.apis.tags import sensor_api
from ev3api.model.sensor import Sensor
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sensor_api.SensorApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'type': "type_example",
    }
    try:
        api_response = api_instance.sensor_type_get(
            path_params=path_params,
        )
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling SensorApi->sensor_type_get: %s\n" % e)
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

# TypeSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sensor_type_get.ApiResponseFor200) | Details of sensors

#### sensor_type_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Sensor**](../../models/Sensor.md) |  | 


### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sensor_type_put**
<a name="sensor_type_put"></a>
> sensor_type_put(typesensor)



update a sensor with specific values

### Example

```python
import ev3api
from ev3api.apis.tags import sensor_api
from ev3api.model.sensor import Sensor
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sensor_api.SensorApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'type': "type_example",
    }
    body = Sensor(
        type="type_example",
        driver_name="driver_name_example",
        port="port_example",
        modes=[
            "modes_example"
        ],
        commands=[
            "commands_example"
        ],
        mode="mode_example",
        decimals=1,
        poll_rate_ms=1,
        units="units_example",
    )
    try:
        api_response = api_instance.sensor_type_put(
            path_params=path_params,
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling SensorApi->sensor_type_put: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Sensor**](../../models/Sensor.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
type | TypeSchema | | 

# TypeSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sensor_type_put.ApiResponseFor200) | Sensor successfully updated

#### sensor_type_put.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sensor_type_text_values_get**
<a name="sensor_type_text_values_get"></a>
> [str] sensor_type_text_values_get(type)



Get the values about a specific sensor in textformat

### Example

```python
import ev3api
from ev3api.apis.tags import sensor_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sensor_api.SensorApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'type': "type_example",
    }
    try:
        api_response = api_instance.sensor_type_text_values_get(
            path_params=path_params,
        )
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling SensorApi->sensor_type_text_values_get: %s\n" % e)
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

# TypeSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sensor_type_text_values_get.ApiResponseFor200) | Values in textformat is successfully found
404 | [ApiResponseFor404](#sensor_type_text_values_get.ApiResponseFor404) | Sensor of that type not found

#### sensor_type_text_values_get.ApiResponseFor200
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
items | str,  | str,  |  | 

#### sensor_type_text_values_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sensor_type_values_get**
<a name="sensor_type_values_get"></a>
> [str] sensor_type_values_get(type)



Get the values about a specific sensor

### Example

```python
import ev3api
from ev3api.apis.tags import sensor_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sensor_api.SensorApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'type': "type_example",
    }
    try:
        api_response = api_instance.sensor_type_values_get(
            path_params=path_params,
        )
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling SensorApi->sensor_type_values_get: %s\n" % e)
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

# TypeSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sensor_type_values_get.ApiResponseFor200) | Sensor of that type is successfully found
404 | [ApiResponseFor404](#sensor_type_values_get.ApiResponseFor404) | Sensor of that type not found

#### sensor_type_values_get.ApiResponseFor200
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
items | str,  | str,  |  | 

#### sensor_type_values_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

