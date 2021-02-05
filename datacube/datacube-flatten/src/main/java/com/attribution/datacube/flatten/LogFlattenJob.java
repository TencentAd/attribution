package com.attribution.datacube.flatten;

import com.attribution.datacube.common.factories.datasource.StreamClientFactory;
import com.attribution.datacube.common.factories.record.FlattenedRecordClassFactory;
import com.attribution.datacube.common.flatten.parser.FlattenParserFactory;
import com.attribution.datacube.common.flatten.record.FlattenedConversionLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.attribution.datacube.flatten.flatMapper.LogFlatMapper;
import com.attribution.datacube.flatten.policy.CustomRollingPolicy;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.conv.Conv;
import com.twitter.chill.protobuf.ProtobufSerializer;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import org.apache.flink.core.fs.Path;
import org.apache.flink.formats.parquet.avro.ParquetAvroWriters;
import org.apache.flink.runtime.state.filesystem.FsStateBackend;
import org.apache.flink.streaming.api.CheckpointingMode;
import org.apache.flink.streaming.api.datastream.DataStreamSource;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.CheckpointConfig;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.api.functions.sink.filesystem.StreamingFileSink;
import org.apache.flink.streaming.api.functions.sink.filesystem.bucketassigners.DateTimeBucketAssigner;
import org.apache.flink.streaming.api.functions.sink.filesystem.rollingpolicies.CheckpointRollingPolicy;
import org.apache.flink.streaming.api.functions.sink.filesystem.rollingpolicies.DefaultRollingPolicy;
import org.apache.flink.streaming.api.functions.sink.filesystem.rollingpolicies.OnCheckpointRollingPolicy;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class LogFlattenJob {
    private static final Logger LOG = LoggerFactory.getLogger(LogFlattenJob.class);

    public static void main(String[] args) throws Exception {
        String env = args[0];
        String jobName = args[1];

        // todo 这里的config需要增加配置
        Config config = ConfigFactory.load("conv-flattened-streaming.conf")
                .getConfig(env).getConfig(jobName);

        String logType = config.getString("log-type");
        String savingPath = config.getString("saving-dir");

        Config streamConfig = config.getConfig("stream-config");

        // todo 这里后面需要加入 checkpoint
        int checkpointInterval = config.getInt("checkpoint-interval");
        String checkpointDir = config.getString("checkpoint-dir");

        StreamExecutionEnvironment see = StreamExecutionEnvironment.getExecutionEnvironment();

        // 设置checkpoint
        see.enableCheckpointing(checkpointInterval);
        CheckpointConfig checkpointConfig = see.getCheckpointConfig();
        checkpointConfig.setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
        checkpointConfig.enableExternalizedCheckpoints(CheckpointConfig.ExternalizedCheckpointCleanup.RETAIN_ON_CANCELLATION);
        checkpointConfig.setMinPauseBetweenCheckpoints(10000);
        checkpointConfig.setCheckpointTimeout(60000);
        see.setStateBackend(new FsStateBackend(checkpointDir, true));

        see.getConfig().registerTypeWithKryoSerializer(Conv.ConversionLog.class, ProtobufSerializer.class);

        DataStreamSource<Message> messageDataStreamSource = see
                .addSource(StreamClientFactory.getStreamClient(streamConfig).getSourceFunction());

        LOG.info("set datasource done");

        // todo 中间的处理逻辑
        SingleOutputStreamOperator<FlattenedRecord> flattenedResultStream = messageDataStreamSource
                .flatMap(new LogFlatMapper(FlattenParserFactory.getFlattenedParser(logType)))
                .name("flatten map");
        LOG.info("flat map done");

        StreamingFileSink sink = StreamingFileSink
                .forBulkFormat(new Path(savingPath), ParquetAvroWriters.forReflectRecord(FlattenedRecordClassFactory.getLogClass(logType)))
                // 这里是设置多长时间函缓存一个文件
                .withBucketAssigner(new DateTimeBucketAssigner<>("yyyyMMdd"))
                .build();

        flattenedResultStream.addSink(sink).name("storage sink");

        see.execute("flatten log test");
    }
}
