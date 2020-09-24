package parser

import (
	"attribution/proto/conv"
)

type ConvParserInterface interface {
	Parse(line string) (*conv.ConversionLog, error)
}

