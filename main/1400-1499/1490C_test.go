package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/C
// https://codeforces.com/problemset/status/1490/problem/C
func TestCF1490C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1
2
4
34
35
16
703657519796
outputCopy
NO
YES
NO
NO
YES
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1490C)
}
