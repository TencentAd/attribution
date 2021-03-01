package com.attribution.datacube.flatten.tool;

import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.flink.api.common.serialization.DeserializationSchema;
import org.apache.flink.api.common.serialization.SerializationSchema;
import org.apache.flink.api.common.typeinfo.TypeInformation;
import org.apache.flink.api.java.typeutils.TypeExtractor;

import java.io.*;

public class FlattenedMessageSchema implements DeserializationSchema<FlattenedRecord>, SerializationSchema<FlattenedRecord> {
    @Override
    public FlattenedRecord deserialize(byte[] bytes) throws IOException {
        FlattenedRecord  record = null;
        try {
            ObjectInputStream objectInputStream = new ObjectInputStream(new ByteArrayInputStream(bytes));
            record = (FlattenedRecord) objectInputStream.readObject();
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        return record;
    }

    @Override
    public boolean isEndOfStream(FlattenedRecord o) {
        return false;
    }

    @Override
    public byte[] serialize(FlattenedRecord o) {
        ObjectMapper objectMapper = new ObjectMapper();
        try {
            return objectMapper.writeValueAsBytes(o);
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
        return null;
    }

    @Override
    public TypeInformation<FlattenedRecord> getProducedType() {
        return TypeExtractor.getForClass(FlattenedRecord.class);
    }
}
