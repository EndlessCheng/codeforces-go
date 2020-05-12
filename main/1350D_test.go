package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1350D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 3
1 5 2 6 1
1 6
6
3 2
1 2 3
4 3
3 1 2 3
10 3
1 2 3 4 5 6 7 8 9 10
outputCopy
no
yes
yes
no
yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1350D)
}
