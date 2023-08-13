from unittest import TestCase

from ev3.Buggy import Buggy

hostname = "10.0.100.98"


class TestBuggy(TestCase):
    def test_gyro(self):
        print(Buggy(hostname).gyro())

    def test_distance(self):
        print(Buggy(hostname).distance())

    def test_stop(self):
        Buggy(hostname).stop()

    def test_on(self):
        Buggy(hostname).on(20)

    def test_on_for_degrees(self):
        Buggy(hostname).on_for_degrees(20, 720)

    def test_on_for_rotations(self):
        Buggy(hostname).on_for_rotations(20, 2)

    def test_on_for_seconds(self):
        Buggy(hostname).on_for_seconds(20, 2)

    def test_steer_counts(self):
        Buggy(hostname).steer_counts(20, 2, 50)

    def test_steer_duration(self):
        Buggy(hostname).steer_duration(20, 2, -50)
