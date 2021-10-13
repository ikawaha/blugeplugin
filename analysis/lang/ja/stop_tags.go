package ja

import (
	_ "embed"

	"github.com/blugelabs/bluge/analysis"
)

// StopTagsBytes is a stop tag list.
// see. https://github.com/apache/lucene-solr/blob/master/lucene/analysis/kuromoji/src/resources/org/apache/lucene/analysis/ja/stoptags.txt
//go:embed assets/stop_tags.txt
var StopTagsBytes []byte

// StopTags returns a stop tags map.
func StopTags() analysis.TokenMap {
	rv := analysis.NewTokenMap()
	rv.LoadBytes(StopTagsBytes)
	return rv
}
