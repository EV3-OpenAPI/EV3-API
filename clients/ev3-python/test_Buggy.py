from unittest import TestCase

from ev3.Buggy import Buggy


class TestBuggy(TestCase):
    def test_on_for_seconds(self):
        Buggy("10.0.100.98:8080").on_for_seconds(20, 2)
