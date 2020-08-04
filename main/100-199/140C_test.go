package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol140C(t *testing.T) {
	// just copy from website
	rawText := `
7
1 2 3 4 5 6 7
outputCopy
2
3 2 1
6 5 4
inputCopy
3
2 2 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, Sol140C)
}
