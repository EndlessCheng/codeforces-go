package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/848/B
// https://codeforces.com/problemset/status/848/problem/B
func TestCF848B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 10 8
1 1 10
1 4 13
1 7 1
1 8 2
2 2 0
2 5 14
2 6 0
2 6 1
outputCopy
4 8
10 5
8 8
10 6
10 2
1 8
7 8
10 6
inputCopy
3 2 3
1 1 2
2 1 1
1 1 5
outputCopy
1 3
2 1
1 3
inputCopy
10 500 500
2 88 59
2 470 441
1 340 500
2 326 297
1 74 45
1 302 273
1 132 103
2 388 359
1 97 68
2 494 465
outputCopy
500 494
97 500
340 500
302 500
500 470
500 88
500 326
132 500
500 388
74 500`
	testutil.AssertEqualCase(t, rawText, -1, CF848B)
}
