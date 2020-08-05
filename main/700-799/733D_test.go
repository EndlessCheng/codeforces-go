package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF733D(t *testing.T) {
	// just copy from website
	rawText := `
6
5 5 5
3 2 4
1 4 1
2 1 3
3 2 4
3 3 4
outputCopy
1
1
inputCopy
7
10 7 8
5 10 3
4 2 6
5 5 5
10 2 8
4 2 1
7 7 7
outputCopy
2
1 5`
	testutil.AssertEqualCase(t, rawText, 0, CF733D)
}
