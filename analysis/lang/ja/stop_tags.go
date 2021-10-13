package ja

import (
	_ "embed"

	"github.com/blugelabs/bluge/analysis"
)

// StopTagsBytes is a stop tag list.
// see. https://github.com/apache/lucene-solr/blob/master/lucene/analysis/kuromoji/src/resources/org/apache/lucene/analysis/ja/stoptags.txt
//go:embed assets/stop_tags.txt
var StopTagsBytes []byte

// DefaultStopTags returns a stop tags map (for IPA dict).
func DefaultStopTags() analysis.TokenMap {
	rv := analysis.NewTokenMap()
	rv.LoadBytes(StopTagsBytes)
	return rv
}
