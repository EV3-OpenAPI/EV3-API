import math

from ev3.EV3 import EV3
from ev3.sensors.Gyro import Gyro
from ev3.sensors.Sonic import Sonic
from ev3api.model.motor_steer_counts_post_request import MotorSteerCountsPostRequest
from ev3api.model.motor_steer_duration_post_request import MotorSteerDurationPostRequest
from ev3api.model.steering_unit import SteeringUnit
from ev3api.model.tacho_motor import TachoMotor


class Buggy(EV3):
    def __init__(self, host_address):
        super().__init__(host_address)

        self.left = TachoMotor(port="A", size="l")
        self.right = TachoMotor(port="D", size="l")

        self.steering_unit = SteeringUnit(left=self.left, right=self.right)
        self.driving_unit = [self.left, self.right]

        motor = self.motorApi.motor_tacho_get()[0]
        self.max_speed = motor["max_speed"]
        self.count_per_rot = motor["count_per_rot"]

        self.sonicSensor = Sonic(self.sensorApi)
        self.gyroSensor = Gyro(self.sensorApi)

    def gyro(self) -> int:
        """
        :return: the angle measured by the gyro sensor
        """
        return self.gyroSensor.get_angle()

    def distance(self) -> int:
        """
        :return: the distance measured by the sonic sensor
        """
        return self.sonicSensor.get_distance()

    def stop(self) -> None:
        """
        Stops all motors immediately
        """
        self.motorApi.motor_stop_all_post()

    def on(self, speed_percent: int) -> None:
        """
        Start the motors with the specified speed percentage of their max speed
        :param speed_percent: speed percentage to start the motors with
        """
        self.left["command"] = "run-forever"
        self.right["command"] = "run-forever"

        self.left["speed_setpoint"] = self.max_speed // 100 * speed_percent
        self.right["speed_setpoint"] = self.max_speed // 100 * speed_percent

        self.motorApi.motor_tacho_post(self.driving_unit)

    def _set_relative_position_setpoint(self, degrees: int, motor: TachoMotor) -> None:
        pos_delta = math.floor((degrees * self.count_per_rot) / 360)
        motor["position_setpoint"] = pos_delta

    def on_for_degrees(self, speed_percent: int, degrees: int) -> None:
        """
        Start the motors with the specified speed percentage of their max speed.
        Turn them off after they rotated for the amount of degrees given.
        :param speed_percent: speed percentage to start the motors with
        :param degrees: to turn before turning off
        """
        degrees = degrees * -1 if speed_percent < 0 else degrees
        self._set_relative_position_setpoint(degrees, self.left)
        self._set_relative_position_setpoint(degrees, self.right)

        self.left["command"] = "run-to-rel-pos"
        self.right["command"] = "run-to-rel-pos"

        self.left["speed_setpoint"] = self.max_speed // 100 * speed_percent
        self.right["speed_setpoint"] = self.max_speed // 100 * speed_percent

        self.motorApi.motor_tacho_post(self.driving_unit)

    def on_for_rotations(self, speed_percent: int, rotations: int) -> None:
        """
        Start the motors with the specified speed percentage of their max speed.
        Turn them off after they rotated for the amount of rotations given.
        :param speed_percent: speed percentage to start the motors with
        :param rotations: to turn before turning off
        """
        self._set_relative_position_setpoint(rotations * 360, self.left)
        self._set_relative_position_setpoint(rotations * 360, self.right)

        self.left["command"] = "run-to-rel-pos"
        self.right["command"] = "run-to-rel-pos"

        self.left["speed_setpoint"] = self.max_speed // 100 * speed_percent
        self.right["speed_setpoint"] = self.max_speed // 100 * speed_percent

        self.motorApi.motor_tacho_post(self.driving_unit)

    def on_for_seconds(self, speed_percent: int, seconds: float) -> None:
        """
        Start the motors with the specified speed percentage of their max speed.
        Turn them off after they rotated for the amount of seconds given.
        :param speed_percent: speed percentage to start the motors with
        :param seconds: to run before turning off
        """
        self.left["command"] = "run-timed"
        self.right["command"] = "run-timed"

        self.left["speed_setpoint"] = self.max_speed // 100 * speed_percent
        self.right["speed_setpoint"] = self.max_speed // 100 * speed_percent

        self.left["time_setpoint"] = seconds * 1000
        self.right["time_setpoint"] = seconds * 1000

        self.motorApi.motor_tacho_post(self.driving_unit)

    def steer_counts(self, speed_percent: int, counts: int, turn: int) -> None:
        """
        Turns the robot at the given speed for the given tacho counts
        to the given degree left or right.
        :param speed_percent: of the max speed
        :param counts: tacho counts to rotate before stopping the turn
        :param turn: from -100 (hard left) to +100 (hard right)
        """
        req = MotorSteerCountsPostRequest(
            steering_unit=self.steering_unit,
            speed=self.max_speed // 100 * speed_percent,
            counts=counts,
            turn=turn,
        )

        self.motorApi.motor_steer_counts_post(req)

    def steer_duration(
        self, speed_percent: int, duration_sec: float, turn: int
    ) -> None:
        """
        Turns the robot at the given speed for the given number of seconds
        to the given degree left or right.
        :param speed_percent: of the max speed
        :param duration_sec: to run before stopping the turn
        :param turn: from -100 (hard left) to +100 (hard right)
        """
        req = MotorSteerDurationPostRequest(
            steering_unit=self.steering_unit,
            speed=self.max_speed // 100 * speed_percent,
            duration_ms=duration_sec * 1000,
            turn=turn,
        )

        self.motorApi.motor_steer_duration_post(req)
