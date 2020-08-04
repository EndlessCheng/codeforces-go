package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF274B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
1 3
1 -1 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF274B)
}
