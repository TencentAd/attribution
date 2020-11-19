package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"github.com/TencentAd/attribution/attribution/proto/impression"
)

const (
	campaignID     = "campaign_id"
	impressionTime = "impression_time"

	noCampaignIDError     = "can't find campaign_id"
	noImpressionTimeError = "can't find impression_time"
)

type set struct {
	kv kv.KV
}

func (s *set) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	errHandle := func(err error) {
		log.Print(err)
		_, _ = w.Write([]byte(err.Error()))
	}

	values := req.URL.Query()

	cid := values.Get(campaignID)
	if cid == "" {
		errHandle(fmt.Errorf(noCampaignIDError))
		return
	}

	time := values.Get(impressionTime)
	if time == "" {
		errHandle(fmt.Errorf(noImpressionTimeError))
		return
	}

	for k := range impression.IdType_value {
		id := values.Get(k)
		if id == "" {
			continue
		}

		key := keyGenerate(cid, k, id)
		if err := s.kv.Set(key, time); err != nil {
			errHandle(err)
		}
	}
}

func NewSetHandler(kv kv.KV) *set {
	return &set{kv: kv}
}

type get struct {
	kv kv.KV
}

func (g *get) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	request := &impression.Request{
		CampaignId: values.Get(campaignID),
	}

	for k, v := range impression.IdType_value {
		id := values.Get(k)
		if id == "" {
			continue
		}

		request.Records = append(request.Records, &impression.Record{
			IdType:  impression.IdType(v),
			IdValue: id,
		})
	}

	response := &impression.Response{}
	if err := Query(g.kv, request, response); err == nil {
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
	} else {
		_, _ = w.Write([]byte(err.Error()))
	}
}

func NewGetHandler(kv kv.KV) *get {
	return &get{kv: kv}
}
