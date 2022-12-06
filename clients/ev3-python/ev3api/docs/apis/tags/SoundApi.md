<a name="__pageTop"></a>
# ev3api.apis.tags.sound_api.SoundApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sound_beep_post**](#sound_beep_post) | **post** /sound/beep | 
[**sound_speak_post**](#sound_speak_post) | **post** /sound/speak | 
[**sound_tone_post**](#sound_tone_post) | **post** /sound/tone | 
[**sound_tones_post**](#sound_tones_post) | **post** /sound/tones | 

# **sound_beep_post**
<a name="sound_beep_post"></a>
> sound_beep_post()



Output of the roboter is a beep sound

### Example

```python
import ev3api
from ev3api.apis.tags import sound_api
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sound_api.SoundApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.sound_beep_post()
    except ev3api.ApiException as e:
        print("Exception when calling SoundApi->sound_beep_post: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sound_beep_post.ApiResponseFor200) | Beep successfully played
400 | [ApiResponseFor400](#sound_beep_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#sound_beep_post.ApiResponseFor500) | Server error

#### sound_beep_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_beep_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_beep_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sound_speak_post**
<a name="sound_speak_post"></a>
> sound_speak_post(text)



Output from the roboter is a spoken text

### Example

```python
import ev3api
from ev3api.apis.tags import sound_api
from ev3api.model.text import Text
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sound_api.SoundApi(api_client)

    # example passing only required values which don't have defaults set
    body = Text(
        text="text_example",
    )
    try:
        api_response = api_instance.sound_speak_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling SoundApi->sound_speak_post: %s\n" % e)
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
[**Text**](../../models/Text.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sound_speak_post.ApiResponseFor200) | Text successfully spoken
400 | [ApiResponseFor400](#sound_speak_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#sound_speak_post.ApiResponseFor500) | Server error

#### sound_speak_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_speak_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_speak_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sound_tone_post**
<a name="sound_tone_post"></a>
> sound_tone_post(tone)



### Example

```python
import ev3api
from ev3api.apis.tags import sound_api
from ev3api.model.tone import Tone
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sound_api.SoundApi(api_client)

    # example passing only required values which don't have defaults set
    body = Tone(
        frequency=1,
        length_ms=1,
    )
    try:
        api_response = api_instance.sound_tone_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling SoundApi->sound_tone_post: %s\n" % e)
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
[**Tone**](../../models/Tone.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sound_tone_post.ApiResponseFor200) | Tone successfully played
400 | [ApiResponseFor400](#sound_tone_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#sound_tone_post.ApiResponseFor500) | Server error

#### sound_tone_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_tone_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_tone_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **sound_tones_post**
<a name="sound_tones_post"></a>
> sound_tones_post(tone)



Output from the roboter are played tones

### Example

```python
import ev3api
from ev3api.apis.tags import sound_api
from ev3api.model.tone import Tone
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sound_api.SoundApi(api_client)

    # example passing only required values which don't have defaults set
    body = [
        Tone(
            frequency=1,
            length_ms=1,
        )
    ]
    try:
        api_response = api_instance.sound_tones_post(
            body=body,
        )
    except ev3api.ApiException as e:
        print("Exception when calling SoundApi->sound_tones_post: %s\n" % e)
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
[**Tone**]({{complexTypePrefix}}Tone.md) | [**Tone**]({{complexTypePrefix}}Tone.md) | [**Tone**]({{complexTypePrefix}}Tone.md) |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#sound_tones_post.ApiResponseFor200) | Tone successfully played
400 | [ApiResponseFor400](#sound_tones_post.ApiResponseFor400) | Client error
500 | [ApiResponseFor500](#sound_tones_post.ApiResponseFor500) | Server error

#### sound_tones_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_tones_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### sound_tones_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

