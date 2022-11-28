package ch.zhaw.ev3;

import ch.zhaw.ev3api.invoker.ApiException;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

public class BuggyTest {

    private static Buggy buggy;

    @BeforeAll
    static void beforeAll() throws ApiException {
        buggy = new Buggy("10.0.100.98:8080", EV3.Motors.A, EV3.Motors.D);
    }

    @Test
    void gyro() {
        System.out.println(buggy.gyro());
    }

    @Test
    void distance() throws ApiException {
        System.out.println(buggy.distance());
    }
}
