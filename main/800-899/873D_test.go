package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/873/D
// https://codeforces.com/problemset/status/873/problem/D
func TestCF873D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
outputCopy
2 1 3 
inputCopy
4 1
outputCopy
1 2 3 4 
inputCopy
5 6
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF873D)
}
