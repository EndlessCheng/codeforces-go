package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol20C(t *testing.T) {
	// just copy from website
	rawText := `
5 6
1 2 2
2 5 5
2 3 4
1 4 1
4 3 3
3 5 1
outputCopy
1 4 3 5 
inputCopy
5 6
1 2 2
2 5 5
2 3 4
1 4 1
4 3 3
3 5 1
outputCopy
1 4 3 5 `
	testutil.AssertEqual(t, rawText, CF20C)
}
