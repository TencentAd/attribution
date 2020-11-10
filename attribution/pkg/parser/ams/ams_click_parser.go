/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/9/20, 7:10 PM
 */

package ams

import (
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/common/httpx"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

// 接收点击转发的数据，协议参考文档
// https://developers.e.qq.com/docs/guide/conversion/click?version=1.1
type ClickParser struct {
}

func NewAMSClickParser() *ClickParser {
	return &ClickParser{}
}

func (p *ClickParser) Parse(input interface{}) (*click.ClickLog, error) {
	req := input.(*http.Request)

	clickLog := &click.ClickLog{
		UserData: &user.UserData{},
	}
	userData := clickLog.UserData
	var err error
	query := req.URL.Query()

	clickLog.ClickTime, err = httpx.HttpMustQueryInt64Param(query, "click_time")
	if err != nil {
		return nil, err
	}

	clickLog.ClickId, err = httpx.HttpMustQueryStringParam(query, "click_id")
	if err != nil {
		return nil, err
	}

	clickLog.Callback = httpx.HttpQueryStringParam(query, "callback", "")

	clickLog.CampaignId, err = httpx.HttpQueryInt64Param(query, "campaign_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdgroupId, err = httpx.HttpQueryInt64Param(query, "adgroup_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdId, err = httpx.HttpQueryInt64Param(query, "ad_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdPlatformType, err = httpx.HttpQueryInt32Param(query, "ad_platform_type", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AccountId, err = httpx.HttpQueryInt64Param(query, "account_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AgencyId, err = httpx.HttpQueryInt64Param(query, "agency_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.ClickSkuId = httpx.HttpQueryStringParam(query, "click_sku_id", "")

	var billingEvent int32
	billingEvent, err = httpx.HttpQueryInt32Param(query, "billing_event", 0)
	if err != nil {
		return nil, err
	}
	clickLog.BillingEvent = click.BillingEvent(billingEvent)

	clickLog.DeeplinkUrl = httpx.HttpQueryStringParam(query, "deeplink_url", "")

	clickLog.UniversalLink = httpx.HttpQueryStringParam(query, "universal_url", "")

	clickLog.PageUrl = httpx.HttpQueryStringParam(query, "page_url", "")

	clickLog.DeviceOsType, err = httpx.HttpMustQueryStringParam(query, "device_os_type")
	if err != nil {
		return nil, err
	}

	clickLog.ProcessTime, err = httpx.HttpQueryInt64Param(query, "process_time", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AppId = httpx.HttpQueryStringParam(query, "promoted_object_id", "")

	clickLog.PromotedObjectType, err = httpx.HttpQueryInt32Param(query, "promoted_object_type", 0)
	if err != nil {
		return nil, err
	}

	clickLog.RealCost, err = httpx.HttpQueryInt64Param(query, "real_cost", 0)
	if err != nil {
		return nil, err
	}

	clickLog.RequestId = httpx.HttpQueryStringParam(query, "request_id", "")

	clickLog.ImpressionId = httpx.HttpQueryStringParam(query, "impression_id", "")

	clickLog.SiteSet, err = httpx.HttpQueryInt32Param(query, "site_set", 0)
	if err != nil {
		return nil, err
	}

	clickLog.EncryptedPositionId, err = httpx.HttpQueryInt64Param(query, "encrypted_position_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdgroupName = httpx.HttpQueryStringParam(query, "adgroup_name", "")

	clickLog.SiteSetName = httpx.HttpQueryStringParam(query, "site_set_name", "")

	muid := httpx.HttpQueryStringParam(query, "muid", "")
	if clickLog.DeviceOsType == "android" {
		userData.Imei = muid
	} else if clickLog.DeviceOsType == "ios" {
		userData.Idfa = muid
	}

	userData.AndroidId = httpx.HttpQueryStringParam(query, "hash_android_id", "")

	userData.Ip = httpx.HttpQueryStringParam(query, "ip", "")

	userData.Oaid = httpx.HttpQueryStringParam(query, "oaid", "")

	userData.Ipv6 = httpx.HttpQueryStringParam(query, "ipv6", "")

	userData.HashOaid = httpx.HttpQueryStringParam(query, "hash_oaid", "")

	if glog.V(10) {
		glog.V(10).Infof("%s", proto.MarshalTextString(clickLog))
	}

	return clickLog, nil
}
