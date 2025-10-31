package P2000_2999

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2375(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		diff := make([]int, n+1)
		z := make([]int, n)
		boxL, boxR := 0, 0
		for i := 1; i < n; i++ {
			if i <= boxR {
				z[i] = min(z[i-boxL], boxR-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				boxL, boxR = i, i+z[i]
				z[i]++
			}
			diff[i]++
			diff[i+min(z[i], i)]--
		}
		ans, sd := 1, 1
		for _, d := range diff[:n] {
			sd = (sd + d + mod) % mod
			ans = ans * sd % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { p2375(bufio.NewReader(os.Stdin), os.Stdout) }
