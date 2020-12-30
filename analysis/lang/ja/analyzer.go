package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
	"golang.org/x/text/unicode/norm"
)

// Analyzer returns the analyzer suite in Japanese.
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		CharFilters: []analysis.CharFilter{
			NewUnicodeNormalizeCharFilter(norm.NFKC),
		},
		Tokenizer: NewJapaneseTokenizer(StopTagsFilter(), BaseFormFilter()),
		TokenFilters: []analysis.TokenFilter{
			token.NewLowerCaseFilter(),
			NewStopWordsFilter(),
		},
	}
}
