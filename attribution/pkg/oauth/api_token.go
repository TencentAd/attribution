package oauth

import (
	"context"
	"sync"
	"time"
)

var (
	DefaultIntervalTime = 31 * time.Second
	TokenPrefixKey      = "marketing-api-access-token"
)

type store interface {
	Get(string) (string, error)
	Set(string, string) error
}

type Token struct {
	lock  *sync.RWMutex
	token string
	store store
}

func (t *Token) Get(sync... bool) string {
	if len(sync) > 0 && sync[0] {
		_ = t.fetch()
	}

	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.token
}

func (t *Token) Set(token string) error {
	return t.store.Set(TokenPrefixKey, token)
}

func (t *Token) FetchBackGround(ctx context.Context) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				_ = t.fetch()
				time.Sleep(DefaultIntervalTime)
			}
		}
	}()
}

func (t *Token) fetch() error {
	token, err := t.store.Get(TokenPrefixKey)
	if err != nil {
		return err
	}

	t.lock.Lock()
	t.token = token
	t.lock.Unlock()
	return nil
}

func NewToken(store store) (*Token, error) {
	token := &Token{
		lock: &sync.RWMutex{},
		store: store,
	}
	err := token.fetch()
	return token, err
}
