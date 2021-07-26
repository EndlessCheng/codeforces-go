package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1149/B
// https://codeforces.com/problemset/status/1149/problem/B
func TestCF1149B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 8
abdabc
+ 1 a
+ 1 d
+ 2 b
+ 2 c
+ 3 a
+ 3 b
+ 1 c
- 2
outputCopy
YES
YES
YES
YES
YES
YES
NO
YES
inputCopy
6 8
abbaab
+ 1 a
+ 2 a
+ 3 a
+ 1 b
+ 2 b
+ 3 b
- 1
+ 2 z
outputCopy
YES
YES
YES
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1149B)
}
