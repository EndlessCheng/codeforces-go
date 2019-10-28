package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1251D(t *testing.T) {
	// just copy from website
	rawText := `
3
3 26
10 12
1 4
10 11
1 1337
1 1000000000
5 26
4 4
2 4
6 8
5 6
2 7
outputCopy
11
1337
6`
	testutil.AssertEqualCase(t, rawText, 0, Sol1251D)
}
