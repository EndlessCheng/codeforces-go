package copypasta

import (
	. "fmt"
	"io"
	"math"
	"sort"
)

// https://oi-wiki.org/geometry/2d/
// https://oi-wiki.org/geometry/3d/

// 由于浮点默认是 %g，输出时应使用 Fprintf(out, "%.16f", ans)，这样还可以方便测试

// NOTE: always add `eps` when do printf rounding
// Sprintf("%.1f", 0.25) == "0.2"
// Sprintf("%.1f", 0.25+eps) == "0.3"

// NOTE: 比较两个大浮点数（相减误差大于 eps）
// a > b    a > (1+eps)*b
// a >= b   a > (1-eps)*b
// a == b   math.Abs(a/b-1) < eps 或 (1-eps)*b < a && a < (1+eps)*b

const eps = 1e-8

type vec struct {
	x, y int64
}

func read(in io.Reader) vec {
	var x, y int64
	Fscan(in, &x, &y)
	return vec{x, y}
}

//func (a vec) equals(b vec) bool { return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps }
//func (a vec) less(b vec) bool   { return a.x+eps < b.x || math.Abs(a.x-b.x) < eps && a.y+eps < b.y }
func (a vec) less(b vec) bool { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vec) add(b vec) vec   { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec   { return vec{a.x - b.x, a.y - b.y} }
func (a vec) mul(k int64) vec { return vec{a.x * k, a.y * k} }
func (a vec) div(k int64) vec { return vec{a.x / k, a.y / k} }
func (a vec) len() float64    { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) len2() int64     { return a.x*a.x + a.y*a.y }
func (a vec) dot(b vec) int64 { return a.x*b.x + a.y*b.y }
func (a vec) reverse() vec    { return vec{-a.x, -a.y} }
func (a vec) up() vec {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}
func (a vec) det(b vec) int64 { return a.x*b.y - a.y*b.x }

// + b在a左侧
// - b在a右侧
// 0 ab平行或重合（共基线）
// up() 后按逆时针排序: sort.Slice(ps, func(i, j int) bool { return ps[i].det(ps[j]) > 0 })
// 排序后保证共线的向量是相邻的（因为范围是 [0, 180) ）

func (a vec) mulVec(b vec) vec    { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) angle(b vec) float64 { return math.Acos(float64(a.dot(b)) / (a.len() * b.len())) }

//func (a vec) rotate(rad float64) { // 弧度
//	return vec{a.x*math.Cos(rad) - a.y*math.Sin(rad), a.x*math.Sin(rad) + a.y*math.Cos(rad)}
//}

// a 的单位法线
//func (a vec) normal() vec {
//	l := a.len()
//	return vec{-a.y / l, a.x / l}
//}

type line struct {
	// 创建 line 时，要求 p1 != p2
	p1, p2 vec
}

// 点 a 在直线 l 上的投影
//func (a vec) projection(l line) vec {
//	v := l.p2.sub(l.p1)
//	t := v.dot(a.sub(l.p1)) / v.len()
//	return l.p1.add(v.mul(t))
//}

// 过点 a 的垂直于 l 的直线
func (a vec) perpendicular(l line) line {
	return line{a, a.add(vec{l.p1.y - l.p2.y, l.p2.x - l.p1.x})}
}

// 直线 a b 交点
// 必须用 float64
func (a line) intersection(b line) vec {
	va, vb, u := a.p2.sub(a.p1), b.p2.sub(b.p1), a.p1.sub(b.p1)
	t := vb.det(u) / va.det(vb)
	return a.p1.add(va.mul(t))
}

// 点 a 到直线 l 的距离
func (a vec) disToLine(l line) float64 {
	v, u := l.p2.sub(l.p1), a.sub(l.p1)
	return math.Abs(float64(v.det(u))) / v.len()
}

// 点 a 到线段 l 的距离
//func (a vec) disToSeg(l line) float64 {
//	if l.p1 == l.p2 {
//		return a.sub(l.p1).len()
//	}
//	v, v1, v2 := l.p2.sub(l.p1), a.sub(l.p1), a.sub(l.p2)
//	if v.dot(v1) < -eps {
//		return v1.len()
//	}
//	if v.dot(v2) > eps {
//		return v2.len()
//	}
//	return math.Abs(v.det(v1)) / v.len()
//}

// 点 a 是否在线段 p1-p2 上（a-p1 与 a-p2 共线且方向相反）
func (a vec) onSeg(l line) bool {
	p1 := l.p1.sub(a)
	p2 := l.p2.sub(a)
	return p1.det(p2) == 0 && p1.dot(p2) <= 0 // 含端点
	// 如果 a 已经在 l 上了直接 return p1.dot(p2) < eps
	//return math.Abs(p1.det(p2)) < eps && p1.dot(p2) < eps
}

// 线段规范相交
// CCW (Counter Clock Wise) ?
func (a line) segProperIntersection(b line) bool {
	// TODO
}

type circle struct {
	// TODO
}

// https://en.wikipedia.org/wiki/Inversive_geometry
// https://oi-wiki.org/geometry/inverse/

func vec2Collection() {
	merge := func(a, b []vec) []vec {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]vec, 0, n+m)
		for {
			if i == n {
				return append(res, b[j:]...)
			}
			if j == m {
				return append(res, a[i:]...)
			}
			if a[i].y < b[j].y {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}

	// 最近点对
	// 调用前 ps 必须按照 x 坐标排序
	// sort.Slice(ps, func(i, j int) bool { return ps[i].x < ps[j].x })
	var closestPair func([]vec) float64
	closestPair = func(ps []vec) float64 {
		n := len(ps)
		if n <= 1 {
			return math.MaxFloat64
		}
		m := n >> 1
		x := ps[m].x
		d := math.Min(closestPair(ps[:m]), closestPair(ps[m:]))
		for i, p := range merge(ps[:m], ps[m:]) {
			ps[i] = p
		}
		checkPs := []vec{}
		for _, pi := range ps {
			if math.Abs(float64(pi.x-x)) > d+eps {
				continue
			}
			for j := len(checkPs) - 1; j >= 0; j-- {
				pj := checkPs[j]
				dy := float64(pi.y - pj.y)
				if dy >= d {
					break
				}
				dx := float64(pi.x - pj.x)
				d = math.Min(d, math.Hypot(dx, dy))
			}
			checkPs = append(checkPs, pi)
		}
		return d
	}

	// 读入多边形
	// 输入的点必须按顺时针或逆时针顺序输入
	readPolygon := func(in io.Reader, n int) []line {
		ps := make([]vec, n)
		for i := range ps {
			Fscan(in, &ps[i].x, &ps[i].y)
		}
		ls := make([]line, n)
		for i := 0; i < n-1; i++ {
			ls[i] = line{ps[i], ps[i+1]}
		}
		ls[n-1] = line{ps[n-1], ps[0]}
		return ls
	}

	// 凸包
	// qs[0] == qs[-1]
	convexHull := func(ps []vec) []vec {
		n := len(ps)
		sort.Slice(ps, func(i, j int) bool {
			a, b := ps[i], ps[j]
			return a.x < b.x || a.x == b.x && a.y < b.y
		})
		qs := make([]vec, 2*n)
		for _, pi := range ps {
			for {
				sz := len(qs)
				if sz <= 1 || qs[sz-1].sub(qs[sz-2]).det(pi.sub(qs[sz-1])) > 0 {
					break
				}
				qs = qs[:sz-1]
			}
			qs = append(qs, pi)
		}
		downSize := len(qs)
		for i := n - 2; i >= 0; i-- {
			pi := ps[i]
			for {
				sz := len(qs)
				if sz <= downSize || qs[sz-1].sub(qs[sz-2]).det(pi.sub(qs[sz-1])) > 0 {
					break
				}
				qs = qs[:sz-1]
			}
			qs = append(qs, pi)
		}
		return qs
	}

	// 凸包周长
	convexHullLength := func(ps []vec) (res float64) {
		qs := convexHull(ps)
		for i := 1; i < len(qs); i++ {
			res += qs[i].sub(qs[i-1]).len()
		}
		return
	}

	// 如果是线段的话，还需要判断恰好有四个点，并且没有严格交叉（含重合）
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

	// vs 中不能有重复的点
	minAreaRect := func(vs []vec) float64 {
		mp := map[vec]bool{}
		for _, v := range vs {
			mp[v] = true
		}
		ans := int64(math.MaxInt64)
		for i, va := range vs {
			for j, vb := range vs {
				if j == i {
					continue
				}
				for k, vc := range vs {
					if k == i || k == j {
						continue
					}
					if isOrthogonal(va, vb, vc) && mp[va.add(vc.sub(vb))] {
						if area := va.sub(vb).len2() * vc.sub(vb).len2(); area < ans {
							ans = area
						}
					}
				}
			}
		}
		if ans == math.MaxInt64 {
			ans = 0
		}
		return math.Sqrt(float64(ans))
	}

	_ = []interface{}{readPolygon, convexHullLength, isRectangleAnyOrder, minAreaRect}
}

//

type vec3 struct {
	x, y, z int64
}

func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.x < pj.x || pi.x == pj.x && (pi.y < pj.y || pi.y == pj.y && pi.z < pj.z)
	})
}
