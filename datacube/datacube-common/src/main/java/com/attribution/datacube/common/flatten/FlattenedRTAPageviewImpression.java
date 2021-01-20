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
public class FlattenedRTAPageviewImpression extends FlattenedRecord{
    @Nullable
    @JsonProperty
    String traceId;

    @Nullable
    @JsonProperty
    Long impTime;

    @Nullable
    @JsonProperty
    Integer impCount;

    // todo 这里是可能还有别的字段，但是目前没有

    @Override
    public long getEventTime() {
        return getImpTime();
    }
}
