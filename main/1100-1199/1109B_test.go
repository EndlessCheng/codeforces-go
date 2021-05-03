package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1109/B
// https://codeforces.com/problemset/status/1109/problem/B
func TestCF1109B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
nolon
outputCopy
2
inputCopy
otto
outputCopy
1
inputCopy
qqqq
outputCopy
Impossible
inputCopy
kinnikkinnik
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1109B)
}
