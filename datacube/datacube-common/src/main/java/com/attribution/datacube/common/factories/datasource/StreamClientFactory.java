package com.attribution.datacube.common.factories.datasource;

import com.typesafe.config.Config;

public class StreamClientFactory {
    public static StreamClient getStreamClient(Config config) throws Exception {
        String type = config.getString("type");
        switch (type) {
            case "kafka": {
                return new KafkaStreamClient(config);
            }
            default: {
                throw new Exception("no such stream client");
            }
        }
    }
}
