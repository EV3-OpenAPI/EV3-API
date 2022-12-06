import typing_extensions

from ev3api.paths import PathValues
from ev3api.apis.paths.button_pressed import ButtonPressed
from ev3api.apis.paths.led_flash import LedFlash
from ev3api.apis.paths.motor_tacho import MotorTacho
from ev3api.apis.paths.motor_tacho_type_port import MotorTachoTypePort
from ev3api.apis.paths.motor_stop_all import MotorStopAll
from ev3api.apis.paths.motor_steer_reset import MotorSteerReset
from ev3api.apis.paths.motor_steer_counts import MotorSteerCounts
from ev3api.apis.paths.motor_steer_duration import MotorSteerDuration
from ev3api.apis.paths.power import Power
from ev3api.apis.paths.sensor import Sensor
from ev3api.apis.paths.sensor_type import SensorType
from ev3api.apis.paths.sensor_type_values import SensorTypeValues
from ev3api.apis.paths.sensor_type_text_values import SensorTypeTextValues
from ev3api.apis.paths.sound_beep import SoundBeep
from ev3api.apis.paths.sound_speak import SoundSpeak
from ev3api.apis.paths.sound_tone import SoundTone
from ev3api.apis.paths.sound_tones import SoundTones

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.BUTTON_PRESSED: ButtonPressed,
        PathValues.LED_FLASH: LedFlash,
        PathValues.MOTOR_TACHO: MotorTacho,
        PathValues.MOTOR_TACHO_TYPE_PORT: MotorTachoTypePort,
        PathValues.MOTOR_STOP_ALL: MotorStopAll,
        PathValues.MOTOR_STEER_RESET: MotorSteerReset,
        PathValues.MOTOR_STEER_COUNTS: MotorSteerCounts,
        PathValues.MOTOR_STEER_DURATION: MotorSteerDuration,
        PathValues.POWER: Power,
        PathValues.SENSOR: Sensor,
        PathValues.SENSOR_TYPE: SensorType,
        PathValues.SENSOR_TYPE_VALUES: SensorTypeValues,
        PathValues.SENSOR_TYPE_TEXT_VALUES: SensorTypeTextValues,
        PathValues.SOUND_BEEP: SoundBeep,
        PathValues.SOUND_SPEAK: SoundSpeak,
        PathValues.SOUND_TONE: SoundTone,
        PathValues.SOUND_TONES: SoundTones,
    }
)

path_to_api = PathToApi(
    {
        PathValues.BUTTON_PRESSED: ButtonPressed,
        PathValues.LED_FLASH: LedFlash,
        PathValues.MOTOR_TACHO: MotorTacho,
        PathValues.MOTOR_TACHO_TYPE_PORT: MotorTachoTypePort,
        PathValues.MOTOR_STOP_ALL: MotorStopAll,
        PathValues.MOTOR_STEER_RESET: MotorSteerReset,
        PathValues.MOTOR_STEER_COUNTS: MotorSteerCounts,
        PathValues.MOTOR_STEER_DURATION: MotorSteerDuration,
        PathValues.POWER: Power,
        PathValues.SENSOR: Sensor,
        PathValues.SENSOR_TYPE: SensorType,
        PathValues.SENSOR_TYPE_VALUES: SensorTypeValues,
        PathValues.SENSOR_TYPE_TEXT_VALUES: SensorTypeTextValues,
        PathValues.SOUND_BEEP: SoundBeep,
        PathValues.SOUND_SPEAK: SoundSpeak,
        PathValues.SOUND_TONE: SoundTone,
        PathValues.SOUND_TONES: SoundTones,
    }
)
