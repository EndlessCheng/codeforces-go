package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1066/problem/C
// https://codeforces.com/problemset/status/1066/problem/C
func TestCF1066C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
L 1
R 2
R 3
? 2
L 4
? 1
L 5
? 1
outputCopy
1
1
2
inputCopy
10
L 100
R 100000
R 123
L 101
? 123
L 10
R 115
? 100
R 110
? 115
outputCopy
0
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1066C)
}
