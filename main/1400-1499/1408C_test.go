package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1408/C
// https://codeforces.com/problemset/status/1408/problem/C
func TestCF1408C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 10
1 9
1 10
1
5 7
1 2 3 4 6
2 1000000000
413470354 982876160
9 478
1 10 25 33 239 445 453 468 477
outputCopy
3.000000000000000
3.666666666666667
2.047619047619048
329737645.750000000000000
53.700000000000000`
	testutil.AssertEqualCase(t, rawText, -1, CF1408C)
}
