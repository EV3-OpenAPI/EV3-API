package ch.zhaw.ev3;

import ch.zhaw.ev3api.model.LED;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

class EV3Test {
    private static EV3 ev3;

    @BeforeAll
    static void beforeAll() {
        ev3 = new EV3("10.0.100.98");
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

    @Test
    void voltage() {
        System.out.println(ev3.voltage());
    }

    @Test
    void current() {
        System.out.println(ev3.current());
    }

    @Test
    void max_voltage() {
        System.out.println(ev3.max_voltage());
    }

    @Test
    void min_voltage() {
        System.out.println(ev3.min_voltage());
    }

    @Test
    void technology() {
        System.out.println(ev3.technology());
    }

    @Test
    void button() {
        System.out.println(ev3.button());
    }

    @Test
    void flash() {
        ev3.flash();
    }

    @Test
    void flashRed() {
        ev3.flash(LED.ColorEnum.RED);
    }

    @Test
    void flashGreen() {
        ev3.flash(LED.ColorEnum.GREEN);
    }

    @Test
    void flash128128() {
        ev3.flash(128, 128);
    }

    @Test
    void led_off() {
        ev3.led_off();
    }
}
