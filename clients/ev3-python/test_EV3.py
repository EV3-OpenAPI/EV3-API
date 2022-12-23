from unittest import TestCase

from ev3.EV3 import EV3

hostname = "10.0.100.98:8080"


class TestEV3(TestCase):
    def test_beep(self):
        EV3(hostname).beep()

    def test_play_tone(self):
        EV3(hostname).play_tone(220, 1)

    def test_speak(self):
        EV3(hostname).speak("Hello World")

    def test_voltage(self):
        print(EV3(hostname).voltage())

    def test_current(self):
        print(EV3(hostname).current())

    def test_max_voltage(self):
        print(EV3(hostname).max_voltage())

    def test_min_voltage(self):
        print(EV3(hostname).min_voltage())

    def test_technology(self):
        print(EV3(hostname).technology())

    def test_button(self):
        print(EV3(hostname).button())

    def test_flash(self):
        EV3(hostname).flash()

    def test_led_off(self):
        EV3(hostname).led_off()
