package ch.zhaw.ev3;

import org.junit.jupiter.api.Test;

class EV3Test {

    @Test
    void beep() {
        new EV3("10.0.100.98:8080").beep();
    }
}
