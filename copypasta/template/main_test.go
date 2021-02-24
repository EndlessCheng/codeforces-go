package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	dir, _ := filepath.Abs(".")
	t.Logf("Current problem is [%s]", filepath.Base(dir))

	customTestCases := [][2]string{
		// TODO: 优先测试边界
		{
			``,
			``,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		tarCase := 0 // -1
		testutil.AssertEqualStringCase(t, customTestCases, tarCase, run)
		t.Log("======= custom =======")
	}

	tarCase := 0 // -1
	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", tarCase, run)
}

// 无尽对拍 / 构造 hack 数据
// 如果是 special judge，请用 TestCheck 来对拍
// rand.Seed(time.Now().UnixNano())
func TestCompare(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		rg.NewLine()
		rg.IntSlice(n, 0, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0

		Fprintln(out, ans)
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, runBF)
	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
	return

	// for hacking, write wrong codes in runBF
	testutil.AssertEqualRunResultsInf(t, inputGenerator, run, runBF)
}

// 无尽检查输出是否正确 / 构造 hack 数据
func TestCheck(t *testing.T) {
	return
	assert := assert.New(t)

	inputGenerator := func() (string, testutil.OutputChecker) {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		rg.NewLine()
		a := rg.IntSlice(n, 0, 5)
		return rg.String(), func(myOutput string) (_b bool) {
			// 检查 myOutput 是否符合题目要求
			// * 最好重新看一遍题目描述以免漏判 *
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 myOutput 是否满足最优解等
			in := strings.NewReader(myOutput)

			myA := make([]int, n)
			for i := range myA {
				Fscan(in, &myA[i])
			}
			if !assert.EqualValues(a, myA) {
				return
			}

			return true
		}
	}

	testutil.CheckRunResultsInf(t, inputGenerator, run)
	return

	// for hacking, write wrong codes here
	runHack := func(in io.Reader, out io.Writer) {

	}
	testutil.CheckRunResultsInf(t, inputGenerator, runHack)
}
