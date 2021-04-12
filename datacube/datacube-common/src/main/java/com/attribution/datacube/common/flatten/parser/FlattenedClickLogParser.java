package com.attribution.datacube.common.flatten.parser;

import com.attribution.datacube.common.flatten.record.FlattenedClickLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.click.Click;
import com.tencent.attribution.proto.user.User;

/**
 * 将click数据打平的parser，parse方法将Message类的数据打平成FlattenedRecord类的数据
 */
public class FlattenedClickLogParser extends FlattenParser{

    @Override
    public FlattenedRecord parse(Message message) {
        Click.ClickLog clickLog = (Click.ClickLog) message;
        User.UserData userData = clickLog.getUserData();
        return FlattenedClickLog.builder()
                .imei(userData.getImei())
                .idfa(userData.getIdfa())
                .qqOpenId(userData.getQqOpenid())
                .wuId(userData.getWuid())
                .mac(userData.getMac())
                .androidId(userData.getAndroidId())
                .qadid(userData.getQadid())
                .mappedAndroidId(userData.getMappedAndroidIdList())
                .mappedMac(userData.getMappedMacList())
                .taid(userData.getTaid())
                .oaid(userData.getOaid())
                .hashOaid(userData.getHashOaid())
                .uuid(userData.getUuid())
                .ip(userData.getIp())
                .ipv6(userData.getIpv6())
                .clickTime(clickLog.getClickTime())
                .appId(clickLog.getAppId())
                .clickId(clickLog.getClickId())
                .callback(clickLog.getCallback())
                .campaignId(clickLog.getCampaignId())
                .adgroupId(clickLog.getAdgroupId())
                .adId(clickLog.getAdId())
                .adPlatformType(clickLog.getAdPlatformType())
                .accountId(clickLog.getAccountId())
                .agencyId(clickLog.getAgencyId())
                .clickSkuId(clickLog.getClickSkuId())
                .deeplinkUrl(clickLog.getDeeplinkUrl())
                .universalLink(clickLog.getUniversalLink())
                .pageUrl(clickLog.getPageUrl())
                .deviceOsType(clickLog.getDeviceOsType())
                .processTime(clickLog.getProcessTime())
                .promotedObjectType(clickLog.getPromotedObjectType())
                .realCost(clickLog.getRealCost())
                .requestId(clickLog.getRequestId())
                .impressionId(clickLog.getImpressionId())
                .siteSet(clickLog.getSiteSet())
                .encryptedPositionId(clickLog.getEncryptedPositionId())
                .adgroupName(clickLog.getAdgroupName())
                .siteSetName(clickLog.getSiteSetName())
                .cid(clickLog.getCid())
                .billingEvent(clickLog.getBillingEventValue())
                .platform(clickLog.getPlatformValue())
                .build();
    }
}
