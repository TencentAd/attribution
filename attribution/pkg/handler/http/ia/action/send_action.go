package action

import (
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/model"

	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
)

// 将归因后的数据发送给ams
type SendAction struct {
	client *http.Client
}

func NewSendAction() *SendAction {
	return &SendAction{
		client: &http.Client{},
	}
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

	// TODO
	// 调用sdk上报

	return nil
}
