package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF269B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
2 1
1 2.0
1 3.100
outputCopy
1
inputCopy
3 3
1 5.0
2 5.5
3 6.0
outputCopy
0
inputCopy
6 3
1 14.284235
2 17.921382
1 20.328172
3 20.842331
1 25.790145
1 27.204125
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF269B)
}
