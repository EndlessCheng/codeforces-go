package _00_399

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol372A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	sort.Ints(arr)

	ans := n
	n2 := n / 2
	j := n2
	for i := 0; i < n2; i++ {
		for ; j < n; j++ {
			if 2*arr[i] <= arr[j] {
				ans--
				j++
				break
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol372A(os.Stdin, os.Stdout)
//}
