package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	t := append(make([]byte, 0, len(s)*2+3), '^')
	for _, c := range s {
		t = append(t, '#', byte(c))
	}
	t = append(t, '#', '$')
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	mx := 0
	for i := 2; i < len(halfLen); i++ { 
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
		right := (i+hl)/2 - 2
		if right == len(s)-1 {
			mx = max(mx, hl-1)
		}
	}
	
	t = []byte(s[:len(s)-mx])
	slices.Reverse(t)
	Fprint(out, s+string(t))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
