package ja

import (
	_ "embed"

	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
)

// StopWordsBytes is a stop word list.
// see. https://github.com/apache/lucene-solr/blob/master/lucene/analysis/kuromoji/src/resources/org/apache/lucene/analysis/ja/stopwords.txt
//go:embed assets/stop_words.txt
var StopWordsBytes []byte

// DefaultStopWords returns a stop words map.
func DefaultStopWords() analysis.TokenMap {
	rv := analysis.NewTokenMap()
	rv.LoadBytes(StopWordsBytes)
	return rv
}

// NewStopWordsFilter returns a stop words filter.
func NewStopWordsFilter(m analysis.TokenMap) *token.StopTokensFilter {
	return token.NewStopTokensFilter(m)
}
