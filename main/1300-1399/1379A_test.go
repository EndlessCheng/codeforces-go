package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1379A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
7
abacaba
7
???????
11
aba?abacaba
11
abacaba?aba
15
asdf???f???qwer
11
abacabacaba
outputCopy
Yes
abacaba
Yes
abacaba
Yes
abadabacaba
Yes
abacabadaba
No
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1379A)
}
