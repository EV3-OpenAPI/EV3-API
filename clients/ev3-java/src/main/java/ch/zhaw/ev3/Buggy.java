package ch.zhaw.ev3;

import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.Motor;
import ch.zhaw.ev3api.model.MotorRequest;

import java.util.Arrays;

public class Buggy extends EV3 {
    private Motors left;
    private Motors right;

    private int maxSpeed;

    private Sonic sonic;

    private Gyro gyro;

    public Buggy(String host_address) throws ApiException {
        super(host_address);

        // TODO: verify default config
        this.left = Motors.A;
        this.right = Motors.D;

        this.maxSpeed = motorApi.motorTachoTypePortGet(String.valueOf(Motor.SizeEnum.L), right.toString()).getMaxSpeed();
    }

    public Buggy(String host_address, Motors left, Motors right) throws ApiException {
        super(host_address);

        this.left = left;
        this.right = right;
        this.sonic = new Sonic(this.sensorApi);
        this.gyro = new Gyro(this.sensorApi);

        this.maxSpeed = motorApi.motorTachoTypePortGet(String.valueOf(Motor.SizeEnum.L), right.toString()).getMaxSpeed();
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

    public void on(int speedPercent) throws ApiException {
        MotorRequest mr = getMotors();
        mr.setCommand(MotorRequest.CommandEnum.RUN_FOREVER);
        mr.setSpeed((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(mr);
    }



    public void onForSeconds(int speedPercent, double seconds) throws ApiException {
        // TODO: implement
        MotorRequest mr = getMotors();
        mr.command(MotorRequest.CommandEnum.RUN_TIMED);
        mr.time((int) (seconds * 1000));
        mr.speed((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(mr);
    }

    private MotorRequest getMotors() {
        MotorRequest mr = new MotorRequest();
        Motor l = new Motor().port(String.valueOf(left)).size(Motor.SizeEnum.L); // FIXME
        Motor r = new Motor().port(String.valueOf(right)).size(Motor.SizeEnum.L);
        mr.motors(Arrays.asList(l, r));
        return mr;
    }
}
