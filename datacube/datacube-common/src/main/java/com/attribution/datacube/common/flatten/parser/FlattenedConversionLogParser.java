package com.attribution.datacube.common.flatten.parser;

import com.attribution.datacube.common.flatten.record.FlattenedConversionLog;
import com.attribution.datacube.common.flatten.record.FlattenedRecord;
import com.google.protobuf.Message;
import com.tencent.attribution.proto.click.Click;
import com.tencent.attribution.proto.conv.Conv;
import com.tencent.attribution.proto.user.User;

public class FlattenedConversionLogParser extends FlattenParser{
    @Override
    public FlattenedRecord parse(Message message) {
        Conv.ConversionLog conversionLog = (Conv.ConversionLog)message;
        User.UserData userData = conversionLog.getUserData();
        Click.ClickLog clickLog = conversionLog.getMatchClick().getClickLog();
        return FlattenedConversionLog.builder()
                .eventTime(conversionLog.getEventTime())
                .appId(conversionLog.getAppId())
                .convId(conversionLog.getConvId())
                .campaignId(conversionLog.getCampaignId())
                .index(conversionLog.getIndex())
                .originalContent(conversionLog.getOriginalContent())
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
                .clickId(clickLog.getClickId())
                .callback(clickLog.getCallback())
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
