// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var examples = [][]string{
	{
		`[1,2,5,3,4]`, `[1,4,2,5,2,2]`,
		`[41,71,0]`,
	},
	// TODO 测试参数的下界和上界

}

func Test(t *testing.T) {
	t.Log("Current test is [b]")

	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getSum, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://ac.nowcoder.com/acm/contest/9476/b

func TestCompare(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() (a []int, q []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 40)
		a = rg.IntSlice(n, 1, 1e5)
		qq := rg.Int(1, 20)
		for i := 0; i < qq; i++ {
			l := rg.Int(1, n)
			r := rg.Int(l, n)
			q = append(q, l, r)
		}
		return
	}

	runAC := func(a []int, q []int) (ans []int) {
		// 若要修改 a，必须先 copy 一份，在 copied 上修改
		for ii := 0; ii < len(q); ii += 2 {
			l, r := q[ii], q[ii+1]
			s := 0
			for i := l; i <= r; i++ {
				for j := i + 1; j <= r; j++ {
					s += a[i-1] * a[j-1]
				}
			}
			ans = append(ans, s%(1e9+7))
		}
		return
	}

	// test examples first (or make it global)

	if err := testutil.RunLeetCodeFuncWithExamples(t, runAC, examples, 0); err != nil {
		t.Fatal(err)
	}
	testutil.CompareInf(t, inputGenerator, runAC, getSum /*TODO*/)
}
