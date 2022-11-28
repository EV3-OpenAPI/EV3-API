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
        A,
        B,
        C,
        D
    }
    private final String host_address;
    ApiClient apiClient;
    MotorApi motorApi;
    PowerApi powerApi;
    SensorApi sensorApi;
    SoundApi soundApi;

    public EV3(String host_address) {
        this.host_address = host_address;
        this.apiClient = new ApiClient();
        this.apiClient.setBasePath(String.format("http://%s/api/v1", host_address));

        this.motorApi = new MotorApi(this.apiClient);
        this.powerApi = new PowerApi(this.apiClient);
        this.sensorApi = new SensorApi(this.apiClient);
        this.soundApi = new SoundApi(this.apiClient);
    }

    public String getHost_address() {
        return host_address;
    }

    public void beep() {
        try {
            this.soundApi.soundBeepPost();
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    public void play_tone(int frequency, int lengthMs) {
        try {
            this.soundApi.soundTonePost(new Tone().frequency(frequency).lengthMs(lengthMs));
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    public void speak(String text) {
        try {
            this.soundApi.soundSpeakPost(new Text().text(text));
        } catch (ApiException e) {
            System.out.println(e);
        }
    }

    public int voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltage()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    public int current() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getCurrent()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    public int max_voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltageMax()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    public int min_voltage() {
        try {
            return Objects.requireNonNull(this.powerApi.powerGet().getVoltageMin()).intValue();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return -1;
    }

    public String technology() {
        try {
            return this.powerApi.powerGet().getTechnology();
        } catch (ApiException e) {
            System.out.println(e);
        }

        return null;
    }

    public boolean button() {
        // TODO: implement
        return false;
    }

    public void flash() {
        // TODO: implement
    }

    public void led() {
        // TODO: implement
    }

    public void led_off() {
        // TODO: implement
    }

    public void monitor_on() {
        // TODO: implement
    }

    public void monitor_off() {
        // TODO: implement
    }
}
