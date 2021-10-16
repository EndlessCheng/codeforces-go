package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1056/D
// https://codeforces.com/problemset/status/1056/problem/D
func TestCF1056D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1
outputCopy
1 1 2 
inputCopy
5
1 1 3 3
outputCopy
1 1 1 2 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1056D)
}
