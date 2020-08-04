package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol144D(t *testing.T) {
	// just copy from website
	rawText := `
4 6 1
1 2 1
1 3 3
2 3 1
2 4 1
3 4 1
1 4 2
2
outputCopy
3
inputCopy
5 6 3
3 1 1
3 2 1
3 4 1
3 5 1
1 2 6
4 5 8
4
outputCopy
3
inputCopy
2 1 1
2 1 656
0
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 3, Sol144D)
}
