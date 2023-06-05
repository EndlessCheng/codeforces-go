package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1838/problem/C
// https://codeforces.com/problemset/status/1838/problem/C
func TestCF1838C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 4
5 7
6 4
outputCopy
16  7  1  9
12  8  2  3
13  4 10 11
14  5  6 15

29 23 17  9  5  6  2
33 27 21 15 11  7  1
32 31 25 19 20 16 10
26 30 24 18 14  8  4
35 34 28 22 13 12  3

 2  3  7 11
 8  9  1 10
17 13  5  4
18 14  6 12
19 23 15 21
20 24 16 22`
	testutil.AssertEqualCase(t, rawText, 0, CF1838C)
}
