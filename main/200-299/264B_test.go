package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF264B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 3 4 6 9
outputCopy
4
inputCopy
9
1 2 3 5 6 7 8 9 10
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF264B)
}
