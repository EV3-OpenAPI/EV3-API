import time
import ev3api

from pprint import pprint
from ev3api.apis.tags import led_api
from ev3api.model.led import LED
# Defining the host is optional and defaults to http://127.0.0.1:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ev3api.Configuration(
    host = "http://127.0.0.1:8080/api/v1"
)


# Enter a context with an instance of the API client
with ev3api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = led_api.LedApi(api_client)

    try:
        api_response = api_instance.button_pressed_get()
        pprint(api_response)
    except ev3api.ApiException as e:
        print("Exception when calling LedApi->button_pressed_get: %s\n" % e)