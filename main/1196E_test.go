package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1196E(t *testing.T) {
	// just copy from website
	rawText := `3
1 1
1 4
2 5
outputCopy
YES
2 2
1 2
YES
2 3
1 3
3 3
2 2
2 4
YES
2 3
2 4
2 5
1 3
1 5
3 3
3 5`
	testutil.AssertEqual(t, rawText, Sol1196E)
}
