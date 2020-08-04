package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF283B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 4 1
outputCopy
3
6
8
inputCopy
3
1 2
outputCopy
-1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF283B)
}
