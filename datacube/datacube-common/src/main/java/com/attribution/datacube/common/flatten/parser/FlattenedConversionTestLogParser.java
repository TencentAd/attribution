package com.attribution.datacube.common.flatten.parser;

import com.attribution.datacube.common.flatten.record.FlattenedConversionTestLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.conv.Conv;

public class FlattenedConversionTestLogParser extends FlattenParser{
    @Override
    public FlattenedRecord parse(Message message) {
        Conv.ConversionLog conversionLog = (Conv.ConversionLog) message;
        return FlattenedConversionTestLog.builder()
                .appId(conversionLog.getAppId())
                .convId(conversionLog.getConvId())
                .build();
    }
}
