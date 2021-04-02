package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/452/C
// https://codeforces.com/problemset/status/452/problem/C
func TestCF452C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
0.6666666666666666
inputCopy
4 4
outputCopy
0.4000000000000000
inputCopy
1 2
outputCopy
1.0000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF452C)
}
