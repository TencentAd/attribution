package data

import (
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/protocal"
)

type IdSet struct {
	Imei      string
	Idfa      string
	AndroidId string
	Oaid      string
}

func (s *IdSet) ToRequestData() *protocal.RequestData {
	return &protocal.RequestData{
		Imei:      s.Imei,
		Idfa:      s.Idfa,
		AndroidId: s.AndroidId,
		Oaid:      s.Oaid,
	}
}

func (s *IdSet) Crypt(p *crypto.Parallel, cryptoFunc func(string, string) (string, error), groupId string, result *IdSet) {
	p.AddTask(cryptoFunc, groupId, s.Imei, &result.Imei)
	p.AddTask(cryptoFunc, groupId, s.Idfa, &result.Idfa)
	p.AddTask(cryptoFunc, groupId, s.AndroidId, &result.AndroidId)
	p.AddTask(cryptoFunc, groupId, s.Oaid, &result.Oaid)
}

func (s *IdSet) Encrypt(p *crypto.Parallel, groupId string, result *IdSet) {
	s.Crypt(p, crypto.Encrypt, groupId, result)
}

func (s *IdSet) Decrypt(p *crypto.Parallel, groupId string, result *IdSet) {
	s.Crypt(p, crypto.Decrypt, groupId, result)
}

func ResponseData2IdSet(d *protocal.ResponseData) *IdSet {
	return &IdSet{
		Imei:      d.Imei,
		Idfa:      d.Idfa,
		AndroidId: d.AndroidId,
		Oaid:      d.Oaid,
	}
}
