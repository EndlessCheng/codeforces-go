package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1393/B
// https://codeforces.com/problemset/status/1393/problem/B
func TestCF1393B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 1 1 2 1 1
6
+ 2
+ 1
- 1
+ 2
- 1
+ 2
outputCopy
NO
YES
NO
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1393B)
}
