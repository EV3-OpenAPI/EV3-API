package ch.zhaw.ev3;


import ch.zhaw.ev3api.api.PowerApi;
import ch.zhaw.ev3api.api.SoundApi;
import ch.zhaw.ev3api.invoker.ApiClient;
import ch.zhaw.ev3api.invoker.ApiException;

public class EV3 {
    private String host_address;

    public EV3(String host_address) {
        this.host_address = host_address;
    }

    public void beep() throws ApiException {
        ApiClient defaultClient = new ApiClient();
        defaultClient.setBasePath("http://10.0.100.98:8080/api/v1");
        SoundApi api = new SoundApi(defaultClient);
        api.soundBeepPost();
    }
}
