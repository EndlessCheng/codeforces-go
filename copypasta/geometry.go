package copypasta

import (
	. "fmt"
	"io"
	"math"
	"sort"
)

/*
https://oi-wiki.org/geometry/2d/
https://oi-wiki.org/geometry/3d/

由于浮点默认是 %g，输出时应使用 Fprintf(out, "%.16f", ans)，这样还可以方便测试

always add `eps` when do printf rounding:
Sprintf("%.1f", 0.25) == "0.2"
Sprintf("%.1f", 0.25+eps) == "0.3"

比较浮点数：
a < b    a+eps < b
a <= b   a-eps < b
a == b   math.Abs(a-b) < eps

比较大浮点数（因为是即使 a 和 b 相近，a-b 的误差也可能大于 eps，见 CF1059D）：
a < b    a*(1+eps) < b
a <= b   a*(1-eps) < b
a == b   a*(1-eps) < b && b < a*(1+eps)

dot (dot product，点积，A·B 可以理解为向量 A 在向量 B 上的投影再乘以 B 的长度)

det (determinant，行列式，叉积的模，有向面积):
+ b在a左侧
- b在a右侧
0 ab平行或重合（共基线）

1° = (π/180)rad
1rad = (180/π)°
*/

const eps = 1e-8

type vec struct{ x, y int64 }

