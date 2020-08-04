package _00_399

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF349B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 4 3 2 1 2 3 4 5
outputCopy
55555
inputCopy
2
9 11 1 12 5 8 9 10 6
outputCopy
33
inputCopy
0
1 1 1 1 1 1 1 1 1
outputCopy
-1
inputCopy
1000000
1 1 1 1 1 1 1 1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF349B)
}
