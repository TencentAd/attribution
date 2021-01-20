package com.attribution.datacube.common.flatten;

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
public class FlattenedRTAPageviewConversion extends FlattenedRecord {

    @Nullable
    @JsonProperty
    String traceId;

    @Nullable
    @JsonProperty
    Long actionTime;

    @Nullable
    @JsonProperty
    String actionType;

    @Nullable
    @JsonProperty
    Integer conversionCount;

    @Override
    public long getEventTime() {
        return 0;
    }
}
