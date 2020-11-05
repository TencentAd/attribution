package handler

import (
    "log"
    "net/http"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv"
)

var (
    idList = []string{"hash_imei", "hash_idfa", "hash_phone", "oaid", "hash_oaid", "hash_android_id"}
)

type set struct {
    kv kv.KV
}

func (s *set) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    values := req.URL.Query()
    for _, v := range idList {
        key := values.Get(v)
        if err := s.kv.Set(key); err != nil {
            log.Print(err)
            _, _ = w.Write([]byte(err.Error()))
            return
        }
    }

}

func NewSetHandler(kv kv.KV) *set {
    return &set{kv: kv}
}

type has struct {
    kv kv.KV
}

func (h *has) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    values := req.URL.Query()
    key := values.Get("id")
    has, err := h.kv.Has(key)
    if err != nil {
        log.Print(err)
    }

    if has {
        _, _ = w.Write([]byte("true"))
    } else {
        _, _ = w.Write([]byte("false"))
    }
}

func NewHasHandler(kv kv.KV) *has {
    return &has{kv: kv}
}
