package com.attribution.datacube.common.flatten.record;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonPOJOBuilder;
import lombok.*;
import org.apache.avro.reflect.Nullable;

@AllArgsConstructor
@NoArgsConstructor
@Builder
@JsonPOJOBuilder
@Getter
@Setter
@ToString
public class FlattenedConversionTestLog extends FlattenedRecord {
    @Nullable
    @JsonProperty
    public String appId;

    @Nullable
    @JsonProperty
    public String convId;

    @Override
    public long getEventTime() {
        return 0;
    }
}
