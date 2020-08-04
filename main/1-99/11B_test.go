package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF11B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
3
inputCopy
6
outputCopy
3
inputCopy
0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF11B)
}
