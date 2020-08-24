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
推荐 https://vlecomte.github.io/cp-geo.pdf

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

dot (dot product，点积):
A·B 可以理解为向量 A 在向量 B 上的投影再乘以 B 的长度

det (determinant，行列式，叉积的模，有向面积):
+ b在a左侧
- b在a右侧
0 ab平行或重合（共基线）
关于有向面积 https://cp-algorithms.com/geometry/oriented-triangle-area.html

1° = (π/180)rad
1rad = (180/π)°
常见的是，弧度为 2*math.Pi*(角度占整个360°的多少)

一些反三角函数的范围
反正弦 -1 ~ 1 => -π/2 ~ π/2
Asin(-1) = -π/2
Asin( 0) = 0
Asin( 1) = π/2
反余弦 1 ~ -1 => 0 ~ π
Acos( 1) = 0
Acos( 0) = π/2
Acos(-1) = π
反正切 三四一二象限 => (-π, π]
( x, y) = Atan2(y,x)
(-1,-1) = -3π/4
( 0,-1) = -π/2
( 1,-1) = -π/4
( 1, 0) = 0
( 1, 1) = π/4
( 0, 1) = π/2
(-1, 1) = 3π/4
(-1, 0) = π
( 0, 0) = 0

todo 二维偏序 https://ac.nowcoder.com/acm/contest/4853/F 题解 https://ac.nowcoder.com/discuss/394080

Pick 定理
https://oi-wiki.org/geometry/pick/
https://cp-algorithms.com/geometry/picks-theorem.html
https://cp-algorithms.com/geometry/lattice-points.html

TIPS: 旋转坐标，适用于基于曼哈顿距离的题目
      顺时针旋转 45° (x,y) -> (x+y,y-x) 记作 (x',y')
      曼哈顿距离 |x1-x2|+|y1-y2| = max(|x1'-x2'|,|y1'-y2'|)
TIPS: 另一种处理曼哈顿距离的方法是枚举 (x2,y2) 在以 (x1,y1) 为原点的哪个象限，分四种情况讨论
      例题 LC1330/双周赛18D https://leetcode-cn.com/problems/reverse-subarray-to-maximize-array-value/

https://oeis.org/A053411 Circle numbers
a(n)= number of points (i,j), i,j integers, contained in a circle of diameter n, centered at the origin

https://oeis.org/A136485 Number of unit square lattice cells enclosed by origin centered circle of diameter n

*/

const eps = 1e-8

/* 二维向量（点）*/
type vec struct{ x, y int64 }

