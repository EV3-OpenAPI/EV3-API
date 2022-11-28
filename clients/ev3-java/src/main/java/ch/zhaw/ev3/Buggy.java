package ch.zhaw.ev3;

import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.Text;

public class Buggy extends EV3 {
    private Motors left;
    private Motors right;

    private Sonic sonic;

    private Gyro gyro;

    public Buggy(String host_address) {
        super(host_address);

        // TODO: verify default config
        this.left = Motors.A;
        this.right = Motors.D;
    }

    public Buggy(String host_address, Motors left, Motors right) throws ApiException {
        super(host_address);
        this.sonic = new Sonic(this.sensorApi);
        this.gyro = new Gyro(this.sensorApi);
    }

    public int gyro() {
        try {
            return gyro.getAngle();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    public int distance() throws ApiException {
        return sonic.getDistance();
    }

    public void stop() throws ApiException {
        motorApi.motorStopAllPost();
    }
}
