package metadata

import (
    "fmt"
    "log"
    "net/http"

    "github.com/TencentAd/attribution/attribution/pkg/oauth"
)

const (
    token = "token"
)

type tokenHandle struct {
    token *oauth.Token
}

func (t *tokenHandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    errHandle := func(err error) {
        log.Print(err)
        _, _ = w.Write([]byte(err.Error()))
    }

    values := req.URL.Query()
    accessToken := values.Get(token)
    if accessToken == "" {
        errHandle(fmt.Errorf("can't find metadata value"))
        return
    }

    if err := t.token.Set(accessToken); err != nil {
        errHandle(err)
        return
    }

    current := t.token.Get(true)
    _, _ = w.Write([]byte(fmt.Sprintf("current access metadata is %v", current)))
}

func NewTokenHandler(token *oauth.Token) *tokenHandle {
    return &tokenHandle{token: token}
}
