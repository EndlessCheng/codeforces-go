package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/995/C
// https://codeforces.com/problemset/status/995/problem/C
func TestCF995C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
999999 0
0 999999
999999 0
outputCopy
1 1 -1 
inputCopy
1
-824590 246031
outputCopy
1 
inputCopy
8
-67761 603277
640586 -396671
46147 -122580
569609 -2112
400 914208
131792 309779
-850150 -486293
5272 721899
outputCopy
1 1 1 1 1 1 1 -1 `
	testutil.AssertEqualCase(t, rawText, 0, CF995C)
}
