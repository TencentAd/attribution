package com.attribution.datacube.common.flatten;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonPOJOBuilder;
import lombok.*;
import org.apache.avro.reflect.Nullable;

import java.util.List;

@AllArgsConstructor
@NoArgsConstructor
@Builder
@JsonPOJOBuilder
@Getter
@Setter
@ToString
public class FlattenedRTAPageviewStrategy extends FlattenedRecord {

    @Nullable
    @JsonProperty
    String traceId;

    @Nullable
    @JsonProperty
    Long processTime;

    /**
     * 这个是作为策略的primary key
     */
    @Nullable
    @JsonProperty
    Long rtaStrategyId;

    @Nullable
    @JsonProperty
    String rtaValidPlatform;

    @Nullable
    @JsonProperty
    List<Long> strategyIds;

    @Nullable
    @JsonProperty
    Integer hitStrategyCount;

    @Override
    public long getEventTime() {
        return getProcessTime();
    }
}
