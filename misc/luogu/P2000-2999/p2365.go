package P2000_2999

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type vec365 struct{ x, y int }

func (a vec365) sub(b vec365) vec365 { return vec365{a.x - b.x, a.y - b.y} }
func (a vec365) det(b vec365) int    { return a.x*b.y - a.y*b.x }
func (a vec365) dot(b vec365) int    { return a.x*b.x + a.y*b.y }

func p2365(in io.Reader, out io.Writer) {
	var n, k, totalCost, sumNum, sumCost int
	Fscan(in, &n, &k)
	nums := make([]int, n)
	cost := make([]int, n)
	for i := range nums {
		Fscan(in, &nums[i], &cost[i])
		totalCost += cost[i]
	}

	q := []vec365{{}}
	for i, x := range nums {
		sumNum += x
		sumCost += cost[i]

		p := vec365{-sumNum - k, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}

		p = vec365{sumCost, p.dot(q[0]) + sumNum*sumCost + k*totalCost}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	Fprint(out, q[len(q)-1].y)
}

//func main() { p2365(bufio.NewReader(os.Stdin), os.Stdout) }
