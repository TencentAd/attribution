package uitl

import (
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/storage"
	"math/big"
)

type HdfsKeyManager struct {
	HdfsStorage *storage.HdfsStorage
}

func (h HdfsKeyManager) GetEncryptKey(s string) (*big.Int, error) {
	return nil, nil
}

func NewHdfsKeyManager(hdfsConf conf.HdfsConf) *HdfsKeyManager {
	hdfsStorage, _ := storage.NewHdfsStorage(hdfsConf)
	return &HdfsKeyManager{HdfsStorage: hdfsStorage}
}


