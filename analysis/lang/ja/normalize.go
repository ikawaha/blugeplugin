package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"golang.org/x/text/unicode/norm"
)

type UnicodeNormalizeCharFilter struct {
	form norm.Form
}

func (f UnicodeNormalizeCharFilter) Filter(input []byte) []byte {
	return f.form.Bytes(input)
}

func NewUnicodeNormalizeCharFilter(form norm.Form) analysis.CharFilter {
	return UnicodeNormalizeCharFilter{
		form: form,
	}
}
