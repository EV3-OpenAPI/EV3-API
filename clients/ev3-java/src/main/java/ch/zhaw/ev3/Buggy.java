package ch.zhaw.ev3;

import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.MotorSteerCountsPostRequest;
import ch.zhaw.ev3api.model.MotorSteerDurationPostRequest;
import ch.zhaw.ev3api.model.SteeringUnit;
import ch.zhaw.ev3api.model.TachoMotor;


import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Buggy extends EV3 {
    public final TachoMotor left;
    public final TachoMotor right;
    private SteeringUnit steeringUnit;
    private List<TachoMotor> drivingUnit;

    private Sonic sonic;

    private Gyro gyro;

    private int maxSpeed;
    private int countsPerRot;

    /**
     * Create a new buggy with a specific ip-address
     * @param host_address the specific ip-address
     */
    public Buggy(String host_address) {
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
     */
    public Buggy(String host_address, TachoMotor left, TachoMotor right) {
        super(host_address);

        this.left = left;
        this.right = right;
        init();
    }

    private void init() {
        this.steeringUnit = new SteeringUnit().left(this.left).right(this.right);
        this.drivingUnit = Arrays.asList(this.left, this.right);

        try {
            TachoMotor motor = motorApi.motorTachoGet().get(0);
            this.maxSpeed = motor.getMaxSpeed();
            this.countsPerRot = motor.getCountPerRot();

            this.sonic = new Sonic(this.sensorApi);
            this.gyro = new Gyro(this.sensorApi);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
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
     */
    public int distance() {
        try {
            return sonic.getDistance();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method stops the running motor immediately, whether the motor runs or not.
     */
    public void stop() {
        try {
            motorApi.motorStopAllPost();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method starts the motor with a specific speed,
     * the motor will run forever as long the battery has power.
     * @param speedPercent of the max speed
     */
    public void on(int speedPercent) {
        left.command(TachoMotor.CommandEnum.RUN_FOREVER).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_FOREVER).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));

        try {
            motorApi.motorTachoPost(drivingUnit);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method starts the given motor with a specific speed,
     * the motor will run forever as long the battery has power.
     * @param speedPercent of the max speed
     */
    public void on(TachoMotor side, int speedPercent) {
        side.command(TachoMotor.CommandEnum.RUN_FOREVER).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        List<TachoMotor> motors = Arrays.asList(side);

        try {
            motorApi.motorTachoPost(motors);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    private void setRelativePositionSetpoint(int degrees, TachoMotor motor) {
        int pos_delta = (int) Math.round((degrees * countsPerRot) / 360.0);
        motor.positionSetpoint(pos_delta);
    }

    /**
     * This method turns on the motors with the given speed percent.
     * They will be on until they rotated for the specified degrees.
     * @param speedPercent of the max speed
     * @param degrees to turn before turning off
     */
    public void onForDegrees(int speedPercent, int degrees) {
        degrees = speedPercent < 0 ? degrees * -1 : degrees;
        setRelativePositionSetpoint(degrees, left);
        setRelativePositionSetpoint(degrees, right);

        left.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));

        try {
            motorApi.motorTachoPost(drivingUnit);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method turns on the given motor with the given speed percent.
     * They will be on until they rotated for the specified degrees.
     * @param speedPercent of the max speed
     * @param degrees to turn before turning off
     */
    public void onForDegrees(TachoMotor side, int speedPercent, int degrees) {
        degrees = speedPercent < 0 ? degrees * -1 : degrees;
        setRelativePositionSetpoint(degrees, side);

        side.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        List<TachoMotor> motors = Arrays.asList(side);

        try {
            motorApi.motorTachoPost(motors);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method turns on the motors with the given speed percent.
     * They will be on until they rotated for the specified number of rotations.
     * @param speedPercent of the max speed
     * @param rotations to turn before turning off
     */
    public void onForRotations(int speedPercent, int rotations) {
        setRelativePositionSetpoint(rotations * 360, left);
        setRelativePositionSetpoint(rotations * 360, right);

        left.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));

        try {
            motorApi.motorTachoPost(drivingUnit);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

    }

    /**
     * This method turns on the given motor with the given speed percent.
     * They will be on until they rotated for the specified number of rotations.
     * @param speedPercent of the max speed
     * @param rotations to turn before turning off
     */
    public void onForRotations(TachoMotor side, int speedPercent, int rotations) {
        setRelativePositionSetpoint(rotations * 360, side);

        side.command(TachoMotor.CommandEnum.RUN_TO_REL_POS).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        List<TachoMotor> motors = Arrays.asList(side);

        try {
            motorApi.motorTachoPost(motors);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

    }

    /**
     * This method turns on the motors with the given speed percent.
     * They will be on until they ran for the specified amount of seconds.
     * @param speedPercent of the max speed
     * @param seconds to run before turning off
     */
    public void onForSeconds(int speedPercent, double seconds) {
        left.command(TachoMotor.CommandEnum.RUN_TIMED).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        right.command(TachoMotor.CommandEnum.RUN_TIMED).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));

        left.timeSetpoint((int) (seconds * 1000));
        right.timeSetpoint((int) (seconds * 1000));

        try {
            motorApi.motorTachoPost(drivingUnit);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method turns on the given motor with the given speed percent.
     * They will be on until they ran for the specified amount of seconds.
     * @param speedPercent of the max speed
     * @param seconds to run before turning off
     */
    public void onForSeconds(TachoMotor side, int speedPercent, double seconds) {
        side.command(TachoMotor.CommandEnum.RUN_TIMED).speedSetpoint((int) (maxSpeed / 100.0 * speedPercent));
        List<TachoMotor> motors = Arrays.asList(side);

        side.timeSetpoint((int) (seconds * 1000));

        try {
            motorApi.motorTachoPost(motors);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * Turns the robot at the given speed for the given tacho counts
     * to the given degree left or right.
     * @param speedPercent of the max speed
     * @param counts tacho counts to rotate before stopping the turn
     * @param turn from -100 (hard left) to +100 (hard right)
     */
    public void steerCounts(int speedPercent, int counts, int turn) {
        MotorSteerCountsPostRequest r = new MotorSteerCountsPostRequest();
        r.steeringUnit(steeringUnit);
        r.setSpeed((int) (maxSpeed / 100.0 * speedPercent));
        r.setCounts(counts);
        r.setTurn(turn);

        try {
            motorApi.motorSteerCountsPost(r);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * Turns the robot at the given speed for the given number of seconds
     * to the given degree left or right.
     * @param speedPercent of the max speed
     * @param durationSec to run before stopping the turn
     * @param turn from -100 (hard left) to +100 (hard right)
     */
    public void steerDuration(int speedPercent, double durationSec, int turn) {
        MotorSteerDurationPostRequest r = new MotorSteerDurationPostRequest();
        r.steeringUnit(steeringUnit);
        r.setSpeed((int) (maxSpeed / 100.0 * speedPercent));
        r.setDurationMs((int) (durationSec * 1000));
        r.setTurn(turn);

        try {
            motorApi.motorSteerDurationPost(r);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

}
