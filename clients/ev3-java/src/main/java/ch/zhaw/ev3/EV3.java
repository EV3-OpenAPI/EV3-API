package ch.zhaw.ev3;


import ch.zhaw.ev3api.api.MotorApi;
import ch.zhaw.ev3api.api.PowerApi;
import ch.zhaw.ev3api.api.SensorApi;
import ch.zhaw.ev3api.api.SoundApi;
import ch.zhaw.ev3api.invoker.ApiClient;
import ch.zhaw.ev3api.invoker.ApiException;

public class EV3 {
    private String host_address;
    private ApiClient apiClient;
    private MotorApi motorApi;
    private PowerApi powerApi;
    private SensorApi sensorApi;
    private SoundApi soundApi;

    public EV3(String host_address) {
        this.host_address = host_address;
        this.apiClient = new ApiClient();
        this.apiClient.setBasePath(String.format("http://%s/api/v1", host_address));
        System.out.println(this.apiClient.getBasePath());

        this.motorApi = new MotorApi(this.apiClient);
        this.powerApi = new PowerApi(this.apiClient);
        this.sensorApi = new SensorApi(this.apiClient);
        this.soundApi = new SoundApi(this.apiClient);
    }

    public void beep() {
        try {
            this.soundApi.soundBeepPost();
        } catch (ApiException e) {
            System.out.println(e);
        }
    }
}
