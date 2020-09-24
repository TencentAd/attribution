package conv_parser

import (
	"attribution/proto/conv"
	"encoding/json"
)

// 按照json格式解析http的body
type ConvParserStruct struct{}

func NewConvParser() *ConvParserStruct {
	Parser := &ConvParserStruct{
	}
	return Parser
}

func (p *ConvParserStruct) Parse(line string) (*conv.ConversionLog, error) {
	convLog := new(conv.ConversionLog)
	err := json.Unmarshal([]byte(line), convLog)
	if err != nil {
		return nil, err
	}
	return convLog, nil
}
