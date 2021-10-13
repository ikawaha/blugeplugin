package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
	"github.com/ikawaha/kagome-dict/ipa"
	"golang.org/x/text/unicode/norm"
)

// Analyzer returns the analyzer suite in Japanese.
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		CharFilters: []analysis.CharFilter{
			NewUnicodeNormalizeCharFilter(norm.NFKC),
		},
		Tokenizer: NewJapaneseTokenizer(
			ipa.Dict(),
			StopTagsFilter(DefaultStopTags()),
			BaseFormFilter(DefaultInflected),
		),
		TokenFilters: []analysis.TokenFilter{
			token.NewLowerCaseFilter(),
			NewStopWordsFilter(DefaultStopWords()),
		},
	}
}
