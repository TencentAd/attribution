package parser

import (
	"flag"
	"fmt"
	"strings"

	"attribution/pkg/parser/ams"
	"attribution/pkg/parser/jsonline"
	"attribution/proto/conv"
)

var (
	convParserName  = flag.String("conv_parser_name", "ams", "")
)

type ConvParserInterface interface {
	Parse(data interface{}) ([]*conv.ConversionLog, error)
}

func CreateConvParser() (ConvParserInterface, error) {
	switch strings.ToLower(*convParserName) {
	case "ams":
		return ams.NewConvParser(), nil
	case "jsonline":
		return jsonline.NewConvParser(), nil
	default:
		return nil, fmt.Errorf("click parser [%s] not support", *convParserName)
	}
}
