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

        private final String name;

        Modes(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return name;
        }
    }

    private int offset = 0;

    /**
     * Create a new Gyro object with the sensor
     * @param sensorApi
     * @throws ApiException
     */
    public Gyro(SensorApi sensorApi) throws ApiException {
        super(Drivers.GYRO, sensorApi);
    }

    /**
     * This method will return the angle, which will be measured by the Gyro-sensor.
     * @return the measured angle
     * @throws ApiException
     */
    public int getAngle() throws ApiException {
        setMode(Modes.GYRO_ANG.name);
        return Integer.parseInt(getValues().get(0)) - offset;
    }

    /**
     * Reset offset to the measured angel value.
     * @throws ApiException
     */
    public void reset() throws ApiException {
        offset = getAngle();
    }
}
