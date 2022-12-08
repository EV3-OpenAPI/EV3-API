package ch.zhaw.ev3;

import ch.zhaw.ev3.motors.TachoMotor;
import ch.zhaw.ev3api.invoker.ApiException;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

public class BuggyTest {

    private static Buggy buggy;

    @BeforeAll
    static void beforeAll() throws ApiException {
        buggy = new Buggy("160.85.132.65:8080");
    }

    @Test
    void gyro() {
        System.out.println(buggy.gyro());
    }

    @Test
    void distance() throws ApiException {
        System.out.println(buggy.distance());
    }

    @Test
    void on() throws ApiException {
        buggy.on(20);
    }

    @Test
    void onForSeconds() throws ApiException {
        buggy.onForSeconds(20, 3.0);
    }

    @Test
    void stop() throws ApiException {
        buggy.stop();
    }

    @Test
    void startSleepStop() throws ApiException, InterruptedException {
        buggy.on(5);

        Thread.sleep(3000);

        buggy.stop();
    }
}
