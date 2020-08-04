package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF35D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
1 1 1
outputCopy
2
inputCopy
3 6
1 1 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF35D)
}
