package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF617D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 -1
1 1
1 2
outputCopy
1
inputCopy
-1 -1
-1 3
4 3
outputCopy
2
inputCopy
1 1
2 3
3 2
outputCopy
3
inputCopy
-494824697 -964138793
-494824697 671151995
-24543485 877798954
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF617D)
}
