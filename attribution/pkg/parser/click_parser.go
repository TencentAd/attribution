package parser

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TencentAd/attribution/attribution/pkg/parser/ams"
	"github.com/TencentAd/attribution/attribution/pkg/parser/jsonline"
	"github.com/TencentAd/attribution/attribution/proto/click"
)

var (
	clickParserName = flag.String("click_parser_name", "ams", "")
)

type ClickParserInterface interface {
	Parse(input interface{}) (*click.ClickLog, error)
}

func CreateClickParser() (ClickParserInterface, error) {
	switch strings.ToLower(*clickParserName) {
	case "ams":
		return ams.NewAMSClickParser(), nil
	case "jsonline":
		return jsonline.NewJsonClickParser(), nil

	default:
		return nil, fmt.Errorf("click parser [%s] not support", *clickParserName)
	}
}