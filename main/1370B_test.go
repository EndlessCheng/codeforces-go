package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1370B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 2 3 4 5 6
2
5 7 9 10
5
1 3 3 4 5 90 100 101 2 3
outputCopy
3 6
4 5
3 4
1 9
2 3
4 5
6 10`
	testutil.AssertEqualCase(t, rawText, 0, CF1370B)
}
