package hdfs

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/TencentAd/attribution/attribution/pkg/common/loader"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/structure"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/util"
	"github.com/colinmarc/hdfs"
	"github.com/golang/glog"
)

var (
	keyHdfsPath = flag.String("key_hdfs_path", "", "")
)

type KeyManager struct {
	client   *hdfs.Client
	fileName string

	buffer *loader.FileDoubleBuffer
}

type PairData struct {
	m map[string]*structure.CryptoPair
}

func NewPairData() interface{} {
	return &PairData{}
}

func NewHDFSKeyManager() interface{} {
	l := loader.NewHdfsHotLoaderWithDefaultClient(*keyHdfsPath, loadData, NewPairData)
	if l == nil {
		return nil
	}
	buffer := loader.NewFileDoubleBuffer(l)
	buffer.SetNotify(
		func(err error) {
			if err != nil {
				glog.Errorf("failed to load hdfs keys, err: %v", err)
			}
		})
	return &KeyManager{
		buffer: buffer,
	}
}

func loadData(reader io.Reader, i interface{}) error {
	data := i.(*PairData)
	data.m = make(map[string]*structure.CryptoPair)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, "\t")
		if len(split) != 2 {
			return fmt.Errorf("not valid record: %s", line)
		}

		groupId := split[0]
		encKey, err := util.Hex2BigInt(split[1])
		if err != nil {
			return err
		}
		data.m[groupId] = util.CreateCryptoPair(encKey)
	}

	return scanner.Err()
}

func (manager *KeyManager) GetEncryptKey(groupId string) (*big.Int, error) {
	pair, ok := manager.buffer.Data().(*PairData).m[groupId]
	if !ok {
		return nil, fmt.Errorf("crypto pair not found")
	}

	return pair.EncKey, nil
}

func (manager *KeyManager) GetDecryptKey(groupId string) (*big.Int, error) {
	pair, ok := manager.buffer.Data().(*PairData).m[groupId]
	if !ok {
		return nil, fmt.Errorf("crypto pair not found")
	}

	return pair.DecKey, nil
}
