//package com.attribution.datacube.service;
//
//import com.google.protobuf.Message;
//import com.twitter.chill.protobuf.ProtobufSerializer;
//import com.typesafe.config.Config;
//import com.typesafe.config.ConfigFactory;
//import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
//import com.attribution.datacube.proto.pageview.PageviewService;
//import org.apache.flink.streaming.api.datastream.DataStreamSource;
//import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
//
//public class RTAStreamingProcess {
//    public static void main(String[] args) throws Exception {
//        String env = args[0];
//        String jobName = args[1];
//
//        // todo 这里的config需要增加配置
//        Config config = ConfigFactory.load("rta-streaming.conf")
//                .getConfig(env).getConfig(jobName);
//        int checkpointInterval = config.getInt("checkpoint-interval");
//        String checkpointDir = config.getString("checkpoint-dir");
//
//        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();
//        see.getConfig().registerTypeWithKryoSerializer(PageviewService.Pageview.class, ProtobufSerializer.class);
//
//        DataStreamSource<Message> messageDataStreamSource = see
//                .addSource(StreamClientFactory.getStreamClient(config).getSourceFunction());
//
//        // todo 中间的处理逻辑
//        messageDataStreamSource.print();
//
//        see.execute("test kafka");
//    }
//}
