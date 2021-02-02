package com.attribution.datacube.common.flatten.parser;

import com.typesafe.config.Config;

public class FlattenParserFactory {
    public static FlattenParser getFlattenedParser(String type) throws Exception {
        switch (type) {
            case "click": {
                return new FlattenedClickLogParser();
            }
            case "conversion" : {
                return new FlattenedConversionLogParser();
            }
            default: {
                throw new Exception("no such parser type");
            }
        }
    }
}
