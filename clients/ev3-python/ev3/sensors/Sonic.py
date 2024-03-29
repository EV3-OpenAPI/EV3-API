from enum import Enum
from ev3.sensors.Sensor import Sensor
from ev3api.api.sensor_api import SensorApi


class Sonic(Sensor):
    class MODES(Enum):
        DIST_CM = "US-DIST-CM"
        DIST_IN = "US-DIST-IN"
        LISTEN = "US-LISTEN"
        SI_CM = "US-SI-CM"
        SI_IN = "US-SI-IN"
        DC_CM = "US-DC-CM"
        DC_IN = "US-DC-IN"

    def __init__(self, sensor_api: SensorApi):
        super(Sonic, self).__init__(sensor_api, Sensor.Drivers.SONIC)

    def get_distance(self) -> int:
        self.set_mode(self.MODES.DIST_CM.value)
        return int(self.get_values()[0])
