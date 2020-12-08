package action

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/crypto/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
)

var (
	AmsDecryptUrl = flag.String("ams_decrypt_url", "http://tracking.e.qq.com/crypto/decrypt", "")
)

// 发送给AMS进行第二次解密
type FinalDecryptAction struct {
	client *http.Client
}

func NewFinalDecryptAction() *FinalDecryptAction {
	return &FinalDecryptAction{
		client: &http.Client{},
	}
}

func (action *FinalDecryptAction) name() string {
	return "FinalDecryptAction"
}

func (action *FinalDecryptAction) Run(i interface{}) {
	ExecuteNamedAction(action, i)
}

func (action *FinalDecryptAction) run(c *data.ImpAttributionContext) error {
	req := action.buildCryptoRequest(c)
	return action.doHttpRequest(c, req)
}


func (action *FinalDecryptAction) buildCryptoRequest(c *data.ImpAttributionContext) *protocal.CryptoRequest {
	req := &protocal.CryptoRequest{
		CampaignId: c.CampaignId,
	}
	for _, set := range c.IntersectData {
		req.Data = append(req.Data, set.ToRequestData())
	}

	return req
}

func (action *FinalDecryptAction) doHttpRequest(c *data.ImpAttributionContext, req *protocal.CryptoRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", *AmsDecryptUrl, bytes.NewReader(body))
	if err != nil {
		return err
	}

	httpResp, err := action.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	var resp protocal.CryptoResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("encrypt twice err: %s", resp.Message)
	}

	for _, d := range resp.Data {
		c.FinalDecryptData = append(c.FinalDecryptData, data.ResponseData2IdSet(d))
	}

	return nil
}