package ja

import (
	"github.com/blugelabs/bluge/analysis/token"
)

// NewStopWordsFilter returns a stop words filter.
func NewStopWordsFilter() *token.StopTokensFilter {
	return token.NewStopTokensFilter(StopWords())
}
