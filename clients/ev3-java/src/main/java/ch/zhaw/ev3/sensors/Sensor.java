package ch.zhaw.ev3.sensors;

import ch.zhaw.ev3api.api.SensorApi;
import ch.zhaw.ev3api.invoker.ApiException;

import java.util.List;

public class Sensor {
    public enum Drivers {
        GYRO("gyro"),
        SONIC("us");

        private final String name;

        Drivers(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return name;
        }
    }

    final SensorApi api;
    private final Drivers driver;
    private final List<String> modes;

    private ch.zhaw.ev3api.model.Sensor sensor;

    public Sensor(Drivers driver, SensorApi sensorApi) throws ApiException {
        this.driver = driver;
        this.api = sensorApi;
        updateSensor();
        this.modes = sensor.getModes();
    }

    void updateSensor() throws ApiException {
        sensor = api.sensorTypeGet(driver.name);
    }

    public Drivers getDriver() {
        return driver;
    }

    List<String> getModes() {
        return modes;
    }

    String getMode() throws ApiException {
        updateSensor();
        return sensor.getMode();
    }

    void setMode(String mode) throws ApiException {
        if (!modes.contains(mode)) {
            return; // TODO: error handling
        }
        if (getMode().equals(mode)) {
            return; // No need to change mode
        }
        api.sensorTypePut(driver.name, new ch.zhaw.ev3api.model.Sensor().mode(mode));
    }

    List<String> getValues() throws ApiException {
        return api.sensorTypeValuesGet(driver.name);
    }
}
