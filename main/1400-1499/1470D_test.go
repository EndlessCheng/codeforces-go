package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1470D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3 2
3 2
2 1
4 2
1 4
2 3
outputCopy
YES
2
1 3 
NO
inputCopy
1
17 27
1 8
2 9
3 10
4 11
5 12
6 13
7 14
8 9
8 14
8 15
9 10
9 15
10 11
10 15
10 17
11 12
11 17
12 13
12 16
12 17
13 14
13 16
14 16
14 15
15 16
15 17
16 17
outputCopy
YES
8
1 3 4 5 6 9 14 17 `
	testutil.AssertEqualCase(t, rawText, 0, CF1470D)
}
