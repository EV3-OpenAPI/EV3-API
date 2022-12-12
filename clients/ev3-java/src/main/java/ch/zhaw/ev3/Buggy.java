package ch.zhaw.ev3;

import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.SteeringUnit;
import ch.zhaw.ev3api.model.TachoMotor;


import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Buggy extends EV3 {
    private final TachoMotor left;
    private final TachoMotor right;
    private SteeringUnit steeringUnit;
    private List<TachoMotor> drivingUnit;

    private Sonic sonic;

    private Gyro gyro;

    private int maxSpeed;
    private int countsPerRot;

    /**
     * Create a new buggy with a specific ip-address
     * @param host_address the specific ip-address
     * @throws ApiException
     */
    public Buggy(String host_address) throws ApiException {
        super(host_address);

        this.left = new TachoMotor().port(TachoMotor.PortEnum.A).size(TachoMotor.SizeEnum.L);
        this.right = new TachoMotor().port(TachoMotor.PortEnum.D).size(TachoMotor.SizeEnum.L);
        init();
    }

    /**
     * Create a new buggy with a specific ip_address, a motor left object and a motor right object.
     * @param host_address the specific ip-address
     * @param left the left motor
     * @param right the right motor
     * @throws ApiException
     */
    public Buggy(String host_address, TachoMotor left, TachoMotor right) throws ApiException {
        super(host_address);

        this.left = left;
        this.right = right;
        init();
    }

    private void init() throws ApiException {
        this.steeringUnit = new SteeringUnit().left(this.left).right(this.right);
        this.drivingUnit = Arrays.asList(this.left, this.right);

        TachoMotor motor = motorApi.motorTachoGet().get(0);
        this.maxSpeed = motor.getMaxSpeed();
        this.countsPerRot = motor.getCountPerRot();

        this.sonic = new Sonic(this.sensorApi);
        this.gyro = new Gyro(this.sensorApi);
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
        left.command(TachoMotor.CommandEnum.RUN_FOREVER).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_FOREVER).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(drivingUnit);
    }

    private void setRelativePositionSetpoint(int degrees, TachoMotor motor) {
        int pos_delta = (int) Math.round((degrees * countsPerRot) / 360.0);
        motor.positionSetpoint(pos_delta);
    }

    public void onForDegrees(int speedPercent, int degrees) throws ApiException {
        degrees = speedPercent < 0 ? degrees * -1 : degrees;
        setRelativePositionSetpoint(degrees, left);
        setRelativePositionSetpoint(degrees, right);

        left.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(drivingUnit);
    }

    public void onForRotations(int speedPercent, int rotations) throws ApiException {
        setRelativePositionSetpoint(rotations * 360, left);
        setRelativePositionSetpoint(rotations * 360, right);

        left.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(drivingUnit);

    }

    /**
     * This method starts the motor with a specific speed,
     * the buggy will drives exactly as many seconds as stated
     * @param speedPercent maxspeed in percent
     * @param seconds how long the buggy should drive
     * @throws ApiException
     */
    public void onForSeconds(int speedPercent, double seconds) throws ApiException {
        left.command(TachoMotor.CommandEnum.RUN_TIMED).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TIMED).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));

        left.timeSetpoint((int) (seconds * 1000));
        right.timeSetpoint((int) (seconds * 1000));

        System.out.println(drivingUnit);
        motorApi.motorTachoPost(drivingUnit);
    }

}
