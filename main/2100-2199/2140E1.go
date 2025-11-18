package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2140E1(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([]int, k)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
		}
		if m == 1 {
			Fprintln(out, 1)
			continue
		}

		f := make([]byte, 1<<n)
		f[1] = 1
		for sz := 2; sz <= n; sz++ {
			if (n-sz)%2 == 0 {
			o:
				for mask := 1<<sz - 1; mask > 0; mask-- {
					for _, i := range a {
						if i >= sz {
							break
						}
						if f[mask>>(i+1)<<i|mask&(1<<i-1)] > 0 {
							f[mask] = 1
							continue o
						}
					}
					f[mask] = 0
				}
			} else {
			o2:
				for mask := 1<<sz - 1; mask > 0; mask-- {
					for _, i := range a {
						if i >= sz {
							break
						}
						if f[mask>>(i+1)<<i|mask&(1<<i-1)] == 0 {
							f[mask] = 0
							continue o2
						}
					}
					f[mask] = 1
				}
			}
		}

		ans := 1 << n
		for _, v := range f {
			ans += int(v)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2140E1(bufio.NewReader(os.Stdin), os.Stdout) }
