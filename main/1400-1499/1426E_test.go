package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
0 1 1
1 1 0
outputCopy
0 1
inputCopy
15
5 5 5
5 5 5
outputCopy
0 15
inputCopy
3
0 0 3
3 0 0
outputCopy
3 3
inputCopy
686
479 178 29
11 145 530
outputCopy
22 334
inputCopy
319
10 53 256
182 103 34
outputCopy
119 226`
	testutil.AssertEqualCase(t, rawText, 0, CF1426E)
}
