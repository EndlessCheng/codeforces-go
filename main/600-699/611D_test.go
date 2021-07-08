package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/611/D
// https://codeforces.com/problemset/status/611/problem/D
func TestCF611D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
123434
outputCopy
8
inputCopy
8
20152016
outputCopy
4
inputCopy
10
5558403946
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, -1, CF611D)
}