func (a vec) add(b vec) vec     { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec     { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int64   { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int64   { return a.x*b.y - a.y*b.x }
func (a vec) len2() int64       { return a.x*a.x + a.y*a.y }
func (a vec) dis2(b vec) int64  { return a.sub(b).len2() }
func (a vec) len() float64      { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) dis(b vec) float64 { return a.sub(b).len() }
func (a vec) vecF() vecF        { return vecF{float64(a.x), float64(a.y)} }

func (a *vec) adds(b vec) { a.x += b.x; a.y += b.y }
func (a *vec) subs(b vec) { a.x -= b.x; a.y -= b.y }

// 不常用
func (a vec) less(b vec) bool       { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vecF) less(b vecF) bool     { return a.x+eps < b.x || a.x < b.x+eps && a.y+eps < b.y }
func (a vecF) equals(b vecF) bool   { return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps }
func (a vec) onSameLine(b vec) bool { return a.det(b) == 0 }
func (a vec) mul(k int64) vec       { return vec{a.x * k, a.y * k} }
func (a *vec) muls(k int64)         { a.x *= k; a.y *= k }
func (a vecF) div(k float64) vecF   { return vecF{a.x / k, a.y / k} }
func (a *vecF) divs(k float64)      { a.x /= k; a.y /= k }
func (a vec) mulVec(b vec) vec      { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) polarAngle() float64   { return math.Atan2(float64(a.y), float64(a.x)) }
func (a vec) reverse() vec          { return vec{-a.x, -a.y} }
func (a vec) up() vec {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}

func (a vec) angleTo(b vec) float64 {
	v := float64(a.dot(b)) / (a.len() * b.len())
	v = math.Min(math.Max(v, -1), 1)
	return math.Acos(v)
}

// 极角排序
func polarAngleSort(ps []vec) {
	// (-π, π]
	// (-1e9,-1) -> (-1e9, 0)
	// 可以先把每个向量的极角算出来再排序
	sort.Slice(ps, func(i, j int) bool { return ps[i].polarAngle() < ps[j].polarAngle() })

	// EXTRA: 若将所有向量 up() 后极角排序，由于 up() 后向量的范围是 [0°, 180°)，极角排序后共线的向量一定会相邻
	for i := range ps {
		ps[i] = ps[i].up()
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].det(ps[j]) > 0 })
}

// 向量旋转，传入旋转的弧度
func (a vecF) rotate(rad float64) vecF {
	return vecF{a.x*math.Cos(rad) - a.y*math.Sin(rad), a.x*math.Sin(rad) + a.y*math.Cos(rad)}
}

// a 的单位法线，a 不能是零向量
func (a vecF) normal() vecF { l := a.len(); return vecF{-a.y / l, a.x / l} }

// 余弦定理，以及两边及夹角，计算对边长度
func cosineRule(a, b, angle float64) float64 {
	return math.Sqrt(a*a*b*b - 2*a*b*math.Cos(angle))
}
func cosineRuleVec(va, vb vecF, angle float64) float64 {
	return math.Sqrt(va.len2() + vb.len2() - 2*va.len()*vb.len()*math.Cos(angle))
}

// 三角形外心（外接圆圆心，三条边的垂直平分线的交点）
// https://codeforces.com/problemset/problem/1/C
func circumcenter(a, b, c vecF) vecF {
	a1, b1, a2, b2 := b.x-a.x, b.y-a.y, c.x-a.x, c.y-a.y
	c1, c2, d := a1*a1+b1*b1, a2*a2+b2*b2, 2*(a1*b2-a2*b1)
	return vecF{a.x + (c1*b2-c2*b1)/d, a.y + (a1*c2-a2*c1)/d}
}

// EXTRA: 外接圆半径 R
// 最好用 det + 正弦定理，误差小
// https://baike.baidu.com/item/%E5%A4%96%E6%8E%A5%E5%9C%86%E5%8D%8A%E5%BE%84%E5%85%AC%E5%BC%8F
func circumcenterR(a, b, c vecF) float64 {
	ab, ac := b.sub(a), c.sub(a)
	return a.dis(b) * b.dis(c) * c.dis(a) / (2 * math.Abs(ab.det(ac)))
}
func circumcenterR2(a, b, c vecF) float64 {
	ab, ac := b.sub(a), c.sub(a)
	return a.dis2(b) * b.dis2(c) * c.dis2(a) / (4 * ab.det(ac) * ab.det(ac))
}

// 三角形垂心（三条高的交点）
// 欧拉线上的四点中，九点圆圆心到垂心和外心的距离相等，而且重心到外心的距离是重心到垂心距离的一半。注意内心一般不在欧拉线上，除了等腰三角形外
// https://zh.wikipedia.org/wiki/%E6%AD%90%E6%8B%89%E7%B7%9A
// https://baike.baidu.com/item/%E4%B8%89%E8%A7%92%E5%BD%A2%E4%BA%94%E5%BF%83%E5%AE%9A%E5%BE%8B
func orthocenter(a, b, c vecF) vecF {
	return a.add(b).add(c).sub(circumcenter(a, b, c).mul(2))
}

// 三角形内心（三条角平分线的交点）
// 三点坐标按对边长度加权平均
func incenter(a, b, c vecF) vecF {
	bc, ac, ab := b.dis(c), a.dis(c), a.dis(b)
	sum := bc + ac + ab
	return vecF{(bc*a.x + ac*b.x + ab*c.x) / sum, (bc*a.y + ac*b.y + ab*c.y) / sum}
}

/* 二维直线（线段）*/
type line struct{ p1, p2 vec }

// 方向向量 directional vector
func (a line) vec() vec              { return a.p2.sub(a.p1) }
func (a lineF) point(t float64) vecF { return a.p1.add(a.vec().mul(t)) }

// 点 a 是否在线段 l 上（a-p1 与 a-p2 共线且方向相反）
func (a vec) onSeg(l line) bool {
	p1, p2 := l.p1.sub(a), l.p2.sub(a)
	return p1.det(p2) == 0 && p1.dot(p2) <= 0 // 含端点
	//return math.Abs(p1.det(p2)) < eps && p1.dot(p2) < eps
}

// 点 a 是否在射线 o-d 上
func (a vec) onRay(o, d vec) bool {
	a = a.sub(o)
	return d.det(a) == 0 && d.dot(a) >= 0 // 含端点
}

// 直线 a b 交点
// NOTE: 若输入均为有理数，则输出也为有理数
func (a lineF) intersection(b lineF) vecF {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	t := vb.det(u) / va.det(vb) // a b 不能平行，即 va.det(vb) != 0
	return a.point(t)
}

// 求射线 a b 交点，返回各自到首个交点所需的时间（射线速度由 .vec().len() 决定）
// 无交点返回 -1
// 交点为 a.point(ta) 或 b.point(tb)
// 若题目给了方向向量和速度：https://codeforces.com/problemset/problem/1359/F
func (a lineF) rayIntersection(b lineF) (ta, tb float64) {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	if d := va.det(vb); d != 0 {
		d1, d2 := vb.det(u), va.det(u)
		if d > 0 && d1 >= 0 && d2 >= 0 || d < 0 && d1 <= 0 && d2 <= 0 {
			return d1 / d, d2 / d
		}
		return -1, -1
	}
	if u.det(va) != 0 { // 平行但未共基线
		return -1, -1
	}
	if l := u.len(); va.dot(vb) > 0 { // 同向
		if u.dot(vb) >= 0 {
			return 0, l / vb.len()
		}
		return l / va.len(), 0
	} else { // 异向
		if u.dot(vb) >= 0 {
			t := l / (va.len() + vb.len())
			return t, t
		}
		return -1, -1
	}
}

// 点 a 到直线 l 的距离
// 若不取绝对值得到的是有向距离
func (a vecF) disToLine(l lineF) float64 {
	v, p1a := l.vec(), a.sub(l.p1)
	return math.Abs(v.det(p1a)) / v.len()
}

// 点 a 到线段 l 的距离
func (a vecF) disToSeg(l lineF) float64 {
	if l.p1 == l.p2 {
		return a.sub(l.p1).len()
	}
	v, p1a, p2a := l.vec(), a.sub(l.p1), a.sub(l.p2)
	if v.dot(p1a) < eps {
		return p1a.len()
	}
	if v.dot(p2a) > -eps {
		return p2a.len()
	}
	return math.Abs(v.det(p1a)) / v.len()
}

// 点 a 到线段 l 的距离 <= r
func (a vec) withinRange(l line, r int64) bool {
	v, p1a, p2a := l.vec(), a.sub(l.p1), a.sub(l.p2)
	if v.dot(p1a) <= 0 {
		return p1a.len2() <= r*r
	}
	if v.dot(p2a) >= 0 {
		return p2a.len2() <= r*r
	}
	return v.det(p1a)*v.det(p1a) <= v.len2()*r*r
}

// 点 a 在直线 l 上的投影
func (a vecF) projection(l lineF) vecF {
	v, p1a := l.vec(), a.sub(l.p1)
	t := v.dot(p1a) / v.len2()
	return l.p1.add(v.mul(t))
}

// 线段规范相交
// CCW (counterclockwise)
func (a lineF) segProperIntersection(b lineF) bool {
	sign := func(x float64) int {
		if x < -eps {
			return -1
		}
		if x < eps {
			return 0
		}
		return 1
	}
	va, vb := a.vec(), b.vec()
	d1, d2 := va.det(b.p1.sub(a.p1)), va.det(b.p2.sub(a.p1))
	d3, d4 := vb.det(a.p1.sub(b.p1)), vb.det(a.p2.sub(b.p1))
	return sign(d1)*sign(d2) < 0 && sign(d3)*sign(d4) < 0
}

// 过点 a 的垂直于 l 的直线
func (a vec) perpendicular(l line) line { return line{a, a.add(vec{l.p1.y - l.p2.y, l.p2.x - l.p1.x})} }

/* 圆 */
type circle struct {
	vec
	r int64
}

// 圆心角对应的点
func (o circle) point(rad float64) vecF {
	return vecF{float64(o.x) + float64(o.r)*math.Cos(rad), float64(o.y) + float64(o.r)*math.Sin(rad)}
}
func (o circleF) point(rad float64) vecF {
	return vecF{o.x + o.r*math.Cos(rad), o.y + o.r*math.Sin(rad)}
}

// 给定半径和一条有向的弦，求该弦右侧的圆心（即 ao 在 ab 右侧）
func getCircleCenter(a, b vec, r int64) vecF {
	disAB2 := b.sub(a).len2()
	//if disAB2 > 4*r*r {
	//	continue
	//}
	midX, midY := float64(a.x+b.x)/2, float64(a.y+b.y)/2
	d := math.Sqrt(float64(r*r) - float64(disAB2/4))
	angle := math.Atan2(float64(b.y-a.y), float64(b.x-a.x))
	return vecF{midX + d*math.Sin(angle), midY - d*math.Cos(angle)}
}

// 直线与圆的交点
func (o circleF) intersectionLine(l lineF) (ps []vecF, t1, t2 float64) {
	v := l.vec()
	a, b, c, d := v.x, l.p1.x-o.x, v.y, l.p1.y-o.y
	e, f, g := a*a+c*c, 2*(a*b+c*d), b*b+d*d-o.r*o.r
	switch delta := f*f - 4*e*g; {
	case delta < -eps: // 相离
		return
	case delta < eps: // 相切
		t := -f / (2 * e)
		return []vecF{l.point(t)}, t, t
	default: // 相交
		t1 = (-f - math.Sqrt(delta)) / (2 * e)
		t2 = (-f + math.Sqrt(delta)) / (2 * e)
		ps = []vecF{l.point(t1), l.point(t2)}
		return
	}
}

// 两圆交点
func (o circle) intersectionCircle(b circle) (ps []vecF, normal bool) {
	a := o
	if a.r < b.r {
		a, b = b, a
	}
	ab := b.sub(a.vec)
	dab2 := ab.len2()
	diffR, sumR := a.r-b.r, a.r+b.r
	if dab2 == 0 {
		if diffR == 0 {
			return
		} // 重合
		return nil, true
	}
	normal = true
	if sumR*sumR < dab2 || diffR*diffR > dab2 {
		return
	}

	angleAB := ab.polarAngle()
	angleDelta := math.Acos(float64(sumR*diffR+dab2) / (float64(2*a.r) * ab.len())) // 余弦定理
	p1, p2 := a.point(angleAB-angleDelta), a.point(angleAB+angleDelta)
	ps = append(ps, p1)
	if math.Abs(angleDelta) > eps {
		ps = append(ps, p2)
	}
	return
}

// 圆的面积并 - 两圆的特殊情形
// todo https://codeforces.com/contest/600/problem/D

// 圆的面积并
// todo

// 点到圆的切线，返回向量即可
func (o circle) tangents(p vec) (ls []vecF) {
	po := o.sub(p)
	d2 := po.len2()
	if d2 < o.r*o.r {
		return
	}
	if d2 == o.r*o.r {
		return []vecF{po.rotate(math.Pi / 2)}
	} // 圆上一点的切线
	ang := math.Asin(float64(o.r) / po.len())
	return []vecF{po.rotate(-ang), po.rotate(ang)}
}

// 两圆公切线
// 返回每条切线在圆 o 和圆 ob 的切点
// NOTE: 下面的代码是基于 int64 的，没有判断 eps
func (o circle) tangents2(b circle) (ls []lineF, hasInf bool) {
	a := o
	if a.r < b.r {
		a, b = b, a
	}
	ab := b.sub(a.vec)
	dab2 := ab.len2()
	diffR, sumR := a.r-b.r, a.r+b.r
	if dab2 == 0 && diffR == 0 { // 完全重合
		return nil, true
	}
	if dab2 < diffR*diffR { // 内含
		return
	}
	angleAB := ab.polarAngle()
	if dab2 == diffR*diffR { // 内切
		return []lineF{{a.point(angleAB), b.point(angleAB)}}, false
	}

	dab := ab.len()
	ang := math.Acos(float64(diffR) / dab)
	ls = append(ls, lineF{a.point(angleAB + ang), b.point(angleAB + ang)}) // 两条外公切线
	ls = append(ls, lineF{a.point(angleAB - ang), b.point(angleAB - ang)})
	if dab2 == sumR*sumR { // 一条内公切线
		ls = append(ls, lineF{a.point(angleAB), b.point(angleAB + math.Pi)})
	} else if dab2 > sumR*sumR { // 两条内公切线
		ang = math.Acos(float64(sumR) / dab)
		ls = append(ls, lineF{a.point(angleAB + ang), b.point(angleAB + ang + math.Pi)})
		ls = append(ls, lineF{a.point(angleAB - ang), b.point(angleAB - ang + math.Pi)})
	}
	return
}

// 最小圆覆盖
func minCircleCoverAllPoints(ps []vec) circle {
	// todo
	panic("todo")
}

// 求一固定半径的圆最多能覆盖多少个点（圆边上也算覆盖） len(ps)>0 && r>0
// Angular Sweep 算法 O(n^2logn)
// https://www.geeksforgeeks.org/angular-sweep-maximum-points-can-enclosed-circle-given-radius/
// LC1453/周赛189D https://leetcode-cn.com/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/solution/python3-angular-sweepsuan-fa-by-lih/
func maxCoveredPoints(ps []vec, r int64) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const eps = 1e-8
	type event struct {
		angle float64
		delta int
	}

	n := len(ps)
	ans := 1
	for i, a := range ps {
		events := make([]event, 0, 2*n-2)
		for j, b := range ps {
			if j == i {
				continue
			}
			ab := b.sub(a)
			if ab.len2() > 4*r*r {
				continue
			}
			at := math.Atan2(float64(ab.y), float64(ab.x))
			ac := math.Acos(ab.len() / float64(2*r))
			events = append(events, event{at - ac, 1}, event{at + ac, -1})
		}
		sort.Slice(events, func(i, j int) bool { a, b := events[i], events[j]; return a.angle+eps < b.angle || a.angle < b.angle+eps && a.delta > b.delta })
		mx, cnt := 0, 1 // 1 指当前固定的点 a
		for _, e := range events {
			cnt += e.delta
			mx = max(mx, cnt)
		}
		ans = max(ans, mx)
	}
	return ans
}

