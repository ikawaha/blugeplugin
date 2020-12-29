package ja

import (
	"reflect"
	"testing"

	"github.com/blugelabs/bluge/analysis"
)

func TestJapaneseAnalyzer(t *testing.T) {
	tests := []struct {
		title string
		input  []byte
		output analysis.TokenStream
	}{
		{
			title: "tokenize",
			input: []byte("関西国際空港"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("関西"),
					PositionIncr: 1,
					Start:        0,
					End:          6,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("国際"),
					PositionIncr: 1,
					Start:        6,
					End:          12,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("空港"),
					PositionIncr: 1,
					Start:        12,
					End:          18,
					Type: analysis.Ideographic,
				},
			},
		},
		{
			title: "filtered results:うなぎ",
			input: []byte("私は鰻"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("私"),
					PositionIncr: 1,
					Start:        0,
					End:          3,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("鰻"),
					PositionIncr: 2,
					Start:        6,
					End:          9,
					Type: analysis.Ideographic,
				},
			},
		},
		{
			// 人魚は、南の方の海にばかり棲んでいるのではありません。
			// 人魚	名詞,一般,*,*,*,*,人魚,ニンギョ,ニンギョ
			// は	助詞,係助詞,*,*,*,*,は,ハ,ワ
			// 、	記号,読点,*,*,*,*,、,、,、
			// 南	名詞,一般,*,*,*,*,南,ミナミ,ミナミ
			// の	助詞,連体化,*,*,*,*,の,ノ,ノ
			// 方	名詞,非自立,一般,*,*,*,方,ホウ,ホー
			// の	助詞,連体化,*,*,*,*,の,ノ,ノ
			// 海	名詞,一般,*,*,*,*,海,ウミ,ウミ
			// に	助詞,格助詞,一般,*,*,*,に,ニ,ニ
			// ばかり	助詞,副助詞,*,*,*,*,ばかり,バカリ,バカリ
			// 棲ん	動詞,自立,*,*,五段・マ行,連用タ接続,棲む,スン,スン
			// で	助詞,接続助詞,*,*,*,*,で,デ,デ
			// いる	動詞,非自立,*,*,一段,基本形,いる,イル,イル
			// の	名詞,非自立,一般,*,*,*,の,ノ,ノ
			// で	助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
			// は	助詞,係助詞,*,*,*,*,は,ハ,ワ
			// あり	動詞,自立,*,*,五段・ラ行,連用形,ある,アリ,アリ
			// ませ	助動詞,*,*,*,特殊・マス,未然形,ます,マセ,マセ
			// ん	助動詞,*,*,*,不変化型,基本形,ん,ン,ン
			// 。	記号,句点,*,*,*,*,。,。,。
			title: "filtered results:赤い蝋燭と人魚",
			input: []byte("人魚は、南の方の海にばかり棲んでいるのではありません。"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("人魚"),
					PositionIncr: 1,
					Start:        0,
					End:          6,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("南"),
					PositionIncr: 3,
					Start:        12,
					End:          15,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("方"),
					PositionIncr: 2,
					Start:        18,
					End:          21,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("海"),
					PositionIncr: 2,
					Start:        24,
					End:          27,
					Type: analysis.Ideographic,
				},
				&analysis.Token{
					Term:         []byte("棲ん"),
					PositionIncr: 3,
					Start:        39,
					End:          45,
					Type: analysis.Ideographic,
				},
			},
		},
	}

	analyzer := Analyzer()
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			actual := analyzer.Analyze(test.input)
			if !reflect.DeepEqual(actual, test.output) {
				t.Errorf("want %+v, got %+v", test.output, actual)
			}
		})
	}
}

func BenchmarkJapaneseAnalyzer(b *testing.B){
	sen := []byte("人魚は、南の方の海にばかり棲んでいるのではありません。")
	analyzer := Analyzer()
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		analyzer.Analyze(sen)
	}
}