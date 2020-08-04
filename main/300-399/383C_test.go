package _00_399

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF383C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2 1 1 2
1 2
1 3
2 4
2 5
1 2 3
1 1 2
2 1
2 2
2 4
outputCopy
3
3
0`
	testutil.AssertEqualCase(t, rawText, 0, CF383C)
}
