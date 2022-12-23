from enum import StrEnum

from ev3.sensors.Sensor import Sensor
from ev3api.api.sensor_api import SensorApi


class Gyro(Sensor):
    class Modes(StrEnum):
        GYRO_ANG = "GYRO-ANG"
        GYRO_RATE = "GYRO-RATE"
        GYRO_FAS = "GYRO-FAS"
        GYRO_GaA = "GYRO-G&A"
        GYRO_CAL = "GYRO-CAL"
        TILT_RATE = "TILT-RATE"
        TILT_ANG = "TILT-ANG"

    def __init__(self, sensor_api: SensorApi):
        super().__init__(sensor_api)

        self.offset = 0

    def get_angle(self) -> int:
        super().set_mode(self.Modes.GYRO_ANG)
        return int(self.get_values()[0]) - self.offset

    def reset(self) -> None:
        offset = self.get_angle()
