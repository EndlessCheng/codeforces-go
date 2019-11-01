package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func SolP3864(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	mp := map[string]byte{
		"ABC": '2', "DEF": '3', "GHI": '4', "JKL": '5', "MNO": '6', "PRS": '7', "TUV": '8', "WXY": '9',
	}
	var s string
	Fscan(in, &s)
	ans := []string{}
	for i := 0; i < 4617; i++ {
		var name string
		if n, _ := Fscan(in, &name); n == 0 {
			break
		}
		if len(name) != len(s) {
			continue
		}
		ok := true
		for i, c := range name {
			ok2 := false
			for k, v := range mp {
				if strings.Contains(k, string(c)) && v == s[i] {
					ok2 = true
					break
				}
			}
			if !ok2 {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, name)
		}
	}
	if len(ans) == 0 {
		Fprintln(out, "NONE")
		return
	}
	sort.Strings(ans)
	for _, s := range ans {
		Fprintln(out, s)
	}
}

//func main() {
//	SolP3864(os.Stdin, os.Stdout)
//}
