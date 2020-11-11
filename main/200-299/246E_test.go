package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF246E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
pasha 0
gerald 1
gerald 1
valera 2
igor 3
olesya 1
5
1 1
1 2
1 3
3 1
6 1
outputCopy
2
2
0
1
0
inputCopy
6
valera 0
valera 1
valera 1
gerald 0
valera 4
kolya 4
7
1 1
1 2
2 1
2 2
4 1
5 1
6 1
outputCopy
1
0
0
0
2
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF246E)
}
