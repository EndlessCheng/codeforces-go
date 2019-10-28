package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1251D(t *testing.T) {
	// just copy from website
	rawText := `
4

7 1000000
1 1000
2 1000
3 1000
4 4
60 100
80 80
100 100

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
100
11
1337
6`
	testutil.AssertEqualCase(t, rawText, 0, Sol1251D)
}
