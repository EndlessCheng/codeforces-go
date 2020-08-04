package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol613A(t *testing.T) {
	// just copy from website
	rawText := `
3 0 0
0 1
-1 2
1 2
outputCopy
12.566370614359172464
inputCopy
4 1 -1
0 0
1 2
2 0
1 1
outputCopy
21.991148575128551812`
	testutil.AssertEqualCase(t, rawText, 0, Sol613A)
}
