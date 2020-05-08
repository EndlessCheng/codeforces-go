package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1351B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3
3 1
3 2
1 3
3 3
1 3
outputCopy
Yes
Yes
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1351B)
}
