package protocal

import "github.com/TencentAd/attribution/attribution/pkg/crypto"

func ProcessData(p *crypto.Parallel, groupId string, reqData *RequestData,
	cryptoFunc func(string, string) (string, error), resp *ResponseData)  {

	p.AddTask(cryptoFunc, groupId, reqData.Imei, &resp.Imei)
	p.AddTask(cryptoFunc, groupId, reqData.Idfa, &resp.Idfa)
	p.AddTask(cryptoFunc, groupId, reqData.AndroidId, &resp.AndroidId)
	p.AddTask(cryptoFunc, groupId, reqData.Oaid, &resp.Oaid)
}
