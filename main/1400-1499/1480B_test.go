package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1480/B
// https://codeforces.com/problemset/status/1480/problem/B
func TestCF1480B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 17 1
2
16
10 999 3
10 20 30
100 50 30
1000 1000 4
200 300 400 500
1000 1000 1000 1000
999 999 1
1000
1000
999 999 1
1000000
999
outputCopy
YES
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1480B)
}
