from enum import Enum
from ev3api.api.motor_api import MotorApi
from ev3api.api.power_api import PowerApi
from ev3api.api.sensor_api import SensorApi
from ev3api.api.sound_api import SoundApi
from ev3api.api.led_api import LedApi
from ev3api.model.led import LED
from ev3api.model.tone import Tone
from ev3api.model.text import Text
from ev3api.exceptions import ApiException
from ev3api.configuration import Configuration
from ev3api.api_client import ApiClient


class EV3:

    def __init__(self, host_address):
        self.configuration = Configuration(
            host=f"http://{host_address}/api/v1"
        )
        self.hostAddress = host_address
        self.api_client = ApiClient(self.configuration)
        self.motorApi = MotorApi(self.api_client)
        self.powerApi = PowerApi(self.api_client)
        self.sensorApi = SensorApi(self.api_client)
        self.soundApi = SoundApi(self.api_client)
        self.ledApi = LedApi(self.api_client)

    def get_host_address(self) -> str:
        """
        This method returns the ip-address from the EV3

        :return: the host address
        """

        return self.hostAddress

    def beep(self) -> None:
        """
        The EV3 will beep.
        """

        self.soundApi.sound_beep_post()

    def play_tone(self, frequency: int, length_ms: int) -> None:
        """
        The EV3 will play a tone.

        :param frequency: in herz
        :param length_ms: duration in ms
        """

        self.soundApi.sound_tone_post(Tone(
            frequency=frequency,
            length_ms=length_ms,
        ))

    def speak(self, text: str) -> None:
        """
        The EV3 will speak a specific text.

        :param text: the text to be spoken
        """

        self.soundApi.sound_speak_post(Text(
            text=text,
        ))

    def voltage(self) -> float:
        """
        Returns the current voltage of the battery.

        :return: voltage of the battery
        """
        return self.powerApi.power_get()["voltage"]

    def current(self) -> float:
        """
        Returns the current current of the battery

        :return: current of the battery
        """
        return self.powerApi.power_get()["current"]

    def max_voltage(self) -> float:
        """
        Returns the maximum voltage for the battery

        :return: maximum voltage of the battery
        """
        return self.powerApi.power_get()["voltage_max"]

    def min_voltage(self) -> float:
        """
        Returns the minimum voltage for the battery

        :return: minimum voltage of the battery
        """
        return self.powerApi.power_get()["voltage_min"]

    def technology(self) -> str:
        """
        Returns the technology of the battery
        :return: the technology of the battery
        """
        try:
            return self.powerApi.power_get()["technology"]
        except ApiException as e:
            print(e)
        return None

    def button(self) -> bool:
        """
        Returns true if any buttons are currently pressed

        :return: True if any buttons are currently pressed
        """
        return self.buttonApi.button_pressed_get()

    def flash(self) -> None:
        """
        The EV3 will flash the LEDs immediately.
        """
        leds = [
            LED(
                side="left",
                color="orange",
            )
        ]
        self.ledApi.led_flash_post(leds)

    def led_off(self) -> None:
        """
        This method will switch off the LEDs of from the EV3.
        """
        self.ledApi.led_off_post()
