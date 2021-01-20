package com.attribution.datacube.attribution;

import com.twitter.chill.protobuf.ProtobufSerializer;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import com.attribution.datacube.common.vars.ConfigVars;
import org.apache.flink.api.common.serialization.SimpleStringSchema;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer;

import java.util.Properties;

public class StringStreamingProcess {
    public static void main(String[] args) throws Exception {
        String env = "test";
        String jobName = "string";

        // todo 这里的config需要增加配置
        Config config = ConfigFactory.load("attribution-streaming.conf")
                .getConfig(env).getConfig(jobName);

        Config streamConfig = config.getConfig("stream-config");

        String topic = streamConfig.getString(ConfigVars.TOPIC);

//        int checkpointInterval = config.getInt("checkpoint-interval");
//        String checkpointDir = config.getString("checkpoint-dir");

        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();
        see.getConfig().registerTypeWithKryoSerializer(String.class, ProtobufSerializer.class);

        Properties producerProperties = new Properties();
        producerProperties.setProperty("bootstrap.servers", "localhost:9810");
        FlinkKafkaProducer<String> producer = new FlinkKafkaProducer<>("test-producer", new SimpleStringSchema(), producerProperties);

        DataStreamSource<String> messageDataStreamSource = see
                .addSource(new FlinkKafkaConsumer<String>(topic, new SimpleStringSchema(), producerProperties).setStartFromEarliest());

        // todo 中间的处理逻辑
        messageDataStreamSource.print();

        see.execute("test kafka");
    }
}
