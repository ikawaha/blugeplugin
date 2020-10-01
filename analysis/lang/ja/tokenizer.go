package ja

import (
	"strings"
	"unsafe"

	"github.com/blugelabs/bluge/analysis"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type JapaneseTokenizer struct{
	filter *filter.POSFilter
}

func (t *JapaneseTokenizer) Tokenize(input []byte) analysis.TokenStream{
	tnz, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	tokens := tnz.Analyze(*(*string)(unsafe.Pointer(&input)), tokenizer.Search)
	ret := make(analysis.TokenStream, 0, len(tokens))
	for _, v := range tokens {
		start:= v.Position
		end := v.Position+len(v.Surface)
		ret = append(ret, &analysis.Token{
			Start:        start,
			End:          end,
			Term:         input[start:end],
			PositionIncr: 1,
			Type:         analysis.Ideographic,
			KeyWord:      false,
		})
	}
	return t.stopTagsFilter(ret, tokens)
}

func Tokenizer(stopTags analysis.TokenMap) analysis.Tokenizer {
	ps := make([]filter.POS, 0, len(stopTags))
	for k := range stopTags {
		ps = append(ps, strings.Split(k, "-"))
	}
	return &JapaneseTokenizer{
		filter: filter.NewPOSFilter(ps...),
	}
}

func (t JapaneseTokenizer) stopTagsFilter(input analysis.TokenStream, tokens []tokenizer.Token) analysis.TokenStream {
	var tail, skipped int
	for i, v := range input {
		if t.filter.Match(tokens[i].POS()) {
			skipped += v.PositionIncr
		} else {
			v.PositionIncr += skipped
			skipped = 0
			input[tail] = v
			tail++
		}
	}
	return input[:tail]
}