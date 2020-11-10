package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1009F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
2 3
3 4
outputCopy
0
0
0
0
inputCopy
4
1 2
1 3
1 4
outputCopy
1
0
0
0
inputCopy
4
1 2
2 3
2 4
outputCopy
2
1
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1009F)
}
