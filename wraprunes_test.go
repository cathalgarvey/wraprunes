package wraprunes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	Value   string
	WrapLen int
}

var (
	testcases = map[tcase][]string{
		//                |         |         |         |         |         |
		//       123456789012345678901234567890123456789012345678901234567890
		tcase{"This is a simple string, that keeps going for a while.", 10}: []string{"This is a ", "simple str", "ing, that ", "keeps goin", "g for a wh", "ile."},
		tcase{"This is a simple string, that keeps going for a while.", 6}:  []string{"This i", "s a si", "mple s", "tring,", " that ", "keeps ", "going ", "for a ", "while."},
		tcase{"This ís a símple ßtring≤ that keeps going for a while.", 6}:  []string{"This í", "s a sí", "mple ß", "tring≤", " that ", "keeps ", "going ", "for a ", "while."},
	}
)

func TestWrapping(t *testing.T) {
	for test, result := range testcases {
		assert.Equal(t, result, Wrap(test.Value, test.WrapLen))
	}
}
