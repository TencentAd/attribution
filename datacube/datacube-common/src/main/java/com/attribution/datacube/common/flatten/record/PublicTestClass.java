package com.attribution.datacube.common.flatten.record;

public class PublicTestClass {
    public String appId;

    public String convId;

    @Override
    public String toString() {
        return "PublicTestClass{" +
                "appId='" + appId + '\'' +
                ", convId='" + convId + '\'' +
                '}';
    }

    public PublicTestClass(String appId, String convId) {
        this.appId = appId;
        this.convId = convId;
    }
}
