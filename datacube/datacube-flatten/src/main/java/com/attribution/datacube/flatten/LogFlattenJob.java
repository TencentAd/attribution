package com.attribution.datacube.flatten;

import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
import com.attribution.datacube.common.flatten.parser.FlattenParserFactory;
import com.attribution.datacube.common.flatten.record.FlattenedConversionLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.attribution.datacube.flatten.flatMapper.LogFlatMapper;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.conv.Conv;
import com.twitter.chill.protobuf.ProtobufSerializer;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.core.fs.Path;
import org.apache.flink.formats.parquet.avro.ParquetAvroWriters;
import org.apache.flink.streaming.api.CheckpointingMode;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.CheckpointConfig;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.api.functions.sink.filesystem.StreamingFileSink;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class LogFlattenJob {
    private static final Logger LOG = LoggerFactory.getLogger(LogFlattenJob.class);

    public static void main(String[] args) throws Exception {
        // todo 这里的逻辑就是消费数据，然后使用一个flatMapper来进行打平
        String env = "test";
        String jobName = "flattened";

        // todo 这里的config需要增加配置
        Config config = ConfigFactory.load("conv-flattened-streaming.conf")
                .getConfig(env).getConfig(jobName);

        String logType = config.getString("log-type");
        String path = config.getString("saving-dir");

        Config streamConfig = config.getConfig("stream-config");

        // todo 这里后面需要加入 checkpoint
        int checkpointInterval = config.getInt("checkpoint-interval");
        String checkpointDir = config.getString("checkpoint-dir");

        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();
        see.getConfig().registerTypeWithKryoSerializer(Conv.ConversionLog.class, ProtobufSerializer.class);

        // 设置checkpoint
//        see.enableCheckpointing(checkpointInterval);
//        CheckpointConfig checkpointConfig = see.getCheckpointConfig();
//        checkpointConfig.enableExternalizedCheckpoints(CheckpointConfig.ExternalizedCheckpointCleanup.RETAIN_ON_CANCELLATION);
//        checkpointConfig.setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
//        checkpointConfig.setMinPauseBetweenCheckpoints(500);
//        checkpointConfig.setCheckpointTimeout(60000);
//        checkpointConfig.setMaxConcurrentCheckpoints(1);

        DataStreamSource<Message> messageDataStreamSource = see
                .addSource(StreamClientFactory.getStreamClient(streamConfig).getSourceFunction());

        LOG.info("set datasource done");

        // todo 中间的处理逻辑
        SingleOutputStreamOperator<FlattenedRecord> flattenedResultStream = messageDataStreamSource
                .flatMap(new LogFlatMapper(FlattenParserFactory.getFlattenedParser(logType)));

        LOG.info("flat map done");

        StreamingFileSink<FlattenedRecord> sink = StreamingFileSink
                .forBulkFormat(new Path(path), ParquetAvroWriters.forReflectRecord(FlattenedRecord.class))
                .build();
        flattenedResultStream.addSink(sink);

        see.execute("test kafka");
    }
}
