package com.attribution.datacube.common.flatten.parser;

import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.google.protobuf.Message;

import java.io.Serializable;

public abstract class FlattenParser implements Serializable {
    private static final long serialVersionUID = 1111013L;

    public abstract FlattenedRecord parse(Message message);
}
