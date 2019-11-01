package copypasta

import (
	. "fmt"
	"io"
	"math"
	"sort"
)

// 由于浮点默认是 %g，输出时应使用 Fprintf(out, "%.12f", ans)，这样还可以方便测试

const eps = 1e-6

type vec struct {
	x, y int64
}

func read(in io.Reader) vec {
	var x, y int64
	Fscan(in, &x, &y)
	return vec{x, y}
}

func (a vec) less(b vec) bool  { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vec) add(b vec) vec    { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec    { return vec{a.x - b.x, a.y - b.y} }
func (a vec) mul(k int64) vec  { return vec{a.x * k, a.y * k} }
func (a vec) div(k int64) vec  { return vec{a.x / k, a.y / k} }
func (a vec) mulVec(b vec) vec { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) dot(b vec) int64  { return a.x*b.x + a.y*b.y }

// + b在a左侧
// - b在a右侧
// 0 ab平行或重合（共基线）
// up() 后按逆时针排序 sort.Slice(ps, func(i, j int) bool { return ps[i].cross(ps[j]) > 0 })
func (a vec) cross(b vec) int64 { return a.x*b.y - a.y*b.x }
func (a vec) reverse() vec      { return vec{-a.x, -a.y} }
func (a vec) up() vec {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}

func (a vec) len() float64 { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) len2() int64  { return a.x*a.x + a.y*a.y }

// Use equals when x,y is float64, otherwise just use ==
//func (a vec) equals(b vec) bool {
//	return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps
//}

// 如果是线段的话，还需要判断恰好有四个点，并且没有严格交叉（含重合）
func isOrthogonal(a, b, c vec) bool { return a.sub(b).dot(c.sub(b)) == 0 }
func isRectangle(a, b, c, d vec) bool {
	return isOrthogonal(a, b, c) && isOrthogonal(b, c, d) && isOrthogonal(c, d, a)
}
func isRectangleAnyOrder(a, b, c, d vec) bool {
	return isRectangle(a, b, c, d) || isRectangle(a, b, d, c) || isRectangle(a, c, b, d)
}

func geometryCollection() {
	sortVec := func(arr []vec) {
		sort.Slice(arr, func(i, j int) bool {
			v0, v1 := arr[i], arr[j]
			return v0.x < v1.x || v0.x == v1.x && v0.y < v1.y
			//return v0.x+eps < v1.x || math.Abs(v0.x-v1.x) < eps && v0.y+eps < v1.y
		})
	}

	// tests if angle abc is a right angle
	isOrthogonal := func(a, b, c vec) bool {
		return a.sub(b).dot(c.sub(b)) == 0
	}

	isRectangle := func(a, b, c, d vec) bool {
		return isOrthogonal(a, b, c) &&
			isOrthogonal(b, c, d) &&
			isOrthogonal(c, d, a)
	}

	isRectangleAnyOrder := func(a, b, c, d vec) bool {
		return isRectangle(a, b, c, d) ||
			isRectangle(a, b, d, c) ||
			isRectangle(a, c, b, d)
	}

	_ = []interface{}{sortVec, isRectangleAnyOrder}
}

//

type line struct {
	p0, l vec
}

//

type vec3 struct {
	x, y, z int
	idx     int
}

func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.x < pj.x || pi.x == pj.x && (pi.y < pj.y || pi.y == pj.y && pi.z < pj.z)
	})
}
