package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io/ioutil"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/514/C
// https://codeforces.com/problemset/status/514/problem/C
func TestCF514C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
aaaaa
acacaca
aabaa
ccacacc
caaac
outputCopy
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF514C)
}

func BenchmarkCF514C(b *testing.B) {
	for _i := 0; _i < b.N; _i++ {
		sb := &strings.Builder{}
		n := 1000
		sb.WriteString(fmt.Sprintln(n, 1))
		s := strings.Repeat("a", n-1) + "c\n"
		sb.WriteString(s)
		for preA := 0; preA < n-1; preA++ {
			sb.WriteString(strings.Repeat("a", preA) + "b" + strings.Repeat("a", n-1-preA) + "\n")
		}
		sb.WriteString(s)
		CF514C(strings.NewReader(sb.String()), ioutil.Discard)
	}
}
