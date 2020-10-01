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
			title: "filtered results",
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
