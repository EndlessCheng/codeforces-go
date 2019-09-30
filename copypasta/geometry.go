package copypasta

import (
	"sort"
)

const eps = 1e-6

type vec struct {
	x, y int64
}

func (a vec) add(b vec) vec    { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec    { return vec{a.x - b.x, a.y - b.y} }
func (a vec) mul(k int64) vec  { return vec{a.x * k, a.y * k} }
func (a vec) div(k int64) vec  { return vec{a.x / k, a.y / k} }
func (a vec) mulVec(b vec) vec { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) dot(b vec) int64  { return a.x*b.x + a.y*b.y }

// + b在a左侧
// - b在a右侧
// 0 ab平行或重合（共基线）
func (a vec) cross(b vec) int64 { return a.x*b.y - a.y*b.x }
func (a vec) reverse() vec      { return a.mul(-1) }

//func (a vec) norm() float64     { return math.Hypot(a.x, a.y) }

// Use equals when x,y is float64, otherwise just use ==
//func (a vec) equals(b vec) bool {
//	return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps
//}

func isOrthogonal(a, b, c vec) bool           { return a.sub(b).dot(c.sub(b)) == 0 }
func isRectangle(a, b, c, d vec) bool         { return isOrthogonal(a, b, c) && isOrthogonal(b, c, d) && isOrthogonal(c, d, a) }
func isRectangleAnyOrder(a, b, c, d vec) bool { return isRectangle(a, b, c, d) || isRectangle(a, b, d, c) || isRectangle(a, c, b, d) }

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
