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

	// TODO: 测试边界情况
	customTestCases := [][2]string{
		{
			``,
			``,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		testutil.AssertEqualStringCase(t, customTestCases, 0, run)
		t.Log("======= custom =======")
	}

	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, run)
}

// 无尽对拍 / 构造 hack 数据
// 如果是 special judge，请用 TestCheck 来对拍
// rand.Seed(time.Now().UnixNano())
func TestCompare(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		rg.NewLine()
		rg.IntSlice(n, 1, n)
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0

		Fprint(out, ans)
	}

	// 先用 runBF 跑下样例，大致检查下 runBF 的正确性
	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, runBF)
	return

	// 对拍
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
		n := rg.Int(1, 10)
		rg.NewLine()
		a := rg.IntSlice(n, 1, n)
		//Println(rg.String())
		return rg.String(), func(output string) (_b bool) {
			// 检查 output 是否符合题目要求
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 output 是否满足最优解等
			in := strings.NewReader(output)
			var outN int
			Fscan(in, &outN)
			if !assert.Equal(n, outN) {
				return
			}

			outA := make([]int, outN)
			for i := range outA {
				Fscan(in, &outA[i])
			}
			if !assert.EqualValues(a, outA) {
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
