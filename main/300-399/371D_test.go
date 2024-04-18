package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol371D(t *testing.T) {
	// just copy from website
	rawText := `
2
5 10
6
1 1 4
2 1
1 2 5
1 1 4
2 1
2 2
outputCopy
4
5
8
inputCopy
3
5 10 8
6
1 1 12
2 2
1 1 6
1 3 2
2 2
2 3
outputCopy
7
10
5`
	testutil.AssertEqual(t, rawText, cf371D)
}
