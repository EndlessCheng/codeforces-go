package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol220B(t *testing.T) {
	// just copy from website
	rawText := `
7 2
3 1 2 2 3 3 7
1 7
3 4
outputCopy
3
1`
	testutil.AssertEqualCase(t, rawText, 0, Sol220B)
}
