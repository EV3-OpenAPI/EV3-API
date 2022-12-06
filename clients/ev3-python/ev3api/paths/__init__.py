# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from ev3api.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    BUTTON_PRESSED = "/button/pressed"
    LED_FLASH = "/led/flash"
    MOTOR_TACHO = "/motor/tacho"
    MOTOR_TACHO_TYPE_PORT = "/motor/tacho/{type}/{port}"
    MOTOR_STOP_ALL = "/motor/stopAll"
    MOTOR_STEER_RESET = "/motor/steer/reset"
    MOTOR_STEER_COUNTS = "/motor/steer/counts"
    MOTOR_STEER_DURATION = "/motor/steer/duration"
    POWER = "/power"
    SENSOR = "/sensor"
    SENSOR_TYPE = "/sensor/{type}"
    SENSOR_TYPE_VALUES = "/sensor/{type}/values"
    SENSOR_TYPE_TEXT_VALUES = "/sensor/{type}/text_values"
    SOUND_BEEP = "/sound/beep"
    SOUND_SPEAK = "/sound/speak"
    SOUND_TONE = "/sound/tone"
    SOUND_TONES = "/sound/tones"
