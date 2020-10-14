package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1427E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
5
3 + 3
3 ^ 6
3 + 5
3 + 6
8 ^ 9
inputCopy
123
outputCopy
10
123 + 123
123 ^ 246
141 + 123
246 + 123
264 ^ 369
121 + 246
367 ^ 369
30 + 30
60 + 60
120 ^ 121
inputCopy
999999
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1427E)
}
