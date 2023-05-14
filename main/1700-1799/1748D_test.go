package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1748/D
// https://codeforces.com/problemset/status/1748/problem/D
func TestCF1748D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
12 39 5
6 8 14
100 200 200
3 4 6
2 2 2
18 27 3
420 666 69
987654321 123456789 999999999
outputCopy
2147483695
14
-1
-1
2
27
44023415742
586760266413239733`
	testutil.AssertEqualCase(t, rawText, 0, CF1748D)
}
