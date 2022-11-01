package org.example;

import ch.zhaw.ev3_api.client.invoker.ApiClient;
import ch.zhaw.ev3_api.client.api.MotorApi;
import ch.zhaw.ev3_api.client.invoker.ApiException;
import ch.zhaw.ev3_api.client.model.Motor;
import ch.zhaw.ev3_api.client.model.MotorRequest;

import java.io.IOException;
import java.util.Arrays;

public class Main {

    public static void main(String[] args) {
        ApiClient defaultClient = new ApiClient();
        defaultClient.setBasePath("http://10.0.100.98:8080/api/v1");

        MotorApi apiInstance = new MotorApi(defaultClient);
        try {
            Motor a = new Motor().size("l").port("A");
            Motor d = new Motor().size("l").port("D");
            MotorRequest mr = new MotorRequest().motors(Arrays.asList(a, d));

            apiInstance.motorTachoPost(mr);

            Thread.sleep(200);

            apiInstance.motorStopAllPost();
        } catch (InterruptedException | ApiException e) {
            throw new RuntimeException(e);
        }
    }
}
