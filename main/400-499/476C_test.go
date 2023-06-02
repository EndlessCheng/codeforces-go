package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/476/C
// https://codeforces.com/problemset/status/476/problem/C
func TestCF476C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
0
inputCopy
2 2
outputCopy
8
inputCopy
10000000 10000000
outputCopy
425362313`
	testutil.AssertEqualCase(t, rawText, 0, CF476C)
}
