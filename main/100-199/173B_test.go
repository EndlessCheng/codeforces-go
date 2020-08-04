package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF173B(t *testing.T) {
	// just copy from website
	rawText := `
3 3
.#.
...
.#.
outputCopy
2
inputCopy
4 3
##.
...
.#.
.#.
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF173B)
}
