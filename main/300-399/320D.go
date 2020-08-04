package _00_399

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol320D(reader io.Reader, writer io.Writer) {
	max := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val > ans {
				ans = val
			}
		}
		return ans
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}

	ans := make([]int, n)
	posStack := []int{}
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for len(posStack) > 0 && arr[posStack[len(posStack)-1]] < arr[i] {
			var pos int
			posStack, pos = posStack[:len(posStack)-1], posStack[len(posStack)-1]
			cnt = max(cnt+1, ans[pos])
			ans[i] = cnt
		}
		posStack = append(posStack, i)
	}
	Fprint(out, max(ans...))
}

//func main() {
//	Sol320D(os.Stdin, os.Stdout)
//}
