package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1335/D
// https://codeforces.com/problemset/status/1335/problem/D
func TestCF1335D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
154873296
386592714
729641835
863725149
975314628
412968357
631457982
598236471
247189563
outputCopy
154873396
336592714
729645835
863725145
979314628
412958357
631457992
998236471
247789563`
	testutil.AssertEqualCase(t, rawText, 0, CF1335D)
}
