package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1738/C
// https://codeforces.com/problemset/status/1738/problem/C
func TestCF1738C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 3 5
4
1 3 5 7
4
1 2 3 4
4
10 20 30 40
outputCopy
Alice
Alice
Bob
Alice`
	testutil.AssertEqualCase(t, rawText, 0, CF1738C)
}