func (a vec) add(b vec) vec   { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec   { return vec{a.x - b.x, a.y - b.y} }
func (a vec) mul(k int64) vec { return vec{a.x * k, a.y * k} }
func (a vec) len() float64    { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) len2() int64     { return a.x*a.x + a.y*a.y }
func (a vec) dot(b vec) int64 { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int64 { return a.x*b.y - a.y*b.x }

//func (a vec) equals(b vec) bool { return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps }
//func (a vec) less(b vec) bool   { return a.x+eps < b.x || math.Abs(a.x-b.x) < eps && a.y+eps < b.y }
func (a vec) less(b vec) bool       { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vec) div(k int64) vec       { return vec{a.x / k, a.y / k} }
func (a vec) mulVec(b vec) vec      { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) angleTo(b vec) float64 { return math.Acos(float64(a.dot(b)) / (a.len() * b.len())) }
func (a vec) polarAngle() float64   { return math.Atan2(float64(a.y), float64(a.x)) }
func (a vec) reverse() vec          { return vec{-a.x, -a.y} }
func (a vec) up() vec {
	// 所有向量 up() 后按逆时针排序:
	// sort.Slice(ps, func(i, j int) bool { return ps[i].det(ps[j]) > 0 })
	// 由于 up() 后所有向量的范围是 [0°, 180°)，在排序后共线的向量一定会相邻
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}

// 向量旋转，传入旋转的弧度
//func (a vec) rotate(rad float64) {
//	return vec{a.x*math.Cos(rad) - a.y*math.Sin(rad), a.x*math.Sin(rad) + a.y*math.Cos(rad)}
//}

// a 的单位法线，a 不能是零向量
//func (a vec) normal() vec {
//	l := a.len()
//	return vec{-a.y / l, a.x / l}
//}

//

type line struct{ p1, p2 vec }

// 方向向量 directional vector
func (a line) vec() vec { return a.p2.sub(a.p1) }

// 直线 a b 交点
func (a line) intersection(b line) vec {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	t := vb.det(u) / va.det(vb) // a b 不能平行，即 va.det(vb) != 0
	return a.p1.add(va.mul(t))
}

// 点 a 到直线 l 的距离
// 若不取绝对值得到的是有向距离
func (a vec) disToLine(l line) float64 {
	v, p1a := l.vec(), a.sub(l.p1)
	return math.Abs(float64(v.det(p1a))) / v.len()
}

// 点 a 到线段 l 的距离
func (a vec) disToSeg(l line) float64 {
	if l.p1 == l.p2 {
		return a.sub(l.p1).len()
	}
	v, p1a, p2a := l.vec(), a.sub(l.p1), a.sub(l.p2)
	if float64(v.dot(p1a)) < -eps {
		return p1a.len()
	}
	if float64(v.dot(p2a)) > eps {
		return p2a.len()
	}
	return math.Abs(float64(v.det(p1a))) / v.len()
}

// 点 a 在直线 l 上的投影
func (a vec) projection(l line) vec {
	v, p1a := l.vec(), a.sub(l.p1)
	t := v.dot(p1a) / v.len2()
	return l.p1.add(v.mul(t))
}

// 线段规范相交
// CCW (counterclockwise)
func (a line) segProperIntersection(b line) bool {
	sign := func(x float64) int {
		if math.Abs(x) < eps {
			return 0
		}
		if x < 0 {
			return -1
		}
		return 1
	}
	_ = sign
	va, vb := a.vec(), b.vec()
	d1, d2 := va.det(b.p1.sub(a.p1)), va.det(b.p2.sub(a.p1))
	d3, d4 := vb.det(a.p1.sub(b.p1)), vb.det(a.p2.sub(b.p1))
	return d1*d2 < 0 && d3*d4 < 0
	//return sign(float64(d1))*sign(float64(d2)) < 0 && sign(float64(d3))*sign(float64(d4)) < 0
}

// 点 a 是否在线段 l 上（a-p1 与 a-p2 共线且方向相反）
func (a vec) onSeg(l line) bool {
	p1, p2 := l.p1.sub(a), l.p2.sub(a)
	return p1.det(p2) == 0 && p1.dot(p2) <= 0 // 含端点
	//return math.Abs(p1.det(p2)) < eps && p1.dot(p2) < eps
}

// 过点 a 的垂直于 l 的直线
func (a vec) perpendicular(l line) line {
	return line{a, a.add(vec{l.p1.y - l.p2.y, l.p2.x - l.p1.x})}
}

//

type circle struct {
	// TODO implement
}

// TODO 直线与圆的交点
// TODO 两个圆的公切线

// 反演变换
// https://en.wikipedia.org/wiki/Inversive_geometry
// TODO: https://oi-wiki.org/geometry/inverse/

func vec2Collection() {
	readVec := func() vec {
		var x, y int64
		//Fscan(in, &x, &y)
		return vec{x, y}
	}

	// TODO: 扫描线：线段求交 O(nlogn)

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
		copy(ps, merge(ps[:m], ps[m:]))
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

	polygonArea := func(ps []vec) float64 {
		n := len(ps)
		area := 0.0
		for i := 1; i < n-1; i++ {
			area += float64(ps[i].sub(ps[0]).det(ps[i+1].sub(ps[0])))
		}
		return area / 2
	}

	// 求凸包 葛立恒扫描法 Graham's scan
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

	// TODO: 旋转卡壳求最远点对

	// 凸包周长
	convexHullLength := func(ps []vec) (res float64) {
		qs := convexHull(ps)
		for i := 1; i < len(qs); i++ {
			res += qs[i].sub(qs[i-1]).len()
		}
		return
	}

	// 判断 ∠abc 是否为直角
	// 如果是线段的话，还需要判断恰好有四个点，并且没有严格交叉（含重合）
	isOrthogonal := func(a, b, c vec) bool { return a.sub(b).dot(c.sub(b)) == 0 }

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

	// 求点集中的最小矩形
	// vs 中不能有重复的点
	minAreaRect := func(vs []vec) (minArea float64) {
		mp := map[vec]bool{}
		for _, v := range vs {
			mp[v] = true
		}
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
						// 要是中间爆了的话就各自 float64 再相乘
						if area := float64(va.sub(vb).len2() * vc.sub(vb).len2()); minArea == 0 || area < minArea {
							minArea = area
						}
					}
				}
			}
		}
		return math.Sqrt(minArea)
	}

	_ = []interface{}{readVec, readPolygon, polygonArea, convexHullLength, isRectangleAnyOrder, minAreaRect}
}

//

type vec3 struct{ x, y, z int64 }

func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.x < pj.x || pi.x == pj.x && (pi.y < pj.y || pi.y == pj.y && pi.z < pj.z)
	})
}
