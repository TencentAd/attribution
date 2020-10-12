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

	"attribution/pkg/common/httpx"
	"attribution/proto/click"
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

	clickLog := &click.ClickLog{}
	var err error

	clickLog.ClickTime, err = httpx.HttpMustQueryInt64Param(req, "click_time")
	if err != nil {
		return nil, err
	}

	clickLog.ClickId, err = httpx.HttpMustQueryStringParam(req, "click_id")
	if err != nil {
		return nil, err
	}

	clickLog.Callback, err = httpx.HttpQueryStringParam(req, "callback", "")
	if err != nil {
		return nil, err
	}

	clickLog.CampaignId, err = httpx.HttpQueryInt64Param(req, "campaign_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdgroupId, err = httpx.HttpQueryInt64Param(req, "adgroup_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdId, err = httpx.HttpQueryInt64Param(req, "ad_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdPlatformType, err = httpx.HttpQueryInt32Param(req, "ad_platform_type", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AccountId, err = httpx.HttpQueryInt64Param(req, "account_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AgencyId, err = httpx.HttpQueryInt64Param(req, "agency_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.ClickSkuId, err = httpx.HttpQueryStringParam(req, "click_sku_id", "")
	if err != nil {
		return nil, err
	}

	var billingEvent int32
	billingEvent, err = httpx.HttpQueryInt32Param(req, "billing_event", 0)
	if err != nil {
		return nil, err
	}
	clickLog.BillingEvent = click.BillingEvent(billingEvent)

	clickLog.DeeplinkUrl, err = httpx.HttpQueryStringParam(req, "deeplink_url", "")
	if err != nil {
		return nil, err
	}

	clickLog.UniversalLink, err = httpx.HttpQueryStringParam(req, "universal_url", "")
	if err != nil {
		return nil, err
	}

	clickLog.PageUrl, err = httpx.HttpQueryStringParam(req, "page_url", "")
	if err != nil {
		return nil, err
	}

	clickLog.DeviceOsType, err = httpx.HttpMustQueryStringParam(req, "device_os_type")
	if err != nil {
		return nil, err
	}

	clickLog.ProcessTime, err = httpx.HttpQueryInt64Param(req, "process_time", 0)
	if err != nil {
		return nil, err
	}

	clickLog.PromotedObjectType, err = httpx.HttpQueryInt32Param(req, "promoted_object_type", 0)
	if err != nil {
		return nil, err
	}

	clickLog.RealCost, err = httpx.HttpQueryInt64Param(req, "real_cost", 0)
	if err != nil {
		return nil, err
	}

	clickLog.RequestId, err = httpx.HttpQueryStringParam(req, "request_id", "")
	if err != nil {
		return nil, err
	}

	clickLog.ImpressionId, err = httpx.HttpQueryStringParam(req, "impression_id", "")
	if err != nil {
		return nil, err
	}

	clickLog.SiteSet, err = httpx.HttpQueryInt32Param(req, "site_set", 0)
	if err != nil {
		return nil, err
	}

	clickLog.EncryptedPositionId, err = httpx.HttpQueryInt64Param(req, "encrypted_position_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdgroupName, err = httpx.HttpQueryStringParam(req, "adgroup_name", "")
	if err != nil {
		return nil, err
	}

	clickLog.SiteSetName, err = httpx.HttpQueryStringParam(req, "site_set_name", "")
	if err != nil {
		return nil, err
	}

	return clickLog, nil
}
