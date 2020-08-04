package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF700B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
1 5 6 2
1 3
3 2
4 5
3 7
4 3
4 6
outputCopy
6
inputCopy
9 3
3 2 1 6 5 9
8 9
3 2
2 7
3 4
7 6
4 5
2 1
2 8
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF700B)
}
