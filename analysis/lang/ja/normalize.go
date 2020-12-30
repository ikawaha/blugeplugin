package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"golang.org/x/text/unicode/norm"
)

// UnicodeNormalizeCharFilter represents unicode char filter.
type UnicodeNormalizeCharFilter struct {
	form norm.Form
}

// Filter applies per-char normalization.
func (f UnicodeNormalizeCharFilter) Filter(input []byte) []byte {
	return f.form.Bytes(input)
}

// NewUnicodeNormalizeCharFilter returns a normalize char filter.
func NewUnicodeNormalizeCharFilter(form norm.Form) analysis.CharFilter {
	return UnicodeNormalizeCharFilter{
		form: form,
	}
}
