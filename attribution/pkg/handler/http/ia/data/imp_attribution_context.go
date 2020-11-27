package data

import (
	"math"
	"strconv"

	"github.com/TencentAd/attribution/attribution/pkg/data"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
)

type ImpAttributionContext struct {
	*data.BaseContext

	ConvParseResult *parse.ConvParseResult
	CampaignId      int64 // 推广计划id
	CampaignIdStr   string

	EncryptData      []*IdSet // 广告主加密
	EncryptTwiceData []*IdSet // ams二次加密
	FirstDecryptData []*IdSet // 二次加密后，第一次解密
	IntersectData    []*IdSet // 归因交集
	FinalDecryptData []*IdSet // 明文数据

	OriginalIndex map[string][]int //原始用户ID => 下标
	MinActionTime int64
	MaxActionTime int64
}

func NewImpAttributionContext(pr *parse.ConvParseResult) (*ImpAttributionContext, error) {
	campaignId := pr.CampaignId

	c := &ImpAttributionContext{
		BaseContext:     data.NewBaseContext(),
		ConvParseResult: pr,
		CampaignId:      campaignId,
		CampaignIdStr:   strconv.FormatInt(campaignId, 10),
	}

	c.buildOriginalIndex()
	c.InitActionTime()
	return c, nil
}

func (c *ImpAttributionContext) GroupId() string {
	return c.CampaignIdStr
}

func (c *ImpAttributionContext) buildOriginalIndex() {
	c.OriginalIndex = make(map[string][]int)

	for i, convLog := range c.ConvParseResult.ConvLogs {
		userData := convLog.UserData
		if userData.Imei != "" {
			c.OriginalIndex[userData.Imei] = append(c.OriginalIndex[userData.Imei], i)
		}
		if userData.Idfa != "" {
			c.OriginalIndex[userData.Idfa] = append(c.OriginalIndex[userData.Idfa], i)
		}
		if userData.AndroidId != "" {
			c.OriginalIndex[userData.AndroidId] = append(c.OriginalIndex[userData.AndroidId], i)
		}
		if userData.Oaid != "" {
			c.OriginalIndex[userData.Oaid] = append(c.OriginalIndex[userData.Oaid], i)
		}
	}
}

// 取最新的行为时间作为action time，尽量多的归因
func (c *ImpAttributionContext) InitActionTime() {
	var minActionTime int64 = math.MaxInt64
	var maxActionTime int64 = 0
	for _, convLog := range c.ConvParseResult.ConvLogs {
		if convLog.EventTime < minActionTime {
			minActionTime = convLog.EventTime
		}
		if convLog.EventTime > maxActionTime {
			maxActionTime = convLog.EventTime
		}
	}
	c.MinActionTime = minActionTime
	c.MaxActionTime = maxActionTime
}
