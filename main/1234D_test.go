package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1234D(t *testing.T) {
	// just copy from website
	rawText := `
abacaba
5
2 1 4
1 4 b
1 5 b
2 4 6
2 1 7
outputCopy
3
1
2
inputCopy
dfcbbcfeeedbaea
15
1 6 e
1 4 b
2 6 14
1 7 b
1 12 c
2 6 8
2 1 6
1 7 c
1 2 f
1 10 a
2 7 9
1 10 a
1 14 b
1 1 f
2 1 11
outputCopy
5
2
5
2
6`
	testutil.AssertEqual(t, rawText, Sol1234D)
}
