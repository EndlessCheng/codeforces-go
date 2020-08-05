package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol484A(t *testing.T) {
	// just copy from website
	rawText := `
7
4 5
4 6
12 14
1 1
1 2
2 4
1 10
outputCopy
5
5
13
1
1
3
7`
	testutil.AssertEqual(t, rawText, Sol484A)
}
