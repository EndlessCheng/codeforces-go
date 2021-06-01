package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/375/B
// https://codeforces.com/problemset/status/375/problem/B
func TestCF375B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
1
outputCopy
1
inputCopy
2 2
10
11
outputCopy
2
inputCopy
4 3
100
011
000
101
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF375B)
}
