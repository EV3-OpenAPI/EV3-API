from unittest import TestCase

from ev3.Buggy import Buggy


class TestEV3(TestCase):
    def test_beep(self):
        Buggy("10.0.100.98:8080")
