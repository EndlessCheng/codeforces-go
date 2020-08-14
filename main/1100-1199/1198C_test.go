package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1198C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
1 2
1 2
1 3
1 2
2 5
1 2
3 1
1 4
5 1
1 6
2 15
1 2
1 3
1 4
1 5
1 6
2 3
2 4
2 5
2 6
3 4
3 5
3 6
4 5
4 6
5 6
outputCopy
Matching
2
IndSet
1
IndSet
2 4
Matching
1 15`
	testutil.AssertEqualCase(t, rawText, 0, CF1198C)
}
