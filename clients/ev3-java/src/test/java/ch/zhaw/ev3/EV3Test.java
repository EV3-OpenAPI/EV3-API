package ch.zhaw.ev3;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

class EV3Test {
    private static EV3 ev3;

    @BeforeAll
    static void beforeAll() {
        ev3 = new EV3("160.85.143.143:8080");
    }

    @Test
    void beep() {
        ev3.beep();
    }

    @Test
    void play_tone() {
        ev3.play_tone(220, 1000);
    }

    @Test
    void speak() {
        ev3.speak("Hello World");
    }
}
