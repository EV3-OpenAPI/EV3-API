from unittest import TestCase

from ev3.EV3 import EV3


class TestEV3(TestCase):
    def test_beep(self):
        EV3("10.0.100.98:8080").beep()

