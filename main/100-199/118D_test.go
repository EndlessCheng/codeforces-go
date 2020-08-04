package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol118D(t *testing.T) {
	// just copy from website
	rawText := `
2 1 1 10
outputCopy
1
inputCopy
2 3 1 2
outputCopy
5
inputCopy
2 4 1 1
outputCopy
0`
	testutil.AssertEqual(t, rawText, Sol118D)
}
