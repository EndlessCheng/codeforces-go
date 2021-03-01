package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/663/A
// https://codeforces.com/problemset/status/663/problem/A
func TestCF663A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
? + ? - ? + ? + ? = 42
outputCopy
Possible
9 + 13 - 39 + 28 + 31 = 42
inputCopy
? - ? = 1
outputCopy
Impossible
inputCopy
? = 1000000
outputCopy
Possible
1000000 = 1000000`
	testutil.AssertEqualCase(t, rawText, 0, CF663A)
}
