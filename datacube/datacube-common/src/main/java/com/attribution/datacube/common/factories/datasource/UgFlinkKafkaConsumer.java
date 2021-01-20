package com.attribution.datacube.common.factories.datasource;

import com.attribution.datacube.proto.tool.ProtoMessageSchema;
import com.google.protobuf.Message;
import org.apache.flink.api.common.serialization.DeserializationSchema;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;

import java.util.Properties;

public class UgFlinkKafkaConsumer<T extends Message> extends FlinkKafkaConsumer<T> {
    private final ProtoMessageSchema protoMessageSchema;

    public UgFlinkKafkaConsumer(String topic, DeserializationSchema<T> valueDeserializer, Properties props) {
        super(topic, valueDeserializer, props);
        this.protoMessageSchema = (ProtoMessageSchema) valueDeserializer;
    }

    @Override
    public void open(Configuration configuration) throws Exception {
        super.open(configuration);
        this.protoMessageSchema.initInvokeFunc();
        this.protoMessageSchema.initMetrics(getRuntimeContext());
    }
}
