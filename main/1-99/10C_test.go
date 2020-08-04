package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF10C(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	// just copy from website
	rawText := `
4
outputCopy
2
inputCopy
5
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF10C)
}
