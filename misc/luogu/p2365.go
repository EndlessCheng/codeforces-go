package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type vec2365 struct{ x, y int }

func (a vec2365) sub(b vec2365) vec2365 { return vec2365{a.x - b.x, a.y - b.y} }
func (a vec2365) det(b vec2365) int     { return a.x*b.y - a.y*b.x }
func (a vec2365) dot(b vec2365) int     { return a.x*b.x + a.y*b.y }

func p2365(in io.Reader, out io.Writer) {
	var n, k, totalCost, sumNum, sumCost int
	Fscan(in, &n, &k)
	nums := make([]int, n)
	cost := make([]int, n)
	for i := range nums {
		Fscan(in, &nums[i], &cost[i])
		totalCost += cost[i]
	}

	q := []vec2365{{}}
	for i, x := range nums {
		sumNum += x
		sumCost += cost[i]

		p := vec2365{-sumNum - k, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}

		p = vec2365{sumCost, p.dot(q[0]) + sumNum*sumCost + k*totalCost}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	Fprint(out, q[len(q)-1].y)
}

//func main() { p2365(bufio.NewReader(os.Stdin), os.Stdout) }
