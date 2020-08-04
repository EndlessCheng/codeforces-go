package _00_199

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol107B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m, h, v, hv int
	Fscan(in, &n, &m, &h)
	h--
	sum := 0
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		sum += v
		if i == h {
			hv = v
		}
	}
	if sum < n {
		Fprint(out, -1)
		return
	}
	others := sum - hv
	sum--
	n--
	ans := 1.0
	for i := 0; i < n; i++ {
		ans *= float64(others-i) / float64(sum-i)
	}
	Fprint(out, 1-ans)
}

//func main() {
//	Sol107B(os.Stdin, os.Stdout)
//}
