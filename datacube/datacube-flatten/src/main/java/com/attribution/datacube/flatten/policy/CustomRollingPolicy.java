package com.attribution.datacube.flatten.policy;

import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import org.apache.flink.streaming.api.functions.sink.filesystem.PartFileInfo;
import org.apache.flink.streaming.api.functions.sink.filesystem.rollingpolicies.CheckpointRollingPolicy;

import java.io.IOException;

public class CustomRollingPolicy<IN, BUCKETID> extends CheckpointRollingPolicy<FlattenedRecord, String> {

    @Override
    public boolean shouldRollOnEvent(PartFileInfo<String> partFileInfo, FlattenedRecord flattenedRecord) throws IOException {
        return true;
    }

    @Override
    public boolean shouldRollOnProcessingTime(PartFileInfo<String> partFileInfo, long l) throws IOException {
        return true;
    }
}
