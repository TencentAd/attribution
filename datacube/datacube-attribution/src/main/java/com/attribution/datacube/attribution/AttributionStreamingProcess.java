package com.attribution.datacube.attribution;

import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.conv.Conv;
import com.twitter.chill.protobuf.ProtobufSerializer;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.api.common.serialization.SimpleStringSchema;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer;

import java.util.Properties;

public class AttributionStreamingProcess {

    public static void main(String[] args) throws Exception {
        String env = "test";
        String jobName = "attribution";

        // todo 这里的config需要增加配置
        Config config = ConfigFactory.load("attribution-streaming.conf")
                .getConfig(env).getConfig(jobName);

        Config streamConfig = config.getConfig("stream-config");

//        int checkpointInterval = config.getInt("checkpoint-interval");
//        String checkpointDir = config.getString("checkpoint-dir");

        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();
        see.getConfig().registerTypeWithKryoSerializer(Conv.ConversionLog.class, ProtobufSerializer.class);

        Properties producerProperties = new Properties();
        producerProperties.setProperty("bootstrap.servers", "localhost:9810");
        FlinkKafkaProducer<String> producer = new FlinkKafkaProducer<>("flink_result_test", new SimpleStringSchema(), producerProperties);

        DataStreamSource<Message> messageDataStreamSource = see
                .addSource(StreamClientFactory.getStreamClient(streamConfig).getSourceFunction());

        // todo 中间的处理逻辑
        messageDataStreamSource.map(new MapFunction<Message, String>() {
            @Override
            public String map(Message message) throws Exception {
                Conv.ConversionLog conversionLog = (Conv.ConversionLog) message;
                System.out.println(conversionLog.getAppId());
                return conversionLog.getAppId();
            }
        }).addSink(producer);

        see.execute("test kafka");
    }
}
