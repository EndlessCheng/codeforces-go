package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol190C(t *testing.T) {
	// just copy from website
	rawText := `
2
int int
outputCopy
Error occurred
inputCopy
3
pair pair int int int
outputCopy
pair<pair<int,int>,int>
inputCopy
1
pair int
outputCopy
Error occurred`
	testutil.AssertEqual(t, rawText, Sol190C)
}
