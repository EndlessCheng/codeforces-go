package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf17C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	const mod = 51123987
	const maxN = 152
	const maxM = 52
	var n int
	var s string
	Fscan(in, &n, &s)
	s = " " + s
	t := [3][maxN]int32{}
	for x := n; x >= 1; x-- {
		for y := range 3 {
			t[y][x] = t[y][x+1]
		}
		t[s[x]-'a'][x] = int32(x)
	}

	f := [maxN][maxM][maxM][maxM]int32{}
	f[1][0][0][0] = 1
	ans := int32(0)
	lim := (n + 2) / 3
	for l := 1; l <= n; l++ {
		for x := range lim + 1 {
			for y := range lim + 1 {
				for z := range lim + 1 {
					c := f[l][x][y][z]
					f[t[0][l]][x+1][y][z] = (f[t[0][l]][x+1][y][z] + c) % mod
					f[t[1][l]][x][y+1][z] = (f[t[1][l]][x][y+1][z] + c) % mod
					f[t[2][l]][x][y][z+1] = (f[t[2][l]][x][y][z+1] + c) % mod
					if x+y+z == n && abs(x-y) <= 1 && abs(y-z) <= 1 && abs(z-x) <= 1 {
						ans = (ans + c) % mod
					}
				}
			}
		}
	}

	Fprintln(out, ans)
}

//func main() { cf17C(bufio.NewReader(os.Stdin), os.Stdout) }
