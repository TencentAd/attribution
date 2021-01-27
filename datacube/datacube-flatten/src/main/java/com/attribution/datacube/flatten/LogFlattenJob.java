package com.attribution.datacube.flatten;

import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
import com.attribution.datacube.common.flatten.parser.FlattenParserFactory;
import com.attribution.datacube.common.flatten.parser.FlattenedConversionLogParser;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.attribution.datacube.flatten.flatMapper.LogFlatMapper;
import com.attribution.datacube.flatten.tool.FlattenedMessageSchema;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.conv.Conv;
import com.twitter.chill.protobuf.ProtobufSerializer;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer;

import java.util.Properties;

public class LogFlattenJob {
    public static void main(String[] args) throws Exception {
        // todo 这里的逻辑就是消费数据，然后使用一个flatMapper来进行打平
        String env = "test";
        String jobName = "flattened";

        // todo 这里的config需要增加配置
        Config config = ConfigFactory.load("conv-flattened-streaming.conf")
                .getConfig(env).getConfig(jobName);

        String logType = config.getString("log-type");

        Config streamConfig = config.getConfig("stream-config");

        // todo 这里后面需要加入 checkpoint
//        int checkpointInterval = config.getInt("checkpoint-interval");
//        String checkpointDir = config.getString("checkpoint-dir");

        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();
        see.getConfig().registerTypeWithKryoSerializer(Conv.ConversionLog.class, ProtobufSerializer.class);

        Properties producerProperties = new Properties();
        producerProperties.setProperty("bootstrap.servers", "localhost:9810");
        FlinkKafkaProducer<FlattenedRecord> producer = new FlinkKafkaProducer<>("flattened_log_test", new FlattenedMessageSchema(), producerProperties);

        DataStreamSource<Message> messageDataStreamSource = see
                .addSource(StreamClientFactory.getStreamClient(streamConfig).getSourceFunction());

        // todo 中间的处理逻辑
        messageDataStreamSource.flatMap(new LogFlatMapper(FlattenParserFactory.getFlattenedParser(logType))).addSink(producer);

        see.execute("test kafka");
    }
}
