from enum import Enum

import ev3api.model.sensor
from ev3api.api.sensor_api import SensorApi


class Sensor:
    class Drivers(Enum):
        GYRO = "gyro"
        SONIC = "us"

    def __int__(self, driver: Drivers, sensor_api: SensorApi):
        self.driver = driver
        self.api = sensor_api

        self.sensor = self.api.sensor_type_get(driver)
        print(self.sensor)
        self.modes = self.sensor["modes"]

    def get_mode(self) -> str:
        return self.sensor["mode"]

    def set_mode(self, mode: str) -> None:
        if mode not in self.modes:
            return

        if mode == self.get_mode():
            return

        self.api.sensor_type_put(self.driver, ev3api.model.sensor.Sensor(mode=mode))

    def get_values(self) -> [str]:
        return self.api.sensor_type_values_get(self.driver)
