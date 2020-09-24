package parser

import (
	"attribution/proto/click"
)

type ClickParserInterface interface {
	Parse(line string) (*click.ClickLog, error)
}

