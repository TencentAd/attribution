package ams

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/model"

	"github.com/TencentAd/attribution/attribution/pkg/common/httpx"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
	"github.com/TencentAd/attribution/attribution/proto/conv"
	"github.com/TencentAd/attribution/attribution/proto/user"
)

// https://developers.e.qq.com/docs/api/user_data/user_action/user_actions_add?version=1.1
type UserActionAddRequestParser struct {
}

func NewUserActionAddRequestParser() *UserActionAddRequestParser {
	return &UserActionAddRequestParser{}
}

func (p *UserActionAddRequestParser) Parse(data interface{}) (*parse.ConvParseResult, error) {
	httpReq := data.(*http.Request)

	body, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		return nil, err
	}
	var req model.UserActionsAddRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}

	// 曝光归因，需要额外的解析campaign_id
	query := httpReq.URL.Query()
	campaignId, err := httpx.HttpMustQueryInt64Param(query, "campaign_id")
	if err != nil {
		return nil, err
	}

	convLogs := make([]*conv.ConversionLog, 0, len(*req.Actions))
	for i, action := range *req.Actions {
		convLog := &conv.ConversionLog{
			UserData: &user.UserData{
				Imei:      action.UserId.HashImei,
				Idfa:      action.UserId.HashIdfa,
				AndroidId: action.UserId.HashAndroidId,
				Oaid:      action.UserId.Oaid,
			},
			EventTime: action.ActionTime,
			Index:     int32(i),
		}
		convLogs = append(convLogs, convLog)
	}

	return &parse.ConvParseResult{
		CampaignId:           campaignId,
		UserActionAddRequest: &req,
		ConvLogs:             convLogs,
	}, nil
}
