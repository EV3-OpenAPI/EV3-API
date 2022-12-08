package ch.zhaw.ev3;

import ch.zhaw.ev3.motors.TachoMotor;
import ch.zhaw.ev3.sensors.Gyro;
import ch.zhaw.ev3.sensors.Sonic;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.Motor;
import ch.zhaw.ev3api.model.MotorRequest;

import java.util.Arrays;

public class Buggy extends EV3 {
    private final TachoMotor left;
    private final TachoMotor right;

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
        this.left = new TachoMotor(motorApi, TachoMotor.Port.A);
        this.right = new TachoMotor(motorApi, TachoMotor.Port.D);
    }

    /**
     * Create a new buggy with a specific ip_adress, a motor left object and a motor right object.
     * @param host_address the specific ip-adress
     * @param left the left motor
     * @param right the right motor
     * @throws ApiException
     */
    public Buggy(String host_address, TachoMotor left, TachoMotor right) throws ApiException {
        super(host_address);

        this.left = left;
        this.right = right;
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
        MotorRequest mr = getMotors();
        mr.setCommand(MotorRequest.CommandEnum.RUN_FOREVER);
        mr.setSpeed((int) (maxSpeed / 100.0 * speedPercent));
        motorApi.motorTachoPost(mr);
    }

    private void setRelativePositionSetpoint(int degrees, TachoMotor motor) throws ApiException {
        int pos_delta = (int) Math.round((degrees * motor.getCountsPerRot()) / 360.0);

    }

    public void onForDegrees(int speedPercent, int degrees) throws ApiException {
        degrees = speedPercent < 0 ? degrees * -1 : degrees;
        setRelativePositionSetpoint(degrees, left);
        setRelativePositionSetpoint(degrees, right);
        /*
        def _set_rel_position_degrees_and_speed_sp(self, degrees, speed):
            degrees = degrees if speed >= 0 else -degrees
            speed = abs(speed)

            position_delta = int(round((degrees * self.count_per_rot) / 360))
            speed_sp = int(round(speed))

            self.position_sp = position_delta
            self.speed_sp = speed_sp

        def _set_brake(self, brake):
            if brake:
                self.stop_action = self.STOP_ACTION_HOLD
            else:
                self.stop_action = self.STOP_ACTION_COAST

        def on_for_degrees(self, speed, degrees, brake=True, block=True):
            """
            Rotate the motor at ``speed`` for ``degrees``
            ``speed`` can be a percentage or a :class:`ev3dev2.motor.SpeedValue`
            object, enabling use of other units.
            """
            speed_sp = self._speed_native_units(speed)
            self._set_rel_position_degrees_and_speed_sp(degrees, speed_sp)
            self._set_brake(brake)
            self.run_to_rel_pos()

            if block:
                self.wait_until('running', timeout=WAIT_RUNNING_TIMEOUT)
                self.wait_until_not_moving()
         */
        MotorRequest mr = getMotors();
    }

    public void onForRotations(int speedPercent, int rotations) {
        // TODO: implement
        /*
        def on_for_rotations(self, speed, rotations, brake=True, block=True):
            """
            Rotate the motor at ``speed`` for ``rotations``
            ``speed`` can be a percentage or a :class:`ev3dev2.motor.SpeedValue`
            object, enabling use of other units.
            """
            speed_sp = self._speed_native_units(speed)
            self._set_rel_position_degrees_and_speed_sp(rotations * 360, speed_sp)
            self._set_brake(brake)
            self.run_to_rel_pos()

            if block:
                self.wait_until('running', timeout=WAIT_RUNNING_TIMEOUT)
                self.wait_until_not_moving()
        */
        MotorRequest mr = getMotors();
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
        mr.speed((int) (left.getMaxSpeed() / 100.0 * speedPercent));
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
