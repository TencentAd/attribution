package ocean

import (
	"github.com/TencentAd/attribution/attribution/pkg/common/httpx"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"net/http"
)

type ClickParser struct {
}

func NewJvliangClickParser() *ClickParser {
	return &ClickParser{}
}

func (c ClickParser) Parse(input interface{}) (*click.ClickLog, error) {
	req := input.(*http.Request)

	clickLog := &click.ClickLog{
		UserData: &user.UserData{},
	}
	userData := clickLog.UserData
	var err error
	query := req.URL.Query()

	clickLog.ClickTime, err = httpx.HttpMustQueryInt64Param(query, "ts")
	if err != nil {
		return nil, err
	}

	clickLog.CampaignId, err = httpx.HttpQueryInt64Param(query, "aid", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AccountId, err = httpx.HttpQueryInt64Param(query, "advertiser_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.Cid, err = httpx.HttpQueryInt64Param(query, "cid", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdgroupId, err = httpx.HttpQueryInt64Param(query, "campaign_id", 0)
	if err != nil {
		return nil, err
	}

	clickLog.AdPlatformType, err = httpx.HttpQueryInt32Param(query, "csite", 0)
	if err != nil {
		return nil, err
	}

	clickLog.RequestId = httpx.HttpQueryStringParam(query, "request_id", "")

	clickLog.Callback = httpx.HttpQueryStringParam(query, "callback_url", "")

	clickLog.DeviceOsType, err = httpx.HttpMustQueryStringParam(query, "os")

	if clickLog.DeviceOsType == "0" {
		userData.Imei = httpx.HttpQueryStringParam(query, "imei", "")
	} else if clickLog.DeviceOsType == "1" {
		userData.Idfa = httpx.HttpQueryStringParam(query, "idfa", "")
	}

	userData.AndroidId = httpx.HttpQueryStringParam(query, "androidid", "")

	userData.Oaid = httpx.HttpQueryStringParam(query, "oaid", "")

	userData.HashOaid = httpx.HttpQueryStringParam(query, "oaid_md5", "")

	userData.Ip = httpx.HttpQueryStringParam(query, "ip", "")

	if glog.V(10) {
		glog.V(10).Infof("%s", proto.MarshalTextString(clickLog))
	}

	return clickLog, nil
}
