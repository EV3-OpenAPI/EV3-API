package ch.zhaw.ev3;

import org.junit.jupiter.api.Test;

class EV3Test {
    private static EV3 ev3;

    @Test
    void beep() {
        new EV3("10.0.100.98:8080").beep();
    }

    @Test
    void play_tone() {
        new EV3("10.0.100.98:8080").play_tone(220, 1000);
    }

    @Test
    void speak() {
        new EV3("10.0.100.98:8080").speak("Hello World");
    }
}
