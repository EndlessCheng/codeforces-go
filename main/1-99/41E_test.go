package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF41E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
2
1 2
2 3
inputCopy
4
outputCopy
4
1 2
2 3
3 4
4 1`
	testutil.AssertEqualCase(t, rawText, 0, CF41E)
}
