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
public class FlattenedRTAPageviewRequest extends FlattenedRecord {

    @Nullable
    @JsonProperty
    String traceId;

    @Nullable
    @JsonProperty
    Long processTime;

    @Nullable
    @JsonProperty
    String fromPlatform;

    @Nullable
    @JsonProperty
    String imei;

    @Nullable
    @JsonProperty
    String idfa;

    @Nullable
    @JsonProperty
    String oaid;

    @Nullable
    @JsonProperty
    Long innerCostMs;

    @Nullable
    @JsonProperty
    Integer requestCount;

    @Override
    public long getEventTime() {
        return getProcessTime();
    }
}
