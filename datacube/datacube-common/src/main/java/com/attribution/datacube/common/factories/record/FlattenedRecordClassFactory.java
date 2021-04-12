package com.attribution.datacube.common.factories.record;

import com.attribution.datacube.common.flatten.record.FlattenedClickLog;
import com.attribution.datacube.common.flatten.record.FlattenedConversionLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;

public class FlattenedRecordClassFactory {
    public static Class<? extends FlattenedRecord> getLogClass(String type) throws Exception {
        switch (type) {
            case "conversion": {
                return FlattenedConversionLog.class;
            }
            case "click": {
                return FlattenedClickLog.class;
            }
            default: {
                throw new Exception("no such type");
            }
        }
    }
}
