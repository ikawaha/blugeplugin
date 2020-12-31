package ja

import (
	"reflect"
	"testing"

	"github.com/blugelabs/bluge/analysis"
)

func TestJapaneseTokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name  string
		opts  []TokenizerOption
		input []byte
		want  analysis.TokenStream
	}{
		{
			name:  "文分割なし",
			input: []byte("私は鰻"),
			want: analysis.TokenStream{
				{
					Start:        0,
					End:          3,
					Term:         []byte("私"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        3,
					End:          6,
					Term:         []byte("は"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        6,
					End:          9,
					Term:         []byte("鰻"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
			},
		},
		{
			name:  "文分割あり",
			input: []byte("私は鰻。ねこはいます。"),
			want: analysis.TokenStream{
				{
					Start:        0,
					End:          3,
					Term:         []byte("私"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        3,
					End:          6,
					Term:         []byte("は"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        6,
					End:          9,
					Term:         []byte("鰻"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        9,
					End:          12,
					Term:         []byte("。"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        12,
					End:          18,
					Term:         []byte("ねこ"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        18,
					End:          21,
					Term:         []byte("は"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        21,
					End:          24,
					Term:         []byte("い"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        24,
					End:          30,
					Term:         []byte("ます"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        30,
					End:          33,
					Term:         []byte("。"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
			},
		},
		{
			name:  "文分割・フィルターあり",
			input: []byte("私は鰻。ねこはいます。"),
			opts: []TokenizerOption{
				StopTagsFilter(),
				BaseFormFilter(),
			},
			want: analysis.TokenStream{
				{
					Start:        0,
					End:          3,
					Term:         []byte("私"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        6,
					End:          9,
					Term:         []byte("鰻"),
					PositionIncr: 2,
					Type:         analysis.Ideographic,
				},
				{
					Start:        12,
					End:          18,
					Term:         []byte("ねこ"),
					PositionIncr: 2,
					Type:         analysis.Ideographic,
				},
				{
					Start:        21,
					End:          24,
					Term:         []byte("いる"),
					PositionIncr: 2,
					Type:         analysis.Ideographic,
				},
			},
		},
		{
			// Start: 0  End: 3  PositionIncr: 1  Token: 私  Type: 1
			// Start: 3  End: 6  PositionIncr: 1  Token: は  Type: 1 <drop>
			// Start: 6  End: 9  PositionIncr: 1  Token: 鰻  Type: 1
			// Start: 9  End: 12  PositionIncr: 1  Token: 。  Type: 1 <drop>
			// --- 文区切り
			// Start: 12  End: 15  PositionIncr: 1  Token: は  Type: 1 <drop>
			// Start: 15  End: 18  PositionIncr: 1  Token: 。  Type: 1 <drop>
			// --- 文区切り
			// Start: 18  End: 24  PositionIncr: 1  Token: ねこ  Type: 1
			// Start: 24  End: 27  PositionIncr: 1  Token: は  Type: 1 <drop>
			// Start: 27  End: 30  PositionIncr: 1  Token: い  Type: 1
			// Start: 30  End: 36  PositionIncr: 1  Token: ます  Type: 1 <drop>
			// Start: 36  End: 39  PositionIncr: 1  Token: 。  Type: 1 <drop>
			name:  "文ごとDropされるとき",
			input: []byte("私は鰻。は。ねこはいます。"),
			opts: []TokenizerOption{
				StopTagsFilter(),
				BaseFormFilter(),
			},
			want: analysis.TokenStream{
				{
					Start:        0,
					End:          3,
					Term:         []byte("私"),
					PositionIncr: 1,
					Type:         analysis.Ideographic,
				},
				{
					Start:        6,
					End:          9,
					Term:         []byte("鰻"),
					PositionIncr: 2,
					Type:         analysis.Ideographic,
				},
				{
					Start:        18,
					End:          24,
					Term:         []byte("ねこ"),
					PositionIncr: 4,
					Type:         analysis.Ideographic,
				},
				{
					Start:        27,
					End:          30,
					Term:         []byte("いる"),
					PositionIncr: 2,
					Type:         analysis.Ideographic,
				},
			},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			tz := NewJapaneseTokenizer(v.opts...)
			if got := tz.Tokenize(v.input); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %+v, want %+v", got, v.want)
			}
		})
	}
}
