package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1334C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ hp, dmg int64 }

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		ans := int64(0)
		Fscan(in, &n)
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].hp, &a[i].dmg)
		}
		minDmg := int64(1e18)
		for i, p := range a {
			if hp := a[(i+1)%n].hp; p.dmg > hp {
				p.dmg = hp
			}
			ans += p.hp - p.dmg
			if p.dmg < minDmg {
				minDmg = p.dmg
			}
		}
		Fprintln(out, ans+minDmg)
	}
}

//func main() { CF1334C(os.Stdin, os.Stdout) }
