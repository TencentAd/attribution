/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/13/20, 10:16 AM
 */

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/loader"
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/golang/glog"
)

func main() {
	flagx.Parse()

	if err := crypto.InitCrypto(); err != nil {
		panic(err)
	}

	loader.StartDoubleBufferLoad(5)

	n := 10
	m := 100
	groupId := "12345"
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < m; j++ {
				data := "fb946424beb90f3a3956737ec3c9b521"
				encData, err := crypto.Encrypt(groupId, data)
				if err != nil {
					glog.Errorf("failed to encrypt")
					return
				}
				decData, err := crypto.Decrypt(groupId, encData)
				if err != nil {
					glog.Errorf("failed to decrypt")
					return
				}
				fmt.Printf("orginal[%s] encrypt[%s] decrypt[%s]\n", data, encData, decData)
			}
		}()
	}

	time.Sleep(time.Second * 10)

	wg.Wait()
}
