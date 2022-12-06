package ch.zhaw.ev3;

import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.Motor;
import ch.zhaw.ev3api.model.MotorRequest;

import java.util.Arrays;

public class Buggy extends EV3 {
    private final Motors left;
    private final Motors right;

    private int maxSpeed;

    private Sonic sonic;

    private Gyro gyro;

    /**
     * Create a new buggy with a specific ip-adress
     * @param host_address the specific ip-adress
     * @throws ApiException
     */
    public Buggy(String host_address) throws ApiException {
        super(host_address);

        // TODO: verify default config
        this.left = Motors.A;
        this.right = Motors.D;

        this.maxSpeed = motorApi.motorTachoTypePortGet(String.valueOf(Motor.SizeEnum.L), right.toString()).getMaxSpeed();
    }

    /**
     * Create a new buggy with a specific ip_adress, a motor left object and a motor right object.
     * @param host_address the specific ip-adress
     * @param left the left motor
     * @param right the right motor
     * @throws ApiException
     */
    public Buggy(String host_address, Motors left, Motors right) throws ApiException {
        super(host_address);

        this.left = left;
        this.right = right;
        this.sonic = new Sonic(this.sensorApi);
        this.gyro = new Gyro(this.sensorApi);

        this.maxSpeed = motorApi.motorTachoTypePortGet(String.valueOf(Motor.SizeEnum.L), right.toString()).getMaxSpeed();
    }

    /**
     * Get the angle measured by the gyro sensor
     * @return the angle
     */
    public int gyro() {
        try {
            return gyro.getAngle();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    /**
     * Get the distance measured by the sonic-sensor,
     * @return the distance to the next object
     * @throws ApiException
     */
    public int distance() throws ApiException {
        return sonic.getDistance();
    }

    /**
     * This method stops the running motor immediately, whether the motor runs or not.
     * @throws ApiException
     */
    public void stop() throws ApiException {
        motorApi.motorStopAllPost();
    }

    /**
     * This method starts the motor with a specific speed,
     * the motor will run forever as long the battery has power.
     * @param speedPercent percent from the maxspeed
     * @throws ApiException
     */
    public void on(int speedPercent) throws ApiException {
        MotorRequest mr = getMotors();
        mr.setCommand(MotorRequest.CommandEnum.RUN_FOREVER);
        mr.setSpeed((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(mr);
    }

    /**
     * This method starts the motor with a specific speed,
     * the buggy will drives exactly as many seconds as stated
     * @param speedPercent maxspeed in percent
     * @param seconds how long the buggy should drive
     * @throws ApiException
     */
    public void onForSeconds(int speedPercent, double seconds) throws ApiException {
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
