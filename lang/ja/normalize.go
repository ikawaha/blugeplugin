package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"golang.org/x/text/unicode/norm"
)

type JapaniseNormalizeFilter struct {}

func (f JapaniseNormalizeFilter) Filter(input []byte) []byte{
	return norm.NFKC.Bytes(input)
}

func NormalizeFilter() analysis.CharFilter{
	return JapaniseNormalizeFilter{}
}
