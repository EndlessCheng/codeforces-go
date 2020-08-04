package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF219D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 1
2 3
outputCopy
0
2 
inputCopy
4
1 4
2 4
3 4
outputCopy
2
1 2 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF219D)
}
