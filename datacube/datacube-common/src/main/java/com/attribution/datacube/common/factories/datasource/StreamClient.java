package com.attribution.datacube.common.factories.datasource;


import com.google.protobuf.Message;
import org.apache.flink.streaming.api.functions.source.RichParallelSourceFunction;

import java.util.Properties;

public abstract class StreamClient {
    protected Properties properties;

    protected String topic;
    protected String pbClassName;

    abstract public RichParallelSourceFunction<Message> getSourceFunction();
}
