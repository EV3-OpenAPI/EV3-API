package ch.zhaw.ev3.sensors;

import ch.zhaw.ev3api.api.SensorApi;
import ch.zhaw.ev3api.invoker.ApiException;

public class Sonic extends Sensor {
    public enum Modes {
        DIST_CM("US-DIST-CM"),
        DIST_IN("US-DIST-IN"),
        LISTEN("US-LISTEN"),
        SI_CM("US-SI-CM"),
        SI_IN("US-SI-IN"),
        DC_CM("US-DC-CM"),
        DC_IN("US-DC-IN");

        private final String mode;

        Modes(String mode) {
            this.mode = mode;
        }


        @Override
        public String toString() {
            return mode;
        }
    }

    public Sonic(SensorApi sensorApi) throws ApiException {
        super(Drivers.SONIC, sensorApi);
    }

    public int getDistance() throws ApiException {
        setMode(Modes.DIST_CM.mode);
        return Integer.parseInt(getValues().get(0));
    }
}
