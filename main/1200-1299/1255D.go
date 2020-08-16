package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1255D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var t, r, c, k, lastRi, lastRj, avg, nAvg int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &r, &c, &k)
		cells := make([][]byte, r)
		cnt := 0
		for i := range cells {
			Fscan(in, &cells[i])
			if cntRow := bytes.Count(cells[i], []byte{'R'}); cntRow > 0 {
				cnt += cntRow
				lastRi = i
				if i&1 == 0 {
					lastRj = bytes.LastIndex(cells[i], []byte{'R'})
				} else {
					lastRj = bytes.Index(cells[i], []byte{'R'})
				}
			}
		}
		cells[lastRi][lastRj] = '.'
		if r&1 == 1 {
			cells[r-1][c-1] = 'R'
		} else {
			cells[r-1][0] = 'R'
		}

		avg = cnt / k
		if cnt%k == 0 {
			nAvg = k
		} else {
			nAvg = (avg+1)*k - cnt
		}
		nowI := 0
		nowCnt := 0
		left := nAvg

		modify := func(i, j int, isRice bool) {
			cells[i][j] = alphabet[nowI]
			if isRice {
				nowCnt++
				if nowCnt == avg {
					nowI++
					nowCnt = 0
					left--
					if left == 0 {
						avg++
						left = k - nAvg
					}
				}
			}
		}
		for i, row := range cells {
			if i&1 == 0 {
				for j, ch := range row {
					modify(i, j, ch == 'R')
				}
			} else {
				for j := c - 1; j >= 0; j-- {
					modify(i, j, cells[i][j] == 'R')
				}
			}
		}
		for _, row := range cells {
			Fprintf(out, "%s\n", row)
		}
	}
}

//func main() {
//	Sol1255D(os.Stdin, os.Stdout)
//}
