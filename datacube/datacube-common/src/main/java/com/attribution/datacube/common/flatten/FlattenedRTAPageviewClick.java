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
public class FlattenedRTAPageviewClick extends FlattenedRecord {
    // TODO 这里没有声明一个id是否是因为存储的时候能够有自增的主键

    @Nullable
    @JsonProperty
    String traceId;

    @Nullable
    @JsonProperty
    Long clickTime;

    @Nullable
    @JsonProperty
    Integer clickCount;

    @Override
    public long getEventTime() {
        return getClickTime();
    }
}
