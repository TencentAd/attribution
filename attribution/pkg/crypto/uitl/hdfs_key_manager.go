package uitl

import (
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/storage"
	"github.com/golang/glog"
	"math/big"
)

type HdfsKeyManager struct {
	hdfsStorage *storage.HdfsStorage
}

func NewHdfsKeyManager(hdfsConf *conf.HdfsConf) *HdfsKeyManager {
	hdfsStorage, _ := storage.NewHdfsStorage(hdfsConf)
	return &HdfsKeyManager{hdfsStorage: hdfsStorage}
}

func (h *HdfsKeyManager) GetEncryptKey(groupId string) (*big.Int, error) {
	encKey, err := h.hdfsStorage.Fetch(groupId)
	if err != nil {
		glog.Errorf("fail to get encrypt key from hdfs, err[%v]", err)
		return nil, err
	}
	result := new(big.Int)
	result, _ = result.SetString(encKey, 16)
	return result, nil
}

// 该方法用于用户人工存储秘钥
func (h *HdfsKeyManager) StorageEncryptKey(groupId string, encKey *big.Int) error {
	err := h.hdfsStorage.Storage(groupId, encKey.Text(16))
	return err
}
