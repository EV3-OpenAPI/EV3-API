package ch.zhaw.ev3;


import ch.zhaw.ev3api.api.*;
import ch.zhaw.ev3api.invoker.ApiClient;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.LED;
import ch.zhaw.ev3api.model.Text;
import ch.zhaw.ev3api.model.Tone;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class EV3 {

    private final String host_address;
    ApiClient apiClient;
    MotorApi motorApi;
    PowerApi powerApi;
    SensorApi sensorApi;
    SoundApi soundApi;
    ButtonApi buttonApi;
    LedApi ledApi;

    /**
     * Creates a new EV3 with a specific ip-address.
     * @param host_address the ip-address
     */
    public EV3(String host_address) {
        this.host_address = host_address;
        this.apiClient = new ApiClient();
        this.apiClient.setBasePath(String.format("http://%s/api/v1", host_address));

        this.motorApi = new MotorApi(this.apiClient);
        this.powerApi = new PowerApi(this.apiClient);
        this.sensorApi = new SensorApi(this.apiClient);
        this.soundApi = new SoundApi(this.apiClient);
        this.buttonApi = new ButtonApi(this.apiClient);
        this.ledApi = new LedApi(this.apiClient);
    }

    /**
     * This method returns the ip-address from the EV3
     * @return the ip-address
     */
    public String getHost_address() {
        return host_address;
    }

    /**
     * The EV3 will beep.
     */
    public void beep() {
        try {
            soundApi.soundBeepPost();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * The EV3 will play a tone.
     * @param frequency the frequency for the tone
     * @param lengthMs the duration of the tone
     */
    public void play_tone(int frequency, int lengthMs) {
        try {
            soundApi.soundTonePost(new Tone().frequency(frequency).lengthMs(lengthMs));
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * The EV3 will speak a specific text.
     * @param text the spoken text for the EV3
     */
    public void speak(String text) {
        try {
            soundApi.soundSpeakPost(new Text().text(text));
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the battery voltage level exists.
     * @return the battery voltage level.
     */
    public int voltage() {
        try {
            return Objects.requireNonNull(powerApi.powerGet().getVoltage()).intValue();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the battery current level exists.
     * @return the battery current level.
     */
    public int current() {
        try {
            return Objects.requireNonNull(powerApi.powerGet().getCurrent()).intValue();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the maximal battery voltage exists.
     * @return the maximal battery voltage
     */
    public int max_voltage() {
        try {
            return Objects.requireNonNull(powerApi.powerGet().getVoltageMax()).intValue();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the minimal battery voltage exists.
     * @return the minimal battery voltage.
     */
    public int min_voltage() {
        try {
            return Objects.requireNonNull(powerApi.powerGet().getVoltageMin()).intValue();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the battery technology description exists.
     * @return the battery technology description
     */
    public String technology() {
        try {
            return powerApi.powerGet().getTechnology();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method returns a flag, if the button is pressed or not
     * @return the boolean if pressed or not
     */
    public boolean button() {
        try {
            List<String> pressedButtons = buttonApi.buttonPressedGet();
            return pressedButtons.size() == 0;
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * The EV3 will flash the LEDs immediately.
     */
    public void flash() {
        List<LED> leds = new ArrayList<>();
        leds.add(new LED().side("left").color(LED.ColorEnum.ORANGE));
        leds.add(new LED().side("right").color(LED.ColorEnum.ORANGE));

        try {
            ledApi.ledFlashPost(leds);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * The EV3 will flash the LEDs immediately.
     */
    public void flash(int red, int green) {
        List<LED> leds = new ArrayList<>();
        leds.add(new LED().side("left").red(red).green(green));
        leds.add(new LED().side("right").red(red).green(green));

        try {
            ledApi.ledFlashPost(leds);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * The EV3 will flash the LEDs immediately.
     */
    public void flash(LED.ColorEnum color) {
        List<LED> leds = new ArrayList<>();
        leds.add(new LED().side("left").color(color));
        leds.add(new LED().side("right").color(color));

        try {
            ledApi.ledFlashPost(leds);
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * This method will switch off the LEDs of from the EV3.
     */
    public void led_off() {
        try {
            ledApi.ledOffPost();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }
}
