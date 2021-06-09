package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/486/E
// https://codeforces.com/problemset/status/486/problem/E
func TestCF486E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
4
outputCopy
3
inputCopy
4
1 3 2 5
outputCopy
3223
inputCopy
4
1 5 2 3
outputCopy
3133`
	testutil.AssertEqualCase(t, rawText, 0, CF486E)
}
