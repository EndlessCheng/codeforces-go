package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1627/D
// https://codeforces.com/problemset/status/1627/problem/D
func TestCF1627D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 20 1 25 30
outputCopy
3
inputCopy
3
6 10 15
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1627D)
}
