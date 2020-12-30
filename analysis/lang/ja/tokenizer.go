package ja

import (
	"strings"
	"unsafe"

	"github.com/blugelabs/bluge/analysis"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// DefaultInflected represents POSs which has inflected form.
var DefaultInflected = filter.NewPOSFilter(
	filter.POS{"動詞"},
	filter.POS{"形容詞"},
	filter.POS{"形容動詞"},
)

// TokenizerOption represents an option of the japanese tokenizer.
type TokenizerOption func(t *JapaneseTokenizer)

// StopTagsFilter returns a stop tags filter option.
func StopTagsFilter() TokenizerOption {
	return func(t *JapaneseTokenizer) {
		tags := StopTags()
		ps := make([]filter.POS, 0, len(tags))
		for k := range tags {
			ps = append(ps, strings.Split(k, "-"))
		}
		t.stopTagFilter = filter.NewPOSFilter(ps...)
	}
}

// BaseFormFilter returns an base form filter option.
func BaseFormFilter() TokenizerOption {
	return func(t *JapaneseTokenizer) {
		t.baseFormFilter = DefaultInflected
	}
}

// JapaneseTokenizer represents a Japanese tokenizer with filters.
type JapaneseTokenizer struct {
	*tokenizer.Tokenizer
	stopTagFilter  *filter.POSFilter
	baseFormFilter *filter.POSFilter
}

// Tokenize tokenizes the input and filters them.
func (t *JapaneseTokenizer) Tokenize(input []byte) analysis.TokenStream {
	tokens := t.Analyze(*(*string)(unsafe.Pointer(&input)), tokenizer.Search)
	t.stopTagFilter.Drop(&tokens)
	ret := make(analysis.TokenStream, 0, len(tokens))
	for i, v := range tokens {
		start := v.Position
		end := v.Position + len(v.Surface)
		term := input[start:end]
		if pos := v.POS(); t.baseFormFilter.Match(pos) {
			if base, ok := v.BaseForm(); ok {
				term = []byte(base)
			}
		}
		incr := 1
		if i > 0 {
			incr = v.Index - tokens[i-1].Index
		}
		ret = append(ret, &analysis.Token{
			Start:        start,
			End:          end,
			Term:         term,
			PositionIncr: incr,
			Type:         analysis.Ideographic,
			KeyWord:      false,
		})
	}
	return ret
}

// NewJapaneseTokenizer returns a Japanese tokenizer.
func NewJapaneseTokenizer(opts ...TokenizerOption) analysis.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	ret := &JapaneseTokenizer{
		Tokenizer: t,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}
