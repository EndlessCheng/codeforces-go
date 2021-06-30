package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/453/B
// https://codeforces.com/problemset/status/453/problem/B
func TestCF453B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1 1 1 1
outputCopy
1 1 1 1 1 
inputCopy
5
1 6 4 2 8
outputCopy
1 5 3 1 8 `
	testutil.AssertEqualCase(t, rawText, 0, CF453B)
}
