package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1838/problem/E
// https://codeforces.com/problemset/status/1838/problem/E
func TestCF1838E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 1000000 1
1
3 4 3
1 2 2
5 7 8
1 2 3 4 1
6 6 18
18 2 2 5 2 16
1 10 2
1
8 10 1234567
1 1 2 1 2 2 2 1
5 1000000000 1000000000
525785549 816356460 108064697 194447117 725595511
outputCopy
1
9
1079
1
1023
906241579
232432822`
	testutil.AssertEqualCase(t, rawText, 0, CF1838E)
}
