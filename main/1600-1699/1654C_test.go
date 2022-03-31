package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1654/C
// https://codeforces.com/problemset/status/1654/problem/C
func TestCF1654C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
14
1
327
2
869 541
2
985214736 985214737
3
2 3 1
3
2 3 3
6
1 1 1 1 1 1
6
100 100 100 100 100 100
8
100 100 100 100 100 100 100 100
8
2 16 1 8 64 1 4 32
10
1 2 4 7 1 1 1 1 7 2
10
7 1 1 1 3 1 3 3 2 3
10
1 4 4 1 1 1 3 3 3 1
10
2 3 2 2 1 2 2 2 2 2
4
999999999 999999999 999999999 999999999
outputCopy
YES
NO
YES
YES
NO
YES
NO
YES
YES
YES
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1654C)
}
