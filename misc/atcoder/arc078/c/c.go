package main

import (
	. "fmt"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(io func(int) bool) (ans int) {
	allNine := true
	for i := 0; ; i++ {
		d := sort.Search(9, func(d int) bool { return (i > 0 || d > 0) && io((ans+d+1)*int(math.Pow10(11-i))-1) })
		allNine = allNine && d < 9
		ans += d
		if ok := io(ans + 1); allNine && !ok || !allNine && ok {
			return
		}
		ans *= 10
	}
}

func main() {
	Println("!", run(func(q int) bool {
		Println("?", q)
		var s []byte
		Scan(&s)
		return s[0] == 'Y'
	}))
}
