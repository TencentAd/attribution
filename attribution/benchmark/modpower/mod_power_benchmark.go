/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/3/20, 9:57 AM
 */

package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/modpower"
	"github.com/TencentAd/attribution/attribution/pkg/handler/file/line"
	"github.com/golang/glog"
)

var (
	prime *big.Int
	encKey *big.Int

	primeHex = flag.String("prime", "CC5A0467495EC2506D699C9FFCE6ED5885AE09384671EE00B93A28712DA814240F2A471B2C77B120FE70DA02D33B0F85CD737B800942D5A2BF80DCCD290FB6553C9834197DC498F6AC69B5ECEF3FD8B05F231E3632E9AECA2B3F50977CF9033AF3005A9C0A339CFB4922971B3AF05A5955983C12B153BB78A2B1FB14C84A3C662ADDE5BCEBE8779FF9F97C6E73BD29D4044242581455EEB0543E2DB35F43997F46F8596A58080DC053BBB71F9A557185DF80738238713D3EDFD77D47B26977B373FB0D969920B3909CCC24792B5B4E94AD29F6AE6BD73ED5FE6528CDDBEA1560BBCD36E8B25008021A26E9A4E51BBCD8436F38D6A222E2138E7042A73A7877D7", "")
	encKeyHex = flag.String("encrypt_key", "9cff5b8e1899bb3e7a356f3eb6b45a85", "")
	filePath = flag.String("file_path", "", "")
	workerNumber = flag.Int("worker_number", 4, "")
)

type ModPowerHandle struct {
	Prime  *big.Int
	EncKey *big.Int
	filename string
}

func NewModPowerHandle(filename string) *ModPowerHandle {
	return & ModPowerHandle{
		filename: filename,
	}
}

func (h *ModPowerHandle) WithPrime(prime *big.Int) *ModPowerHandle {
	h.Prime = prime
	return h
}

func (h *ModPowerHandle) WithEncKey(encKey *big.Int) *ModPowerHandle {
	h.EncKey = encKey
	return h
}

func (h *ModPowerHandle) Run() error {
	lp := line.NewLineProcess(h.filename, h.processLine, func(line string, err error) {
		glog.Errorf("failed to handle line[%s], err[%v]", line, err)
	}).WithParallelism(*workerNumber).WithQueueTimeout(time.Second)

	if err := lp.Run(); err != nil {
		return err
	}

	lp.WaitDone()
	return nil
}

func (h *ModPowerHandle) processLine(line string) error {
	data, err := modpower.Hex2BigInt(line)
	if err != nil {
		return err
	}

	_ = modpower.Encrypt(data, h.EncKey, h.Prime)
	return nil
}

func main() {
	flag.Parse()

	var err error
	prime, err = modpower.Hex2BigInt(*primeHex)
	if err != nil {
		glog.Fatal("'prime' not valid hex")
	}
	encKey, err = modpower.Hex2BigInt(*encKeyHex)
	if err != nil {
		glog.Fatal("'encrypt_key' not valid hex")
	}

	handle := NewModPowerHandle(*filePath).WithPrime(prime).WithEncKey(encKey)
	startTime := time.Now()
	if err := handle.Run(); err != nil {
		panic(err)
	}
	fmt.Println("time used:", time.Since(startTime).Milliseconds())
}
