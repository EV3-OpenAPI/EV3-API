<a name="__pageTop"></a>
# ev3api.apis.tags.led_api.LedApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**button_pressed_get**](#button_pressed_get) | **get** /button/pressed | 
[**led_flash_post**](#led_flash_post) | **post** /led/flash | 

# **button_pressed_get**
<a name="button_pressed_get"></a>
> [str] button_pressed_get()



Get list of sensors currently pressed

### Example

```python
import ev3api
from ev3api.apis.tags import led_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = led_api.LedApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.button_pressed_get()
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling LedApi->button_pressed_get: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#button_pressed_get.ApiResponseFor200) | Success

#### button_pressed_get.ApiResponseFor200
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

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **led_flash_post**
<a name="led_flash_post"></a>
> led_flash_post(led)



### Example

```python
import ev3api
from ev3api.apis.tags import led_api
from ev3api.model.led import LED
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = led_api.LedApi(api_client)

    # example passing only required values which don't have defaults set
    body = [
        LED(
            device="device_example",
            red=1,
            green=1,
        )
    ]
    try:
        api_response = api_instance.led_flash_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling LedApi->led_flash_post: %s\n" % e)
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
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**LED**]({{complexTypePrefix}}LED.md) | [**LED**]({{complexTypePrefix}}LED.md) | [**LED**]({{complexTypePrefix}}LED.md) |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#led_flash_post.ApiResponseFor200) | 

#### led_flash_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

