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

func (s *IdSet) Crypt(cryptoFunc func(string, string) (string, error), groupId string) (*IdSet, error) {
	imei, err := cryptoFunc(groupId, s.Imei)
	if err != nil {
		return nil, err
	}
	idfa, err := cryptoFunc(groupId, s.Idfa)
	if err != nil {
		return nil, err
	}
	androidId, err := cryptoFunc(groupId, s.AndroidId)
	if err != nil {
		return nil, err
	}
	oaid, err := cryptoFunc(groupId, s.Oaid)
	if err != nil {
		return nil, err
	}

	return &IdSet{
		Imei:      imei,
		Idfa:      idfa,
		AndroidId: androidId,
		Oaid:      oaid,
	}, nil
}

func (s *IdSet) Encrypt(groupId string) (*IdSet, error) {
	return s.Crypt(crypto.Encrypt, groupId)
}

func (s *IdSet) Decrypt(groupId string) (*IdSet, error) {
	return s.Crypt(crypto.Decrypt, groupId)
}

func ResponseData2IdSet(d *protocal.ResponseData) *IdSet {
	return &IdSet{
		Imei:      d.Imei,
		Idfa:      d.Idfa,
		AndroidId: d.AndroidId,
		Oaid:      d.Oaid,
	}
}