// 反演变换
// https://en.wikipedia.org/wiki/Inversive_geometry
// TODO: https://oi-wiki.org/geometry/inverse/

// 三角剖分
// todo https://oi-wiki.org/geometry/triangulation/
//      https://cp-algorithms.com/geometry/delaunay.html

func vec2Collection() {
	readVec := func(in io.Reader) vec {
		var x, y int64
		Fscan(in, &x, &y)
		return vec{x, y}
	}

	leftMostVec := func(p0 vec, ps []vec) (idx int) {
		for i, p := range ps {
			if ps[idx].sub(p0).det(p.sub(p0)) > 0 {
				idx = i
			}
		}
		return
	}
	rightMostVec := func(p0 vec, ps []vec) (idx int) {
		for i, p := range ps {
			if ps[idx].sub(p0).det(p.sub(p0)) < 0 {
				idx = i
			}
		}
		return
	}

	// TODO: 扫描线：线段求交 O(nlogn)
	// N 条线段求交的扫描线算法 http://johnhany.net/2013/11/sweep-algorithm-for-segments-intersection/
	// https://codeforces.com/problemset/problem/1359/F
	// 平面扫描思想在 ACM 竞赛中的应用 http://openinx.github.io/2013/01/01/plane-sweep-thinking/

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

	// 平面最近点对
	// 调用前 ps 必须按照 x 坐标排序：
	// sort.Slice(ps, func(i, j int) bool { return ps[i].x < ps[j].x })
	// 保证没有重复的点
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/ClosestPair.java.html
	// 模板题 https://www.luogu.com.cn/problem/P1429
	var closestPair func([]vec) float64
	closestPair = func(ps []vec) float64 {
		n := len(ps) // assert n >= 2
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

	// 多边形面积
	// https://cp-algorithms.com/geometry/area-of-simple-polygon.html
	polygonArea := func(ps []vec) float64 {
		n := len(ps)
		area := 0.0
		for i := 1; i < n-1; i++ {
			area += float64(ps[i].sub(ps[0]).det(ps[i+1].sub(ps[0])))
		}
		return area / 2
	}

	// 求凸包 葛立恒扫描法 Graham's scan
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GrahamScan.java.html
	// NOTE: 坐标值范围不超过 M 的凸多边形的顶点数为 O(√M) 个
	// 模板题 https://www.luogu.com.cn/problem/P2742
	// 限制区间长度的区间最大均值问题 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A
	convexHull := func(ps []vec) []vec {
		n := len(ps)
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
		ch := make([]vec, 0, 2*n)
		for _, p := range ps {
			for {
				sz := len(ch)
				if sz <= 1 || ch[sz-1].sub(ch[sz-2]).det(p.sub(ch[sz-1])) > 0 {
					break
				}
				ch = ch[:sz-1]
			}
			ch = append(ch, p)
		}
		downSize := len(ch)
		for i := n - 2; i >= 0; i-- {
			p := ps[i]
			for {
				sz := len(ch)
				if sz <= downSize || ch[sz-1].sub(ch[sz-2]).det(p.sub(ch[sz-1])) > 0 {
					break
				}
				ch = ch[:sz-1]
			}
			ch = append(ch, p)
		}
		return ch[:len(ch)-1]
	}

	// 旋转卡壳求最远点对（凸包直径） Rotating calipers
	// https://en.wikipedia.org/wiki/Rotating_calipers
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FarthestPair.java.html
	// 模板题 https://www.luogu.com.cn/problem/P1452
	rotatingCalipers := func(ps []vec) (p1, p2 vec) {
		qs := convexHull(ps)
		n := len(qs)
		if n == 2 {
			return qs[0], qs[1]
		}
		i, j := 0, 0 // 对踵点对（左下和右上）
		for k, p := range qs {
			if !qs[i].less(p) {
				i = k
			}
			if qs[j].less(p) {
				j = k
			}
		}
		maxDis2 := int64(0)
		i0, j0 := i, j
		for i != j0 || j != i0 {
			if d2 := qs[i].sub(qs[j]).len2(); d2 > maxDis2 {
				maxDis2 = d2
				p1, p2 = qs[i], qs[j]
			}
			// 判断先转到边 i-(i+1) 的法线方向还是边 j-(j+1) 的法线方向
			if qs[(i+1)%n].sub(qs[i]).det(qs[(j+1)%n].sub(qs[j])) < 0 {
				i = (i + 1) % n
			} else {
				j = (j + 1) % n
			}
		}
		return
	}

	// todo 点集的最大四边形
	// https://www.luogu.com.cn/problem/P4166
	// https://codeforces.com/contest/340/problem/B

	// 凸包周长
	convexHullLength := func(ps []vec) (res float64) {
		qs := convexHull(ps)
		for i := 1; i < len(qs); i++ {
			res += qs[i].sub(qs[i-1]).len()
		}
		return
	}

	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	// 点 p 是否在三角形 △abc 内
	inTriangle := func(a, b, c, p vec) bool {
		pa, pb, pc := a.sub(p), b.sub(p), c.sub(p)
		return abs(b.sub(a).det(c.sub(a))) == abs(pa.det(pb))+abs(pb.det(pc))+abs(pc.det(pa))
	}

	// 判断点 p 是否在凸多边形 ps 内部 O(logN)
	// ps 逆时针顺序
	// https://www.cnblogs.com/yym2013/p/3673616.html
	// https://cp-algorithms.com/geometry/point-in-convex-polygon.html
	// 其他 O(n) 方法 https://blog.csdn.net/WilliamSun0122/article/details/77994526
	inPolygon := func(ps []vec, p vec) bool {
		n := len(ps)
		p0 := p.sub(ps[0])
		// 外：最右射线的右侧或最左射线的左侧
		if ps[1].sub(ps[0]).det(p0) < 0 || ps[n-1].sub(ps[0]).det(p0) > 0 {
			return false
		}
		// p0 是否在右侧或重合
		i := sort.Search(n, func(mid int) bool { return ps[mid].sub(ps[0]).det(p0) <= 0 })
		// 是否在边左侧或与边重合
		return ps[i].sub(ps[i-1]).det(p.sub(ps[i-1])) >= 0
	}

	// todo 判断点 p 是否在多边形 ps 内部
	// 光线投射算法 Ray casting algorithm

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

	// TODO 矩形面积并
	// 扫描线算法
	// 模板题 https://www.luogu.com.cn/problem/P5490

	// TODO 半平面交
	// 推荐 训练指南 的相关章节
	// https://oi-wiki.org/geometry/half-plane/
	// https://zhuanlan.zhihu.com/p/83499723
	// 模板题 https://www.luogu.com.cn/problem/P4196

	_ = []interface{}{
		readVec, leftMostVec, rightMostVec, readPolygon, polygonArea, rotatingCalipers, convexHullLength, inTriangle, inPolygon,
		isRectangleAnyOrder, minAreaRect,
	}
}

/* 三维向量（点）*/
type vec3 struct{ x, y, z int64 }

func (a vec3) less(b vec3) bool {
	return a.x < b.x || a.x == b.x && (a.y < b.y || a.y == b.y && a.z < b.z)
}

/* 三维直线（线段）*/
type line3 struct{ p1, p2 vec3 }

func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && (a.y < b.y || a.y == b.y && a.z < b.z) })

	// 三维凸包
	// todo 模板题 https://www.luogu.com.cn/problem/P4724
}

// 下面这些仅作为占位符表示，实际使用的时候复制上面的模板，类型改成 float64 同时 vecF 替换成 vec 等
type vecF struct{ x, y float64 }
type lineF struct{ p1, p2 vecF }
type vec3F struct{ x, y, z int64 }
type line3F struct{ p1, p2 vec3F }
type circleF struct {
	vecF
	r float64
}

func (vecF) add(vecF) (_ vecF)      { return }
func (vecF) sub(vecF) (_ vecF)      { return }
func (vecF) mul(float64) (_ vecF)   { return }
func (vecF) dot(vecF) (_ float64)   { return }
func (vecF) det(vecF) (_ float64)   { return }
func (vecF) len() (_ float64)       { return }
func (vecF) len2() (_ float64)      { return }
func (vecF) dis(vecF) (_ float64)   { return }
func (vecF) dis2(vecF) (_ float64)  { return }
func (vec) rotate(float64) (_ vecF) { return }
func (lineF) vec() (_ vecF)         { return }
