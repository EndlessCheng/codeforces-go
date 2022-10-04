package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/665/E
// https://codeforces.com/problemset/status/665/problem/E
func TestCF665E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
1 2 3
outputCopy
5
inputCopy
3 2
1 2 3
outputCopy
3
inputCopy
3 3
1 2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF665E)
}
