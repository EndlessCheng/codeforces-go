package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF722C(t *testing.T) {
	// just copy from website
	rawText := `
4
1 3 2 5
3 4 1 2
outputCopy
5
4
3
0
inputCopy
5
1 2 3 4 5
4 2 3 5 1
outputCopy
6
5
5
1
0
inputCopy
8
5 5 4 4 6 6 5 5
5 2 8 7 1 3 4 6
outputCopy
18
16
11
8
8
6
6
0`
	testutil.AssertEqualCase(t, rawText, 0, CF722C)
}
