package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1186/D
// https://codeforces.com/problemset/status/1186/problem/D
func TestCF1186D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4.58413
1.22491
-2.10517
-3.70387
outputCopy
4
2
-2
-4
inputCopy
5
-6.32509
3.30066
-0.93878
2.00000
1.96321
outputCopy
-6
3
-1
2
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1186D)
}
