package ch.zhaw.ev3.sensors;

import ch.zhaw.ev3api.api.SensorApi;
import ch.zhaw.ev3api.invoker.ApiException;

public class Gyro extends Sensor {
    public enum Modes {
        GYRO_ANG("GYRO-ANG"),
        GYRO_RATE("GYRO-RATE"),
        GYRO_FAS("GYRO-FAS"),
        GYRO_GaA("GYRO-G&A"),
        GYRO_CAL("GYRO-CAL"),
        TILT_RATE("TILT-RATE"),
        TILT_ANG("TILT-ANG");

        private final String mode;

        Modes(String mode) {
            this.mode = mode;
        }

        @Override
        public String toString() {
            return mode;
        }
    }

    public Gyro(SensorApi sensorApi) throws ApiException {
        super(Drivers.GYRO, sensorApi);
    }

    public int getAngle() throws ApiException {
        setMode(Modes.GYRO_ANG.mode);
        return Integer.parseInt(getValues().get(0));
    }

    public void reset() {
    }
}
