package ja

import (
	"github.com/blugelabs/bluge/analysis"
)

// DefaultInflected represents POSs which has inflected form.
var DefaultInflected = analysis.TokenMap{
	"動詞":   true,
	"形容詞":  true,
	"形容動詞": true,
}
