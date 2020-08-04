package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF20B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 -5 6
outputCopy
2
2.0000000000
3.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF20B)
}
