package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1381B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2
2 3 1 4
2
3 1 2 4
4
3 2 6 1 5 7 8 4
3
1 2 3 4 5 6
4
6 1 3 7 4 5 8 2
6
4 3 2 5 1 11 9 12 8 6 10 7
outputCopy
YES
NO
YES
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1381B)
}
