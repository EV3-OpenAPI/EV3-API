package ch.zhaw.ev3;

import ch.zhaw.ev3api.invoker.ApiException;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

public class BuggyTest {

    private static Buggy buggy;

    @BeforeAll
    static void beforeAll() {
        buggy = new Buggy("192.168.1.251");
    }

    @Test
    void gyro() {
        System.out.println(buggy.gyro());
    }

    @Test
    void distance() {
        System.out.println(buggy.distance());
    }

    @Test
    void on() {
        buggy.on(20);
    }

    @Test
    void onForSeconds() {
        buggy.onForSeconds(20, 3.0);
    }

    @Test
    void stop() {
        buggy.stop();
    }

    @Test
    void onForDegrees() {
        buggy.onForDegrees(20, 360);
    }

    @Test
    void onForRotations() {
        buggy.onForRotations(20, 2);
    }

    @Test
    void steerCounts() {
        buggy.steerCounts(20, 2 * 360, 0);
    }

    @Test
    void steerDuration() {
        buggy.steerDuration(20, 4, 20);
    }

    @Test
    void startSleepStop() , InterruptedException {
        buggy.on(5);

        Thread.sleep(3000);

        buggy.stop();
    }

    @Test
    void testOn() throws InterruptedException {
        buggy.on(buggy.left, 10);
    }

    @Test
    void testOnForDegrees() {
        buggy.onForDegrees(buggy.left, 10, 360);
    }

    @Test
    void testOnForRotations() {
        buggy.onForRotations(buggy.left, 10, 2);
    }

    @Test
    void testOnForSeconds() {
        buggy.onForSeconds(buggy.left, 10, 3);
    }
}
