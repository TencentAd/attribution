package jsonline

import (
	"encoding/json"

	"github.com/TencentAd/attribution/attribution/proto/click"
)

// 按照json格式解析http的body
type JsonClickParser struct{}

func NewJsonClickParser() *JsonClickParser {
	return &JsonClickParser{}
}

func (p *JsonClickParser) Parse(input interface{}) (*click.ClickLog, error) {
	line := input.(string)
	clickLog := new(click.ClickLog)
	err := json.Unmarshal([]byte(line), clickLog)
	if err != nil {
		return nil, err
	}
	return clickLog, nil
}
