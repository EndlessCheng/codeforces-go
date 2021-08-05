package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol721C(t *testing.T) {
	// just copy from website
	rawText := `
4 3 13
1 2 5
2 3 7
2 4 8
outputCopy
3
1 2 4
inputCopy
6 6 7
1 2 2
1 3 3
3 6 3
2 4 2
4 6 2
6 5 1
outputCopy
4
1 2 4 6
inputCopy
5 5 6
1 3 3
3 5 3
1 2 2
2 4 3
4 5 2
outputCopy
3
1 3 5`
	testutil.AssertEqualCase(t, rawText, -1, CF721C)
}
