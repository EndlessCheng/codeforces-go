package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1073/E
// https://codeforces.com/problemset/status/1073/problem/E
func TestCF1073E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 50 2
outputCopy
1230
inputCopy
1 2345 10
outputCopy
2750685
inputCopy
101 154 2
outputCopy
2189
inputCopy
427896435961371452 630581697708338740 1
outputCopy
716070897
inputCopy
1 99 1
outputCopy
540`
	testutil.AssertEqualCase(t, rawText, 0, CF1073E)
}
