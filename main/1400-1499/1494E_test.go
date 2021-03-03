package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1494/problem/E
// https://codeforces.com/problemset/status/1494/problem/E
func TestCF1494E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 11
+ 1 2 a
+ 2 3 b
+ 3 2 a
+ 2 1 b
? 3
? 2
- 2 1
- 3 2
+ 2 1 c
+ 3 2 d
? 5
outputCopy
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1494E)
}
