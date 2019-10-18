package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func SolP1201(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	names := make([]string, n)
	for i := range names {
		Fscan(in, &names[i])
	}
	left := map[string]int{}
	for _, name := range names {
		left[name] = 0
	}
	for {
		var sender string
		var money, r int
		if n, _ := Fscan(in, &sender, &money, &r); n == 0 {
			break
		}
		if r == 0 {
			continue
		}
		pm := money / r
		left[sender] -= pm * r
		for ; r > 0; r-- {
			var name string
			Fscan(in, &name)
			left[name] += pm
		}
	}
	for _, name := range names {
		Fprintln(out, name, left[name])
	}
}

//func main() {
//	SolP1201(os.Stdin, os.Stdout)
//}
