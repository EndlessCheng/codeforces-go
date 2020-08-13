package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol671A(t *testing.T) {
	// just copy from website
	rawText := `
3 1 1 2 0 0
3
1 1
2 1
2 3
outputCopy
11.084259940083
inputCopy
5 0 4 2 2 0
5
5 2
3 0
5 5
3 5
3 3
outputCopy
33.121375178000`
	testutil.AssertEqualCase(t, rawText, 0, Sol671A)
}