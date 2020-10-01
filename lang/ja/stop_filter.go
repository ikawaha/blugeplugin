package ja

import (
	"github.com/blugelabs/bluge/analysis/token"
)

func StopWordsFilter() *token.StopTokensFilter {
	return token.NewStopTokensFilter(StopWords())
}
