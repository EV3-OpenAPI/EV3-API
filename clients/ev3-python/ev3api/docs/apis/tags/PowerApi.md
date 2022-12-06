<a name="__pageTop"></a>
# ev3api.apis.tags.power_api.PowerApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**power_get**](#power_get) | **get** /power | 

# **power_get**
<a name="power_get"></a>
> PowerInfo power_get()



Get power info

### Example

```python
import ev3api
from ev3api.apis.tags import power_api
from ev3api.model.power_info import PowerInfo
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)

# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = power_api.PowerApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.power_get()
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling PowerApi->power_get: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#power_get.ApiResponseFor200) | Success

#### power_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**PowerInfo**](../../models/PowerInfo.md) |  | 


### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

