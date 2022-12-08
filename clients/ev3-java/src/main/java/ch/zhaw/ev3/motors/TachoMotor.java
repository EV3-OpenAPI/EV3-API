package ch.zhaw.ev3.motors;

import ch.zhaw.ev3api.api.MotorApi;
import ch.zhaw.ev3api.invoker.ApiException;

public class TachoMotor {
    public enum Port {
        A,
        B,
        C,
        D
    }

    public enum Size {
        s,
        m,
        l
    }

    private final MotorApi api;

    private final Port port;
    private final Size size;

    private int maxSpeed = -1;
    private int countsPerRot = -1;

    public TachoMotor(MotorApi api, Port port) {
        this.api = api;
        this.port = port;
        this.size = Size.l;
    }

    public TachoMotor(MotorApi api, Port port, Size size) {
        this.api = api;
        this.port = port;
        this.size = size;
    }

    public int getMaxSpeed() throws ApiException {
        if (maxSpeed == -1) {
            maxSpeed = api.motorTachoTypePortGet(size.toString(), port.toString()).getMaxSpeed();
        }
        return maxSpeed;
    }

    public int getCountsPerRot() throws ApiException {
        if (countsPerRot == -1) {
            countsPerRot = api.motorTachoTypePortGet(size.toString(), port.toString()).getCountPerRot();
        }
        return countsPerRot;
    }
}
