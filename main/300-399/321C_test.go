package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF321C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
1 4
outputCopy
A B B B
inputCopy
10
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
outputCopy
D C B A D C B D C D`
	testutil.AssertEqualCase(t, rawText, 0, CF321C)
}
