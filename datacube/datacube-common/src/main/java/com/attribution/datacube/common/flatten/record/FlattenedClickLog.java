package com.attribution.datacube.common.flatten.record;

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
public class FlattenedClickLog extends FlattenedRecord{
    @Nullable
    @JsonProperty
    String imei;

    @Nullable
    @JsonProperty
    String idfa;

    @Nullable
    @JsonProperty
    String qqOpenId;

    @Nullable
    @JsonProperty
    String wuId;

    @Nullable
    @JsonProperty
    String mac;

    @Nullable
    @JsonProperty
    String androidId;

    @Nullable
    @JsonProperty
    String qadid;

    @Nullable
    @JsonProperty
    List<String> mappedAndroidId;

    @Nullable
    @JsonProperty
    List<String> mappedMac;

    @Nullable
    @JsonProperty
    String taid;

    @Nullable
    @JsonProperty
    String oaid;

    @Nullable
    @JsonProperty
    String hashOaid;

    @Nullable
    @JsonProperty
    String uuid;

    @Nullable
    @JsonProperty
    String ip;

    @Nullable
    @JsonProperty
    String ipv6;

    @Nullable
    @JsonProperty
    Long clickTime;

    @Nullable
    @JsonProperty
    String appId;

    @Nullable
    @JsonProperty
    String clickId;

    @Nullable
    @JsonProperty
    String callback;

    @Nullable
    @JsonProperty
    Long campaignId;

    @Nullable
    @JsonProperty
    Long adgroupId;

    @Nullable
    @JsonProperty
    Long adId;

    @Nullable
    @JsonProperty
    Integer adPlatformType;

    @Nullable
    @JsonProperty
    Long accountId;

    @Nullable
    @JsonProperty
    Long agencyId;

    @Nullable
    @JsonProperty
    String clickSkuId;

    @Nullable
    @JsonProperty
    String deeplinkUrl;

    @Nullable
    @JsonProperty
    String universalLink;

    @Nullable
    @JsonProperty
    String pageUrl;

    @Nullable
    @JsonProperty
    String deviceOsType;

    @Nullable
    @JsonProperty
    Long processTime;

    @Nullable
    @JsonProperty
    Integer promotedObjectType;

    @Nullable
    @JsonProperty
    Long realCost;

    @Nullable
    @JsonProperty
    String requestId;

    @Nullable
    @JsonProperty
    String impressionId;

    @Nullable
    @JsonProperty
    Integer siteSet;

    @Nullable
    @JsonProperty
    Long encryptedPositionId;

    @Nullable
    @JsonProperty
    String adgroupName;

    @Nullable
    @JsonProperty
    String siteSetName;

    @Nullable
    @JsonProperty
    Long cid;

    @Nullable
    @JsonProperty
    Integer billingEvent;

    @Nullable
    @JsonProperty
    Integer platform;

    @Override
    public long getEventTime() {
        return 0;
    }
}
