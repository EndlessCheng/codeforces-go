package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1612/problem/D
// https://codeforces.com/problemset/status/1612/problem/D
func TestCF1612D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
6 9 3
15 38 7
18 8 8
30 30 30
40 50 90
24 28 20
365 216 52
537037812705867558 338887693834423551 3199921013340
outputCopy
YES
YES
YES
YES
NO
YES
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1612D)
}
