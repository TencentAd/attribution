
package click_parser

import (
	"encoding/json"
	"attribution/proto/click"
)

// 按照json格式解析http的body
type ClickParserStruct struct{}

func NewClickParser() *ClickParserStruct {
	Parser := &ClickParserStruct{
	}
	return Parser
}

func (p *ClickParserStruct) Parse(line string) (*click.ClickLog, error) {
	clickLog := new(click.ClickLog)
	err := json.Unmarshal([]byte(line), clickLog)
	if err != nil {
		return nil, err
	}
	return clickLog, nil
}

