package storage

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/colinmarc/hdfs"
	"github.com/golang/glog"
	"strings"
)

type HdfsStorage struct {
	client   *hdfs.Client
	fileName string
}

func NewHdfsStorage(hdfsConf conf.HdfsConf) (*HdfsStorage, error) {
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: hdfsConf.NodeAddress,
		User:      hdfsConf.User,
	})
	if err != nil {
		glog.Errorf("fail to connect to hdfs, err[%v]", err)
		return nil, err
	}

	return &HdfsStorage{
		client:   client,
		fileName: hdfsConf.FileName,
	}, nil
}

func (h HdfsStorage) Storage(groupId string, encryptKey string) error {
	writer, err := h.client.Append(h.fileName)
	if err != nil {
		glog.Errorf("fail to open file for writing, err[%v]", err)
		return err
	}
	str := fmt.Sprintf("%s\t%s\n\r", groupId, encryptKey)
	_, err = writer.Write([]byte(str))
	return err
}

func (h HdfsStorage) Fetch(groupId string) (string, error) {
	reader, err := h.client.Open(h.fileName)
	if err != nil {
		glog.Errorf("fail to open file for reading, err[%v]", err)
		return "", err
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "\t")
		if strings.Compare(split[0], groupId) == 0 {
			return split[1], nil
		}
	}
	return "", errors.New("encrypt key not exist")
}
