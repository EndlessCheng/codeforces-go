package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1036/E
// https://codeforces.com/problemset/status/1036/problem/E
func TestCF1036E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
0 0 4 4
-1 5 4 0
4 0 4 4
5 2 11 2
6 1 6 7
5 6 11 6
10 1 10 7
7 0 9 8
10 -1 11 -1
outputCopy
42
inputCopy
4
-1 2 1 2
-1 0 1 0
-1 0 0 3
0 3 1 0
outputCopy
7
inputCopy
10
841746 527518 -841746 -527518
595261 331297 -595261 -331297
-946901 129987 946901 -129987
670374 -140388 -670374 140388
-684770 309555 684770 -309555
-302589 415564 302589 -415564
-387435 613331 387435 -613331
-624940 -95922 624940 95922
945847 -199224 -945847 199224
24636 -565799 -24636 565799
outputCopy
43`
	testutil.AssertEqualCase(t, rawText, -1, CF1036E)
}
