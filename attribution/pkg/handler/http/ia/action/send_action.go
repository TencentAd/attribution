package action

import (
	"context"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
	"github.com/TencentAd/attribution/attribution/pkg/oauth"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

var (
	DefaultMarketingApiTimeout = 60 * time.Second
	DefaultMarketingApiIsDebug = false
)

// 将归因后的数据发送给ams
type SendAction struct {
	client *ads.SDKClient
	token  *oauth.Token
}

func NewSendAction() *SendAction {
	client := ads.Init(&config.SDKConfig{
		IsDebug: DefaultMarketingApiIsDebug,
	})
	client.UseProduction()
	return &SendAction{client: client}
}

func (action *SendAction) WithToken(token *oauth.Token) *SendAction {
	action.token = token
	return action
}

func (action *SendAction) name() string {
	return "SendAction"
}

func (action *SendAction) Run(i interface{}) {
	ExecuteNamedAction(action, i)
}

func (action *SendAction) run(c *data.ImpAttributionContext) error {
	intersectIdx := make(map[int]struct{})
	for _, fd := range c.FinalDecryptData {
		for _, i := range c.OriginalIndex[fd.Imei] {
			intersectIdx[i] = struct{}{}
		}

		for _, i := range c.OriginalIndex[fd.Idfa] {
			intersectIdx[i] = struct{}{}
		}

		for _, i := range c.OriginalIndex[fd.AndroidId] {
			intersectIdx[i] = struct{}{}
		}

		for _, i := range c.OriginalIndex[fd.Oaid] {
			intersectIdx[i] = struct{}{}
		}
	}

	originalReq := c.ConvParseResult.UserActionAddRequest
	req := &model.UserActionsAddRequest{
		AccountId:       originalReq.AccountId,
		UserActionSetId: originalReq.UserActionSetId,
		Actions:         &[]model.UserAction{},
	}

	for idx := range intersectIdx {
		*req.Actions = append(*req.Actions, (*originalReq.Actions)[idx])
	}

	action.client.SetAccessToken(action.token.Get())
	ctx, _ := context.WithTimeout(context.Background(), DefaultMarketingApiTimeout)
	_, _, err := action.client.UserActions().Add(ctx, *req)
	return err
}
