//package com.attribution.datacube.join;
//
//
//import com.google.protobuf.Message;
//import com.twitter.chill.protobuf.ProtobufSerializer;
//import com.typesafe.config.Config;
//import com.typesafe.config.ConfigFactory;
//import com.attribution.datacube.common.factories.datasource.StreamClient;
//import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
//import com.attribution.datacube.join.process.JoinProcessFunction;
//import com.attribution.datacube.proto.pageview.PageviewService;
//import org.apache.flink.api.common.functions.MapFunction;
//import org.apache.flink.api.java.functions.KeySelector;
//import org.apache.flink.api.java.tuple.Tuple2;
//import org.apache.flink.streaming.api.datastream.DataStream;
//import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
//import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
//
//public class JoinJob {
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
//        StreamClient requestStreamClient = StreamClientFactory.getStreamClient(config.getConfig("request").getConfig("stream-config"));
//        StreamClient strategyStreamClient = StreamClientFactory.getStreamClient(config.getConfig("strategy").getConfig("stream-config"));
//
//        DataStream<Tuple2<String, PageviewService.Pageview>> requestDataStream = see
//                .addSource(requestStreamClient.getSourceFunction())
//                .name("join_request")
//                .uid("join_request")
//                .map(new MapFunction<Message, Tuple2<String, PageviewService.Pageview>>() {
//                    @Override
//                    public Tuple2<String, PageviewService.Pageview> map(Message message) throws Exception {
//                        return Tuple2.of("request", (PageviewService.Pageview) message);
//                    }
//                });
//
//        DataStream<Tuple2<String, PageviewService.Pageview>> strategyDataStream = see
//                .addSource(strategyStreamClient.getSourceFunction())
//                .name("join_strategy")
//                .uid("join_strategy")
//                .map(new MapFunction<Message, Tuple2<String, PageviewService.Pageview>>() {
//                    @Override
//                    public Tuple2<String, PageviewService.Pageview> map(Message message) throws Exception {
//                        return Tuple2.of("strategy", (PageviewService.Pageview) message);
//                    }
//                });
//
//        // 这里是将两个数据流结合起来进行join操作
//        SingleOutputStreamOperator<Tuple2<String, PageviewService.Pageview>> enrichedRequestDataStream =
//                (requestDataStream.union(strategyDataStream))
//                .keyBy(new KeySelector<Tuple2<String, PageviewService.Pageview>, String>() {
//                    @Override
//                    public String getKey(Tuple2<String, PageviewService.Pageview> value) throws Exception {
//                        return value.f1.getTraceId();
//                    }
//                })
//                .process(new JoinProcessFunction()).name("process_function").uid("process_function");
//
//        see.execute("test kafka");
//    }
//}
