package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1311B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 2
3 2 1
1 2
4 2
4 1 2 3
3 2
5 1
1 2 3 4 5
1
4 2
2 1 4 3
1 3
4 2
4 3 2 1
1 3
5 2
2 1 2 3 3
1 4
outputCopy
YES
NO
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1311B)
}
