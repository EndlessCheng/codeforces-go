package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1234/C
// https://codeforces.com/problemset/status/1234/problem/C
func TestCF1234C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
7
2323216
1615124
1
3
4
2
13
24
2
12
34
3
536
345
2
46
54
outputCopy
YES
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1234C)
}
