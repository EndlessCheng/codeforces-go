package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1706/D2
// https://codeforces.com/problemset/status/1706/problem/D1
func TestCF1706D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5 2
4 5 6 8 11
5 12
4 5 6 8 11
3 1
2 9 15
7 3
2 3 5 5 6 9 10
6 56
54 286 527 1436 2450 2681
3 95
16 340 2241
2 2
1 3
outputCopy
2
0
13
1
4
7
0
inputCopy
1
6 56
54 286 527 1436 2450 2681
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1706D2)
}

/*
27 47
47 54
52 57
54 58

 */