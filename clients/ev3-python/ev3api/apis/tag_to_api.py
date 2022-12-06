import typing_extensions

from ev3api.apis.tags import TagValues
from ev3api.apis.tags.button_api import ButtonApi
from ev3api.apis.tags.led_api import LedApi
from ev3api.apis.tags.motor_api import MotorApi
from ev3api.apis.tags.power_api import PowerApi
from ev3api.apis.tags.sensor_api import SensorApi
from ev3api.apis.tags.sound_api import SoundApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.BUTTON: ButtonApi,
        TagValues.LED: LedApi,
        TagValues.MOTOR: MotorApi,
        TagValues.POWER: PowerApi,
        TagValues.SENSOR: SensorApi,
        TagValues.SOUND: SoundApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.BUTTON: ButtonApi,
        TagValues.LED: LedApi,
        TagValues.MOTOR: MotorApi,
        TagValues.POWER: PowerApi,
        TagValues.SENSOR: SensorApi,
        TagValues.SOUND: SoundApi,
    }
)
