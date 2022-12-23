package ch.zhaw.ev3;

import ch.zhaw.ev3api.invoker.ApiException;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

public class BuggyTest {

    private static Buggy buggy;

    @BeforeAll
    static void beforeAll() throws ApiException {
        buggy = new Buggy("10.0.100.98:8080");
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
    void onForDegrees() throws ApiException {
        buggy.onForDegrees(20, 360);
    }

    @Test
    void onForRotations() throws ApiException {
        buggy.onForRotations(20, 2);
    }

    @Test
    void steerCounts() throws ApiException {
        buggy.steerCounts(20, 2*360, 0);
    }

    @Test
    void steerDuration() throws ApiException {
        buggy.steerDuration(20, 4, 20);
    }

    @Test
    void startSleepStop() throws ApiException, InterruptedException {
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
