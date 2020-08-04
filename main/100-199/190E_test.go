package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol190E(t *testing.T) {
	// just copy from website
	rawText := `
4 4
1 2
1 3
4 2
4 3
outputCopy
2
2 1 4
2 2 3
inputCopy
3 1
1 2
outputCopy
1
3 1 2 3`
	testutil.AssertEqualCase(t, rawText, 0, Sol190E)
}
