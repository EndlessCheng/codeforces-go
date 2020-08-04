package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF148D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 3
outputCopy
0.500000000
inputCopy
5 5
outputCopy
0.658730159`
	testutil.AssertEqualCase(t, rawText, 0, CF148D)
}
