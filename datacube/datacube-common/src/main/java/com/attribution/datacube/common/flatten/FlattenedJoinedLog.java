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
public class FlattenedJoinedLog extends FlattenedRecord {
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
    Integer innerCostMs;

    @Nullable
    @JsonProperty
    Integer requestCount;

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

    // TODO imp暂时没有特别的数据

    @Nullable
    @JsonProperty
    Long clickTime;

    @Nullable
    @JsonProperty
    Integer clickCount;

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
        return getProcessTime();
    }
}
