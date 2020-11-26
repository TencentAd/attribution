package handler

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"github.com/TencentAd/attribution/attribution/proto/impression"
)

func handleOne(wg *sync.WaitGroup, ch chan *impression.Record, kv kv.KV, record *impression.Record, key string) {
	defer wg.Done()
	if value, err := kv.Get(key); err == nil {
		time, _ := strconv.ParseUint(value, 10, 64)
		record.ImpressionTime = time
		ch <- record
	}
}

func Query(kv kv.KV, request *impression.Request, response *impression.Response) error {
	if request.CampaignId == "" {
		return fmt.Errorf(noCampaignIDError)
	}

	wg := &sync.WaitGroup{}
	ch := make(chan *impression.Record)
	for _, v := range request.Records {
		wg.Add(1)
		key := keyGenerate(request.CampaignId, impression.IdType_name[int32(v.IdType)], v.IdValue)
		go handleOne(wg, ch, kv, v, key)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		response.Records = append(response.Records, c)
	}
	return nil
}
