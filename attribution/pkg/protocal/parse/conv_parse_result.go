package parse

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"

	"github.com/TencentAd/attribution/attribution/pkg/protocal/ams/conversion"
	"github.com/TencentAd/attribution/attribution/proto/conv"
)

// 转化数据解析结果
type ConvParseResult struct {
	CampaignId           int64
	AppId                string
	ConvId               string
	UserActionAddRequest *model.UserActionsAddRequest // 行为请求
	Actions              []*amsconversion.Action      // 行为数据(自归因)
	ConvLogs             []*conv.ConversionLog
}
