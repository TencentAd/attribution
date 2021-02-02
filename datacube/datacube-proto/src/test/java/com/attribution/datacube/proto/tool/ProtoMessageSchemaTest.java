package com.attribution.datacube.proto.tool;


import com.google.protobuf.InvalidProtocolBufferException;
import com.tencent.attribution.proto.conv.Conv;
import org.junit.Test;

public class ProtoMessageSchemaTest {
    @Test
    public void testParse() throws InvalidProtocolBufferException {
        byte[] bytes = new byte[]{26, 10, 116, 101, 115, 116, 32, 97, 112, 112, 105, 100, 34, 11, 116, 101, 115, 116, 32, 99, 111, 110, 118, 105, 100, -62, 62, 2, 16, 1, -54, 62, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100};
        Conv.ConversionLog conversionLog = Conv.ConversionLog.parseFrom(bytes);
        System.out.println(conversionLog.getAppId());
        System.out.println(conversionLog.getConvId());
    }
}