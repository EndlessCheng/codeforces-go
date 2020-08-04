package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol5C(t *testing.T) {
	// just copy from website
	rawText := `
(()())()(())()()())())()((()(()(())()()())((()(())()(()()()()))()(())()(((()())()(()((())()(())(()))
outputCopy
28 1
inputCopy
((((()(((
outputCopy
2 1
inputCopy
)((())))(()())
outputCopy
6 2
inputCopy
))(
outputCopy
0 1`
	testutil.AssertEqualCase(t, rawText, -1, Sol5C)
}
