package ch.zhaw.ev3;


import ch.zhaw.ev3api.api.MotorApi;
import ch.zhaw.ev3api.api.PowerApi;
import ch.zhaw.ev3api.api.SensorApi;
import ch.zhaw.ev3api.api.SoundApi;
import ch.zhaw.ev3api.invoker.ApiClient;
import ch.zhaw.ev3api.invoker.ApiException;
import ch.zhaw.ev3api.model.Text;
import ch.zhaw.ev3api.model.Tone;

import java.util.Objects;

public class EV3 {

    public enum Motors {
        A("A"),
        B("B"),
        C("C"),
        D("D");

        private final String name;

        Motors(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return name;
        }
    }
    private final String host_address;
    ApiClient apiClient;
    MotorApi motorApi;
    PowerApi powerApi;
    SensorApi sensorApi;
    SoundApi soundApi;

    /**
     * Creates a new EV3 with a specific ip-adress.
     * @param host_address the ip-adress
     */
    public EV3(String host_address) {
        this.host_address = host_address;
        this.apiClient = new ApiClient();
        this.apiClient.setBasePath(String.format("http://%s/api/v1", host_address));

        this.motorApi = new MotorApi(this.apiClient);
        this.powerApi = new PowerApi(this.apiClient);
        this.sensorApi = new SensorApi(this.apiClient);
        this.soundApi = new SoundApi(this.apiClient);
    }

    /**
     * This method returns the ip-adress from the EV3
     * @return the ip-adress
     */
    public String getHost_address() {
        return host_address;
    }

    /**
     * The EV3 will do a beep sound.
     */
    public void beep() {
        try {
            this.soundApi.soundBeepPost();
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    /**
     * The EV3 will play a tone.
     * @param frequency the specific frequenz for the tone
     * @param lengthMs the specific duration of the tone
     */
    public void play_tone(int frequency, int lengthMs) {
        try {
            this.soundApi.soundTonePost(new Tone().frequency(frequency).lengthMs(lengthMs));
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    /**
     * The EV3 will speak a specific text.
     * @param text the spoken text for the EV3
     */
    public void speak(String text) {
        try {
            this.soundApi.soundSpeakPost(new Text().text(text));
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    /**
     * This method always returns immediately, whether or not the battery voltage level exists.
     * @return the battery voltage level.
     */
    public int voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltage()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    /**
     * This method always returns immediately, whether or not the battery current level exists.
     * @return the battery current level.
     */
    public int current() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getCurrent()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    /**
     * This method always returns immediately, whether or not the maximal battery voltage exists.
     * @return the maximal battery voltage
     */
    public int max_voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltageMax()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    /**
     * This method always returns immediately, whether or not the minimal battery voltage exists.
     * @return the minimal battery voltage.
     */
    public int min_voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltageMin()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    /**
     * This method always returns immediately, whether or not the battery technology description exists.
     * @return the battery technology description
     */
    public String technology() {
        try {
            return this.powerApi.powerGet().getTechnology();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return null;
    }

    /**
     * This method returns an flag, if the button is pressed or not
     * @return the boolean if pressed or not
     */
    public boolean button() {
        // TODO: implement
        return false;
    }

    /**
     * The EV3 will flash the LEDs immediately.
     */
    public void flash() {
        // TODO: implement
    }

    /**
     * This method will set the LEDs of from the EV3
     */
    public void led() {
        // TODO: implement
    }

    /**
     * This method will switch off the LEDs of from the EV3.
     */
    public void led_off() {
        // TODO: implement
    }

    /**
     * This method will turn on the monitor from the EV3
     */
    public void monitor_on() {
        // TODO: implement
    }

    /**
     * This method will turn off the monitor from the EV3
     */
    public void monitor_off() {
        // TODO: implement
    }
}
