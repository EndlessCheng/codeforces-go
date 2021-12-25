package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1600/E
// https://codeforces.com/problemset/status/1600/problem/E
func TestCF1600E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
5
outputCopy
Alice
inputCopy
3
5 4 5
outputCopy
Alice
inputCopy
6
5 8 2 1 10 9
outputCopy
Bob`
	testutil.AssertEqualCase(t, rawText, 0, CF1600E)
}
