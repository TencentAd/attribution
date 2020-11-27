package token

import (
    "fmt"
    "log"
    "net/http"

    "github.com/TencentAd/attribution/attribution/pkg/oauth"
)

const (
    token = "token"
)

type handle struct {
    token *oauth.Token
}

func (h *handle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    errHandle := func(err error) {
        log.Print(err)
        _, _ = w.Write([]byte(err.Error()))
    }

    values := req.URL.Query()
    accessToken := values.Get(token)
    if accessToken == "" {
        errHandle(fmt.Errorf("can't find token value"))
        return
    }

    if err := h.token.Set(accessToken); err != nil {
        errHandle(err)
        return
    }

    current := h.token.Get()
    _, _ = w.Write([]byte(fmt.Sprintf("current access token is %v", current)))
}

func NewHandler(token *oauth.Token) *handle {
    return &handle{token: token}
}
