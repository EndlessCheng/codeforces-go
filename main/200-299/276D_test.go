
package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol276D(t *testing.T) {
	// just copy from website
	rawText := `
1 3
outputCopy
3
inputCopy
1 2
outputCopy
3
inputCopy
8 16
outputCopy
31
inputCopy
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, Sol276D)
}
