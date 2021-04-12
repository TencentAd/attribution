package com.attribution.datacube.flatten.flatMapper;

import com.attribution.datacube.common.flatten.parser.FlattenParser;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.google.protobuf.Message;
import org.apache.flink.api.common.functions.RichFlatMapFunction;
import org.apache.flink.util.Collector;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class LogFlatMapper extends RichFlatMapFunction<Message, FlattenedRecord> {
    private static final Logger LOG = LoggerFactory.getLogger(LogFlatMapper.class);
    FlattenParser flattenParser;

    public LogFlatMapper(FlattenParser flattenParser) {
        this.flattenParser = flattenParser;
    }

    @Override
    public void flatMap(Message message, Collector<FlattenedRecord> collector) {
        FlattenedRecord flattenedrecord = flattenParser.parse(message);
        collector.collect(flattenedrecord);
    }
}
