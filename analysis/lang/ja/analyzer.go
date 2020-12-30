package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
	"golang.org/x/text/unicode/norm"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		CharFilters: []analysis.CharFilter{
			NewUnicodeNormalizeCharFilter(norm.NFKC),
		},
		Tokenizer: Tokenizer(StopTagsFilter(), BaseFormFilter()),
		TokenFilters: []analysis.TokenFilter{
			token.NewLowerCaseFilter(),
			StopWordsFilter(),
		},
	}
}
