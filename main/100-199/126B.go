package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol126B(reader io.Reader, writer io.Writer) {
	calcMaxMatchLengths := func(pattern string) []int {
		n := len(pattern)
		maxMatchLengths := make([]int, n)
		maxLength := 0
		for i := 1; i < n; i++ {
			c := pattern[i]
			for maxLength > 0 && pattern[maxLength] != c {
				maxLength = maxMatchLengths[maxLength-1]
			}
			if pattern[maxLength] == c {
				maxLength++
			}
			maxMatchLengths[i] = maxLength
		}
		return maxMatchLengths
	}
	kmpSearch := func(text, pattern string) bool {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		count := 0
		for i := range text {
			c := text[i]
			for count > 0 && pattern[count] != c {
				count = maxMatchLengths[count-1]
			}
			if pattern[count] == c {
				count++
			}
			if count == lenP {
				return true
			}
		}
		return false
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const fail = "Just a legend"
	var s string
	Fscan(in, &s)
	if len(s) < 3 {
		Fprint(out, fail)
		return
	}
	inS := s[1 : len(s)-1]

	lens := calcMaxMatchLengths(s)
	for maxLen := lens[len(s)-1]; maxLen > 0; maxLen = lens[maxLen-1] {
		if kmpSearch(inS, s[:maxLen]) {
			Fprint(out, s[:maxLen])
			return
		}
	}
	Fprint(out, fail)
}

//func main() {
//	Sol126B(os.Stdin, os.Stdout)
//}
