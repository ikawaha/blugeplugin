package ja

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/blugelabs/bluge/analysis"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var splitter = filter.SentenceSplitter{
	Delim:               []rune{'。', '．', '！', '!', '？', '?'},
	Follower:            []rune{'.', '｣', '」', '』', ')', '）', '｝', '}', '〉', '》'},
	SkipWhiteSpace:      false,
	DoubleLineFeedSplit: true,
	MaxRuneLen:          128,
}

// DefaultInflected represents POSs which has inflected form.
var DefaultInflected = filter.NewPOSFilter(
	filter.POS{"動詞"},
	filter.POS{"形容詞"},
	filter.POS{"形容動詞"},
)

// TokenizerOption represents an option of the japanese tokenizer.
type TokenizerOption func(t *JapaneseTokenizer)

// DefaultStopTagsFilter returns a default tags filter option.
// See the StopTags function.
func DefaultStopTagsFilter() TokenizerOption {
	tags := StopTags()
	ps := make([]filter.POS, 0, len(tags))
	for k := range tags {
		ps = append(ps, strings.Split(k, "-"))
	}
	return StopTagsFilter(filter.NewPOSFilter(ps...))
}

// StopTagsFilter returns a stop tags filter option.
func StopTagsFilter(f *filter.POSFilter) TokenizerOption {
	return func(t *JapaneseTokenizer) {
		t.stopTagFilter = f
	}
}

// DefaultBaseFormFilter returns a default base form filter option.
// See DefaultInflected.
func DefaultBaseFormFilter() TokenizerOption {
	return BaseFormFilter(DefaultInflected)
}

// BaseFormFilter returns an base form filter option.
func BaseFormFilter(f *filter.POSFilter) TokenizerOption {
	return func(t *JapaneseTokenizer) {
		t.baseFormFilter = f
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
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(splitter.ScanSentences)
	base := 0
	prevIncr := 0
	var ret analysis.TokenStream
	for scanner.Scan() {
		inp := scanner.Text()
		tokens := t.Analyze(inp, tokenizer.Search)
		before := len(tokens)
		if t.stopTagFilter != nil {
			t.stopTagFilter.Drop(&tokens)
		}
		after := 0
		if len(tokens) > 0 {
			after = tokens[len(tokens)-1].Index + 1
		}
		for i, v := range tokens {
			start := base + v.Position
			end := base + v.Position + len(v.Surface)
			term := input[start:end]
			if t.baseFormFilter != nil {
				if pos := v.POS(); t.baseFormFilter.Match(pos) {
					if base, ok := v.BaseForm(); ok {
						term = []byte(base)
					}
				}
			}
			incr := 0
			if i == 0 {
				incr = prevIncr + v.Index + 1
				prevIncr = 0
			} else {
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
		base += len(inp)
		prevIncr = prevIncr + (before - after)
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
