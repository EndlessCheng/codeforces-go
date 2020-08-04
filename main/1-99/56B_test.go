package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF56B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
1 6 5 4 3 2 7 8
outputCopy
2 6
inputCopy
4
2 3 4 1
outputCopy
0 0
inputCopy
4
1 2 3 4
outputCopy
0 0`
	testutil.AssertEqualCase(t, rawText, 0, CF56B)
}
