package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF914D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 6 3
4
1 1 2 2
1 1 3 3
2 1 9
1 1 3 2
outputCopy
YES
YES
NO
inputCopy
5
1 2 3 4 5
6
1 1 4 2
2 3 6
1 1 4 2
1 1 5 2
2 5 10
1 1 5 2
outputCopy
NO
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF914D)
}
