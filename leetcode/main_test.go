package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompareInf(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil2.NewRandGenerator()
	inputGenerator := func() (a []int) {
		//return
		rg.Clear()
		n := rg.Int(1, 3)
		a = rg.IntSlice(n, 1, 3)
		return
	}

	runAC := func(a []int) (ans int) {
		// 若要修改 a，必须先 copy 一份，在 copied 上修改

		return
	}

	// test examples first
	if err := testutil.RunLeetCodeFuncWithFile(_t, runAC, "d.txt", 0); err != nil {
		_t.Fatal(err)
	}
	return

	testutil.CompareInf(_t, inputGenerator, runAC, nil /*TODO*/)
}

type Foo struct{}

func Constructor(int) (_ Foo) { return }
func (Foo) F(int) (_ int)     { return }
func (Foo) G(int) (_ int)     { return }

func TestCompareClassInf(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	assert := assert.New(_t)
	rg := testutil2.NewRandGenerator()
	for tc := 1; ; tc++ {
		inputInfo := &strings.Builder{}
		rg.Clear()
		n := rg.Int(1, 5)
		m := rg.Int(1, 5)
		inputInfo.WriteString(fmt.Sprintln("Constructor", n)) //
		obj := Constructor(n)

		// 暴力数据-初始化

		for _i := 1; _i <= m; _i++ {
			switch rg.Int(0, 1) { //
			case 0:
				v := rg.Int(1, 5)
				inputInfo.WriteString(fmt.Sprintln("F", v)) //

				// 暴力数据-计算（如有必要则复制随机数据）
				var expectedAns int

				myAns := obj.F(v)
				assert.EqualValues(expectedAns, myAns, "Wrong Answer %d\nInput:\n%v", tc, inputInfo)
			case 1:
				v := rg.Int(1, 5)
				inputInfo.WriteString(fmt.Sprintln("G", v)) //

				var expectedAns int

				myAns := obj.G(v)
				assert.EqualValues(expectedAns, myAns, "Wrong Answer %d\nInput:\n%v", tc, inputInfo)
			default:
				panic("invalid op")
			}
		}

		if tc&(tc-1) == 0 {
			_t.Logf("%d cases checked.", tc)
		}
	}
}

func TestCheckInf(_t *testing.T) {
	var solve func([]int) []int /*TODO*/
	rg := testutil2.NewRandGenerator()
	for tc := 1; ; tc++ {
		rg.Clear()
		n := rg.Int(1, 9)
		a := rg.IntSlice(n, 1, 9)
		myAns := solve(a)
		// check myAns is valid ...
		_ = myAns

		if tc&(tc-1) == 0 {
			_t.Logf("%d cases checked.", tc)
		}
	}
}

func Test_checkTodo(t *testing.T) {
	dir := "."
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		code := string(data)
		if !strings.Contains(code, "return") {
			fmt.Println("TODO:", path)
			return nil
		}
		line := strings.Split(code, "\n")
		for i, l := range line {
			l = strings.TrimSpace(l)
			if strings.HasPrefix(l, "func ") && l[len(l)-1] == '{' {
				nextLine := strings.TrimSpace(line[i+1])
				if nextLine == "" || nextLine == "return" {
					fmt.Println("TODO:", path)
				}
				break
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
