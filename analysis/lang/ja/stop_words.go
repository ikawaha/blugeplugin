package ja

import (
	_ "embed"

	"github.com/blugelabs/bluge/analysis"
)

// StopWordsBytes is a stop word list.
// see. https://github.com/apache/lucene-solr/blob/master/lucene/analysis/kuromoji/src/resources/org/apache/lucene/analysis/ja/stopwords.txt
//go:embed assets/stop_words.txt
var StopWordsBytes []byte

// StopWords returns a stop words map.
func StopWords() analysis.TokenMap {
	rv := analysis.NewTokenMap()
	rv.LoadBytes(StopWordsBytes)
	return rv
}
