package copypasta

import (
	"cmp"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"slices"
	"sort"
)

/*
https://oi-wiki.org/geometry/2d/
https://oi-wiki.org/geometry/3d/
推荐 https://vlecomte.github.io/cp-geo.pdf
https://www.cnblogs.com/Xing-Ling/p/12102489.html
jiangly https://ac.nowcoder.com/acm/contest/view-submission?submissionId=62808640
kuangbin https://kuangbin.github.io/2019/04/28/20190428/

由于浮点数默认是以 %g 输出，正确写法是 Fprintf(out, "%.15f", ans)，这样还可以方便测试

NOTE: always add `eps` when do printf rounding
Sprintf("%.1f", 0.25) == "0.2"
Sprintf("%.1f", 0.25+eps) == "0.3"

比较小浮点数，采用绝对误差
a < b    a+eps < b
a <= b   a-eps < b
a == b   math.Abs(a-b) < eps

比较大浮点数，采用相对误差（因为是即使 a 和 b 相近，a-b 的误差也可能大于 eps，见 CF1059D）
a < b    a*(1+eps) < b
a <= b   a*(1-eps) < b
a == b   a*(1-eps) < b && b < a*(1+eps)

计算几何第一课：如何避免卡精度？如何设置合理的精度？
避免不同数量级的浮点数的加减可以减小误差
见下面这两份代码的区别
https://codeforces.com/problemset/submission/621/116068024
https://codeforces.com/problemset/submission/621/116068186
调精度 https://codeforces.com/contest/1826/problem/F

dot (dot product，数量积，点积，内积)
https://en.wikipedia.org/wiki/Dot_product
https://en.wikipedia.org/wiki/Inner_product_space
中学课本是用力的做功引入数量积的（物体在力 F 的作用下产生位移 s，力 F 所做的功等于 |F||s|cosθ）
据此课本上定义 dot(a,b) = a·b = |a||b|cosθ，然后证明了其等于 x1x2+y1y2
a·b > 0：夹角为 [0°,90°)
a·b = 0：夹角为 90°
a·b < 0：夹角为 (90°,180°]
可以用来判断两个向量是同向还是垂直还是异向

det (determinant，行列式，平行四边形的有向面积，二维向量叉积的 z 分量的值)
https://en.wikipedia.org/wiki/Determinant#Geometric_meaning
a×b 的 z 分量的值 = |a||b|sinθ = x1y2-y1x2，其中 θ 是 a 到 b 的夹角
考虑 det(a,b) = x1y2-y1x2 的值：
> 0：b 在 a 左侧
< 0：b 在 a 右侧
= 0：a 和 b 平行或重合（共基线）
关于有向面积 https://cp-algorithms.com/geometry/oriented-triangle-area.html
https://codeforces.com/gym/105139/problem/B

1° = (π/180)rad
1rad = (180/π)°
常见的是，弧度为 2*math.Pi*(角度占整个360°的多少，设为 a) = math.Pi/180*a

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
( 0, 0) = 0   建议特殊处理

todo 二维偏序 https://ac.nowcoder.com/acm/contest/4853/F
- 题解 https://ac.nowcoder.com/discuss/394080

Pick 定理
https://en.wikipedia.org/wiki/Pick%27s_theorem
https://oi-wiki.org/geometry/pick/
https://cp-algorithms.com/geometry/picks-theorem.html
https://cp-algorithms.com/geometry/lattice-points.html
A=i+b/2-1, A为多边形面积，i为内部格点数，b为边上格点数
利用该定理可以证明：不存在坐标均为整数的正三角形、正六边形等
https://codeforces.com/problemset/problem/1548/D1
https://codeforces.com/problemset/problem/1548/D2

曼哈顿距离
TIPS: 旋转坐标
【图解】https://leetcode.cn/problems/minimize-manhattan-distances/solution/tu-jie-man-ha-dun-ju-chi-heng-deng-shi-b-op84/
顺时针旋转 45° (x,y) -> (x+y,y-x) 记作 (x',y')
曼哈顿距离转成切比雪夫距离（棋盘距离） |x1-x2|+|y1-y2| = max(|x1'-x2'|,|y1'-y2'|)
模板题 https://atcoder.jp/contests/abc178/tasks/abc178_e
LC3102 https://leetcode.cn/problems/minimize-manhattan-distances/ 2216
https://codeforces.com/problemset/problem/1689/D 1900
https://codeforces.com/problemset/problem/1979/E 2400
https://codeforces.com/problemset/problem/314/D 2500
https://www.luogu.com.cn/problem/P5098
todo LC1956 https://leetcode.cn/problems/minimum-time-for-k-virus-variants-to-spread/
点到点集的最大曼哈顿距离 https://codeforces.com/problemset/problem/491/B
更多维数 https://codeforces.com/problemset/problem/1093/G
- http://poj.org/problem?id=2926
LC1131 三维 https://leetcode.cn/problems/maximum-of-absolute-value-expression/
反向：切比雪夫距离转成曼哈顿距离 https://www.luogu.com.cn/problem/P3964
- 逆时针旋转 45° (x,y) -> ((x-y)/2, (x+y)/2)
- 切比雪夫距离的等价表述：一个点到周围八方向的相邻点的距离都是 1
https://atcoder.jp/contests/abc351/tasks/abc351_e

TIPS: 另一种方法是分四种情况讨论，即
|a-b|+|c-d|
= max(a-b, b-a) + max(c-d, d-c)
= max((a-b)+(c-d), (b-a)+(c-d), (a-b)+(d-c), (b-a)+(d-c))
https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/discuss/489882/O(n)-Solution-with-explanation
纯数学证明 https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solution/bu-hui-hua-jian-qing-kan-zhe-pythonjavac-c2s6/
LC1330 https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/
- todo 上面这题求最小 https://atcoder.jp/contests/arc119/tasks/arc119_e
- https://codeforces.com/problemset/problem/1898/D
LC1131 三维 https://leetcode.cn/problems/maximum-of-absolute-value-expression/

曼哈顿最近点对
LC2613 https://leetcode.cn/problems/beautiful-pairs/

梯形个数
https://leetcode.cn/problems/count-number-of-trapezoids-ii/
思路有些类似 https://codeforces.com/problemset/problem/1163/C2

https://oeis.org/A053411 Circle numbers
a(n)= number of points (i,j), i,j integers, contained in a circle of diameter n, centered at the origin

https://oeis.org/A136485 Number of unit square lattice cells enclosed by origin centered circle of diameter n

湖边野猫追捕水中老鼠，为什么速度比低于 4.6033 就逮不到老鼠？ https://zhuanlan.zhihu.com/p/113905393

为什么常见的O(n)时间的找凸多边形内最大内接三角形的那个算法是错误的? https://www.zhihu.com/question/60726749

*/

// isqrt 返回 floor(sqrt(x))
// 函数名来自 Python 中的 math.isqrt
// 由于 float64 无法【精确】表示过大的 int（超出 2^53-1 的），需要微调   
// 注：【精确】指代码中的 float64 必须要对应着唯一的 int，才能保证计算结果正确 https://www.zhihu.com/question/29010688
//
// 举例说明（上舍入错误）
// x                  floor(sqrt(x)) 注意这里算的是错误的结果，需要用 if 调整，见下面代码
// 999999999999999999 1000000000
// 999996000003999999 999998000
// 880200476099999999 938190000
//
// 附：https://www.binaryconvert.com/convert_double.html
// 0b1(Mantissa) * 2 ** (0b(Exponent) - 1075)
//
// 可以用以下题目测试：
// https://atcoder.jp/contests/abc191/tasks/abc191_d 
// https://atcoder.jp/contests/abc243/tasks/abc243_g
// https://codeforces.com/problemset/problem/1036/F
// https://codeforces.com/problemset/problem/1862/D
func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x))) // 可能会算多一点点
	if rt*rt > x {
		rt--
	}
	return rt
}

// isSQ 返回 x 是否为完全平方数
func isSQ(x int) bool {
	rt := isqrt(x)
	return rt*rt == x
}

// 返回 floor(pow(x, 1/n))
// x>=0, n>1
func floorRootN(x, n int) int {
	if x == 0 {
		return 0
	}
	if n > 62 {
		return 1
	}
	res := int(math.Pow(float64(x), 1/float64(n)))
	// 误差修正
	if pow(res, n) > x {
		res--
	} else if pow(res+1, n) <= x {
		res++
	}
	return res
}

// 什么时候 (i-1)/i == (i-2)/(i-1) ?
// 最小的正整数 i 是一个比 1e8 略小的数，比 floor(2^(53/2)) = 94906265 略大一点
// 94911151
// 94917164
// 94920884
// 94923833
// 94926354
// 94928593
// 94930626
// 94932503
// 94934254
// 94935901
// 为什么？如果两个数的差小于 2^-52，就可能会 round 到同一个浮点数上
// LC2280 https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/
//
// 附：i=134229312 和 i=134229313 这两个相邻的数都满足，这是最小的同时满足且相邻的数字
//
// 附：最大的几个负数，比 -2^(52/2) = -67108864 略小一点
// -67114657
// -67118898
// -67121818
// -67124192
// -67126245
func roundErrorExample() {
	cnt := 0
	for i := 10; cnt < 10; i++ {
		if float64(i-1)/float64(i) == float64(i)/float64(i+1) {
			fmt.Println(i)
			cnt++
		}
	}
}

// 浮点数 GCD
// https://codeforces.com/problemset/problem/1/C 2100
func gcdf(a, b float64) float64 {
	// 注意根据题目约束，分析 eps 取值
	// 例如 CF1C，由于保证正多边形边数不超过 100，故 gcdf 的结果不会小于 2*math.Pi/100，eps 可以取 1e-2
	for a > eps { // 如果有负数，改成 math.Abs(a) > eps
		a, b = math.Mod(b, a), a
	}
	return b
}

const eps = 1e-8

/*
二维向量（点）
点与点
*/
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y } // 点积，用来判断两向量同向/异向
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x } // 叉积的 z 分量的值，用来判断两向量的左右关系   cross
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func (a vec) add(b vec) vec       { return vec{a.x + b.x, a.y + b.y} }
func (a vec) mul(k int) vec       { return vec{a.x * k, a.y * k} }
func (a vecF) div(k float64) vecF { return vecF{a.x / k, a.y / k} }
func (a vec) len2() int           { return a.x*a.x + a.y*a.y }                     // 模长的平方
func (a vec) len() float64        { return math.Sqrt(float64(a.x*a.x + a.y*a.y)) } // 模长
func (a vec) dis2(b vec) int      { return b.sub(a).len2() }
func (a vec) dis(b vec) float64   { return b.sub(a).len() }

func (a *vec) adds(b vec)      { a.x += b.x; a.y += b.y }
func (a *vec) subs(b vec)      { a.x -= b.x; a.y -= b.y }
func (a *vec) muls(k int)      { a.x *= k; a.y *= k }
func (a *vecF) divs(k float64) { a.x /= k; a.y /= k }

func (a vec) less(b vec) bool     { return a.x < b.x || a.x == b.x && a.y < b.y }
func (a vecF) less(b vecF) bool   { return a.x+eps < b.x || a.x < b.x+eps && a.y+eps < b.y }
func (a vecF) equals(b vecF) bool { return math.Abs(a.x-b.x) < eps && math.Abs(a.y-b.y) < eps }
func (a vec) parallel(b vec) bool { return a.det(b) == 0 }
func (a vec) mulVec(b vec) vec    { return vec{a.x*b.x - a.y*b.y, a.x*b.y + b.x*a.y} }
func (a vec) polarAngle() float64 { return math.Atan2(float64(a.y), float64(a.x)) }
func (a vec) reverse() vec        { return vec{-a.x, -a.y} } // 反向
func (a vec) sign() int {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return -1
	}
	return 1
}
func (a vec) up() vec {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}
func (a vec) rotateCCW90() vec { return vec{-a.y, a.x} } // 逆时针旋转 90°
func (a vec) rotateCW90() vec  { return vec{a.y, -a.x} } // 顺时针旋转 90°

// 逆时针旋转，传入旋转的弧度
// https://codeforces.com/problemset/problem/618/E 2500
// 注意：如果传入的是角度 angle，那么 rad = float64(angle%360) / 180 * math.Pi，这里模 360 很重要，避免过大的浮点误差
func (a vec) rotateCCW(rad float64) vecF {
	sin, cos := math.Sincos(rad)
	return vecF{
		float64(a.x)*cos - float64(a.y)*sin,
		float64(a.x)*sin + float64(a.y)*cos,
	}
}

// 单位向量
// https://en.wikipedia.org/wiki/Unit_vector
func (a vecF) unit() vecF { return a.div(a.len()) }

// a 的单位法线（a 不能是零向量）
// https://en.wikipedia.org/wiki/Normal_(geometry)
func (a vecF) normal() vecF { return vecF{-a.y, a.x}.unit() }

// 转化为长度为 x 的向量
func (a vecF) trunc(x float64) vecF { return a.unit().mul(x) }

// 两向量夹角
//    det = |A||B|sinθ
//    dot = |A||B|cosθ
// => tanθ = det / dot
// b 在 a 左侧时返回正值
// b 在 a 右侧时返回负值
func (a vec) angleTo(b vec) float64 { return math.Atan2(float64(a.det(b)), float64(a.dot(b))) }

// 极角排序
// todo 给 1e5 个点，求包含原点的三角形个数 https://www.luogu.com.cn/problem/P2992
// - 考虑补集
// - 点和原点连直线，在直线一侧选两个点组成的三角形必然不会包含原点
// - 双指针维护
// 从 n 个向量中选出一些向量，使得向量和的模长最大 https://atcoder.jp/contests/abc139/tasks/abc139_f
// https://codeforces.com/contest/1841/problem/F
func polarAngleSort(ps []vec) {
	// (-π, π]
	// (-1e9,-1) -> (-1e9, 0)
	// 也可以先把每个向量的极角算出来再排序
	sort.Slice(ps, func(i, j int) bool { return ps[i].polarAngle() < ps[j].polarAngle() })

	// EXTRA: 若允许将所有向量 up()，由于 up() 后向量的范围是 [0°, 180°)，可以用叉积排序，且排序后共线的向量一定会相邻
	for i := range ps {
		ps[i] = ps[i].up()
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].det(ps[j]) > 0 })
}

// 余弦定理，输入两边及夹角，计算对边长度
func cosineRule(a, b, angle float64) float64 {
	return math.Sqrt(a*a*b*b - 2*a*b*math.Cos(angle))
}
func cosineRuleVec(va, vb vecF, angle float64) float64 {
	return math.Sqrt(va.len2() + vb.len2() - 2*va.len()*vb.len()*math.Cos(angle))
}

// 三角形外心（外接圆圆心，三条边的垂直平分线的交点）
// 另一种写法是求两中垂线交点
// https://en.wikipedia.org/wiki/Circumscribed_circle
// https://codeforces.com/problemset/problem/1/C
func circumcenter(a, b, c vecF) vecF {
	a1, b1, a2, b2 := b.x-a.x, b.y-a.y, c.x-a.x, c.y-a.y
	c1, c2, d := a1*a1+b1*b1, a2*a2+b2*b2, 2*(a1*b2-a2*b1)
	if math.Abs(d) < eps { // 三点共线（需要根据题目要求特判）
		return vecF{math.NaN(), math.NaN()}
	}
	// 也可以打开括号，先 /d 再做乘法，提高精度
	return vecF{a.x + (c1*b2-c2*b1)/d, a.y + (a1*c2-a2*c1)/d}
}

// EXTRA: 外接圆半径 R
// 下面交换了一下乘除的顺序，减小精度的丢失
// todo https://codeforces.com/problemset/problem/274/C
func circumcenterR(a, b, c vecF) float64 {
	ab, ac := b.sub(a), c.sub(a)
	return 0.5 * a.dis(b) * a.dis(c) / math.Abs(ab.det(ac)) * b.dis(c)
}
func circumcenterR2(a, b, c vecF) float64 {
	ab, ac := b.sub(a), c.sub(a)
	return 0.25 * a.dis2(b) / ab.det(ac) * a.dis2(c) / ab.det(ac) * b.dis2(c)
}

// 三角形垂心（三条高的交点）
// https://en.wikipedia.org/wiki/Altitude_(triangle)#Orthocenter
// 欧拉线上的四点中，九点圆圆心到垂心和外心的距离相等，而且重心到外心的距离是重心到垂心距离的一半。注意内心一般不在欧拉线上，除了等腰三角形外
// https://en.wikipedia.org/wiki/Euler_line
// https://baike.baidu.com/item/%E4%B8%89%E8%A7%92%E5%BD%A2%E4%BA%94%E5%BF%83%E5%AE%9A%E5%BE%8B
func orthocenter(a, b, c vecF) vecF {
	return a.add(b).add(c).sub(circumcenter(a, b, c).mul(2))
}

// 三角形内心（三条角平分线的交点）
// 三点坐标按对边长度加权平均
// https://en.wikipedia.org/wiki/Incenter
func incenter(a, b, c vecF) vecF {
	bc, ac, ab := b.dis(c), a.dis(c), a.dis(b)
	sum := bc + ac + ab
	return vecF{(bc*a.x + ac*b.x + ab*c.x) / sum, (bc*a.y + ac*b.y + ab*c.y) / sum}
}

/*
二维直线（线段）
二维点与线、线与线
*/
type line struct{ p1, p2 vec }

// 方向向量 directional vector
func (a line) vec() vec              { return a.p2.sub(a.p1) }
func (a lineF) point(t float64) vecF { return a.p1.add(a.vec().mul(t)) }

// 点 a 是否严格在 l 左侧
func (a vecF) onLeft(l lineF) bool { return l.vec().det(a.sub(l.p1)) > eps }

// 点 a 是否在直线 l 上
// 判断方法：a-p1 与 a-p2 共线
func (a vec) onLine(l line) bool {
	p1, p2 := l.p1.sub(a), l.p2.sub(a)
	return p1.det(p2) == 0
}

// 点 a 是否在线段 l 上
// 判断方法：a-p1 与 a-p2 共线且方向相反
func (a vec) onSeg(l line) bool {
	p1, p2 := l.p1.sub(a), l.p2.sub(a)
	return p1.det(p2) == 0 && p1.dot(p2) <= 0 // 含端点
}

func (a vecF) onSegF(l lineF) bool {
	p1, p2 := l.p1.sub(a), l.p2.sub(a)
	return math.Abs(p1.det(p2)) < eps && p1.dot(p2) < eps // 含端点
}

// 点 a 是否在射线 o-d 上（d 是向量）
// 判断方法：o-a 与 d 共线且方向相同
func (a vec) onRay(o, d vec) bool {
	a = a.sub(o)
	return d.det(a) == 0 && d.dot(a) >= 0 // 含端点
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

// 判断点 a 到线段 l 的距离 <= r，避免浮点运算
func (a vec) withinRange(l line, r int) bool {
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
// https://codeforces.com/contest/1826/problem/F
func (a vecF) projection(l lineF) vecF {
	v, p1a := l.vec(), a.sub(l.p1)
	t := v.dot(p1a) / v.len2()
	return l.p1.add(v.mul(t))
}

// 点 a 关于直线 l 的对称点
// 求投影 p，然后将 ap 延长一倍
func (a vecF) symmetry(l lineF) vecF {
	return a.add(a.projection(l).sub(a).mul(2))
}

// 判断点 a 是否在以线段 l 为对角线的矩形内（矩形边界与坐标轴平行）
func (a vec) inRect(l line) bool {
	return (a.x >= l.p1.x || a.x >= l.p2.x) &&
		(a.x <= l.p1.x || a.x <= l.p2.x) &&
		(a.y >= l.p1.y || a.y >= l.p2.y) &&
		(a.y <= l.p1.y || a.y <= l.p2.y)
}

// 直线 a b 交点 - 参数式
// 若求线段交点，可以在求出后判断其是否均在两条线段上：由于交点已在基线上，只需要判断交点是否在以线段为对角线的矩形内即可
// NOTE: 若输入均为有理数，则输出也为有理数，对精度要求较高时可使用分数类
func (a lineF) intersection(b lineF) vecF {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	d := va.det(vb)
	if math.Abs(d) < eps { // 两直线平行（特判）
		return vecF{math.NaN(), math.NaN()}
	}
	t := vb.det(u) / d
	return a.point(t)
}

// 直线 a b 交点 - 两点式
// 线段求整数交点 https://codeforces.com/problemset/problem/1036/E
// 相关 https://leetcode.cn/problems/intersection-lcci/
func (a lineF) intersection2(b lineF) vecF {
	va := a.vec()
	d1 := va.det(b.p1.sub(a.p1))
	d2 := va.det(b.p2.sub(a.p1))
	if math.Abs(d1-d2) < eps { // 两直线平行（特判）
		return vecF{math.NaN(), math.NaN()}
	}
	return vecF{b.p1.x*d2 - b.p2.x*d1, b.p1.y*d2 - b.p2.y*d1}.div(d2 - d1)
}

// 射线 a b 交点 c
// 返回 a b 各自到首个交点所需的时间（射线速度由 .vec().len() 决定）
// 无交点时返回 -1
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

// segIntersection 返回线段 a 和 b 的相交状态 state，以及交点
// 如果有无穷个交点，返回其中 x 最小的交点（垂线取 y 最小）left 和 x 最大的交点（垂线取 y 最大）right
// 0：不相交
// 1：规范相交（交点不是线段端点）
// 2：在端点相交
// 3：重合（有无穷个交点）
// 面试题 16.03. 交点 https://leetcode.cn/problems/intersection-lcci/
// https://leetcode.cn/contest/autox2023/problems/TcdlJS/
func (a lineF) segIntersection(b lineF) (state int, left, right vecF) {
	if max(a.p1.x, a.p2.x) < min(b.p1.x, b.p2.x) ||
		min(a.p1.x, a.p2.x) > max(b.p1.x, b.p2.x) ||
		max(a.p1.y, a.p2.y) < min(b.p1.y, b.p2.y) ||
		min(a.p1.y, a.p2.y) > max(b.p1.y, b.p2.y) {
		return // 不相交（两条线段的矩形框不相交）
	}

	va, vb := a.vec(), b.vec()
	if va.det(vb) == 0 { // 平行
		if va.det(b.p1.sub(a.p1)) != 0 {
			return // 不相交（两条线段不在同一条基线上）
		}
		minAx, maxAx := min(a.p1.x, a.p2.x), max(a.p1.x, a.p2.x)
		minAy, maxAy := min(a.p1.y, a.p2.y), max(a.p1.y, a.p2.y)
		minBx, maxBx := min(b.p1.x, b.p2.x), max(b.p1.x, b.p2.x)
		minBy, maxBy := min(b.p1.y, b.p2.y), max(b.p1.y, b.p2.y)
		left = vecF{max(minAx, minBx), max(minAy, minBy)}
		right = vecF{min(maxAx, maxBx), min(maxAy, maxBy)}
		if left == right {
			return 2, left, left // 在端点相交
		}
		//if !left.onSegF(a) {
		//	left.y, right.y = right.y, left.y
		//}
		return 3, left, right // 有无穷个交点
	}

	//d1, d2 := va.det(b.p1.sub(a.p1)), va.det(b.p2.sub(a.p1))
	//d3, d4 := vb.det(a.p1.sub(b.p1)), vb.det(a.p2.sub(b.p1))
	//if sign(d1)*sign(d2) < 0 && sign(d3)*sign(d4) < 0 { return 1 }
	d1 := b.p1.sub(a.p1).det(b.p2.sub(a.p1))
	d2 := b.p1.sub(a.p2).det(b.p2.sub(a.p2))
	if d1 > 0 && d2 > 0 || d1 < 0 && d2 < 0 {
		return // 不相交
	}
	d3 := a.p1.sub(b.p1).det(a.p2.sub(b.p1))
	d4 := a.p1.sub(b.p2).det(a.p2.sub(b.p2))
	if d3 > 0 && d4 > 0 || d3 < 0 && d4 < 0 {
		return // 不相交
	}

	p := a.intersection(b)
	if d1 == 0 || d2 == 0 || d3 == 0 || d4 == 0 {
		return 2, p, p // 在端点相交
	}
	return 1, p, p // 规范相交
}

// 过点 a 的垂直于 l 的直线
func (a vec) perpendicular(l line) line {
	return line{a, a.add(vec{l.p1.y - l.p2.y, l.p2.x - l.p1.x})}
}

/*
圆
线与圆
圆与圆

https://codeforces.com/problemset/problem/908/C 1500

*/
type circle struct {
	vec
	r int
}

// 圆心角对应的点
func (o circle) point(rad float64) vecF {
	sin, cos := math.Sincos(rad)
	return vecF{float64(o.x) + float64(o.r)*cos, float64(o.y) + float64(o.r)*sin}
}
func (o circleF) point(rad float64) vecF {
	sin, cos := math.Sincos(rad)
	return vecF{o.x + o.r*cos, o.y + o.r*sin}
}

// 三点确定一圆
// 用三角形外心求解，见 circumcenter

// 给定半径和一条有向的弦，求该弦右侧的圆心（即 ao 在 ab 右侧）
func getCircleCenter(a, b vec, r int) vecF {
	disAB2 := b.sub(a).len2()
	if disAB2 > 4*r*r { // 不存在
		return vecF{math.NaN(), math.NaN()}
	}
	midX, midY := float64(a.x+b.x)/2, float64(a.y+b.y)/2
	d := math.Sqrt(float64(r*r) - float64(disAB2/4))
	angle := math.Atan2(float64(b.y-a.y), float64(b.x-a.x))
	return vecF{midX + d*math.Sin(angle), midY - d*math.Cos(angle)}
}

// 直线与圆的交点
// t1 <= t2
// 射线是否相交：t2 >= 0
// - https://www.acwing.com/problem/content/4502/
// 线段与圆是否相交：0 <= t1, t2 <= 1
// - https://leetcode.cn/contest/autox2023/problems/TcdlJS/
func (o circleF) intersectionLine(l lineF) (ps []vecF, t1, t2 float64) {
	v := l.vec()
	// 需要保证 v 不会退化成一个点
	if v.x == 0 && v.y == 0 {
		// 根据题意特判
	}
	a, b, c, d := v.x, l.p1.x-o.x, v.y, l.p1.y-o.y
	e, f, g := a*a+c*c, 2*(a*b+c*d), b*b+d*d-o.r*o.r
	delta := f*f - 4*e*g // 注意这会达到值域的 4 次方
	switch {
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
// 另一种写法是解二元二次方程组，精度更优
// https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable/
func (o circle) intersectionCircle(b circle) (ps []vecF, normal bool) {
	a := o
	if a.r < b.r {
		a, b = b, a // 注意这里 swap 了
	}
	ab := b.sub(a.vec)
	dab2 := ab.len2()
	diffR, sumR := a.r-b.r, a.r+b.r
	if dab2 == 0 { // 圆心重合
		if diffR == 0 { // 两圆完全一样
			return
		}
		return nil, true // 一个包含另一个
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

// 与两圆外切的圆的圆心
// 挑战 p.275
// 记圆心在 (x,y)，半径为 r 的圆为 O1，
// 另有一半径为 R 的圆 O，若 O 与 O1 相切，
// 则 O 的圆心轨迹形成了一个圆心在 (x,y)，半径为 R-r 的圆
// 因此，问题变成了求两个圆的交点

// 圆的面积并 - 两圆的特殊情形
// 两圆面积交 = 面积和 - 面积并
// todo https://codeforces.com/problemset/problem/600/D

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
		return []vecF{po.rotateCCW(math.Pi / 2)}
	} // 圆上一点的切线
	ang := math.Asin(float64(o.r) / po.len())
	return []vecF{po.rotateCCW(-ang), po.rotateCCW(ang)}
}

// 两圆公切线
// 返回每条切线在圆 o 和圆 ob 的切点
// NOTE: 下面的代码是基于 int 的，没有判断 eps
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

// 最小圆覆盖 Welzl's algorithm
// 随机增量法，期望复杂度 O(n)
// 详见《计算几何：算法与应用（第 3 版）》第 4.7 节
// https://en.wikipedia.org/wiki/Smallest-circle_problem
// https://oi-wiki.org/geometry/random-incremental/
// 模板题 https://www.luogu.com.cn/problem/P1742 
//       https://www.luogu.com.cn/problem/P2533
//       https://atcoder.jp/contests/abc151/tasks/abc151_f
//       LC1924 https://leetcode.cn/problems/erect-the-fence-ii/
// 椭圆（坐标系旋转缩一下） https://www.luogu.com.cn/problem/P4288
func smallestEnclosingDisc(ps []vecF) circleF {
	rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })
	o := ps[0]
	r2 := 0.
	for i, p := range ps {
		if p.dis2(o) < r2+eps { // p 在最小圆内部或边上
			continue
		}
		o, r2 = p, 0
		for j, q := range ps[:i] {
			if q.dis2(o) < r2+eps { // q 在最小圆内部或边上
				continue
			}
			o = vecF{(p.x + q.x) / 2, (p.y + q.y) / 2}
			r2 = p.dis2(o)
			for _, x := range ps[:j] {
				if x.dis2(o) > r2+eps { // 保证三点不会共线（证明略）
					o = circumcenter(p, q, x)
					r2 = p.dis2(o)
				}
			}
		}
	}
	return circleF{o, math.Sqrt(r2)}
}

// 求一固定半径的圆最多能覆盖多少个点（圆边上也算覆盖） len(ps)>0 && r>0
// Angular Sweep 算法 O(n^2logn)
// https://www.geeksforgeeks.org/angular-sweep-maximum-points-can-enclosed-circle-given-radius/
// LC1453 https://leetcode.cn/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/solution/python3-angular-sweepsuan-fa-by-lih/
func maxCoveredPoints(ps []vec, r int) int {
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
		sort.Slice(events, func(i, j int) bool {
			a, b := events[i], events[j]
			return a.angle+eps < b.angle || a.angle < b.angle+eps && a.delta > b.delta
		})
		mx, cnt := 0, 1 // 1 指当前固定的点 a
		for _, e := range events {
			cnt += e.delta
			mx = max(mx, cnt)
		}
		ans = max(ans, mx)
	}
	return ans
}

// 圆和矩形是否重叠
// x1<x2, y1<y2
// https://www.zhihu.com/question/24251545/answer/27184960
// LC1401 https://leetcode.cn/problems/circle-and-rectangle-overlapping/
// todo 返回和矩形的那些边相交/相切 
//   https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable/
func isCircleRectangleOverlap(r, ox, oy, x1, y1, x2, y2 int) bool {
	cx, cy := float64(x1+x2)/2, float64(y1+y2)/2               // 矩形中心
	hx, hy := float64(x2-x1)/2, float64(y2-y1)/2               // 矩形半长
	x, y := math.Abs(float64(ox)-cx), math.Abs(float64(oy)-cy) // 转换到第一象限的圆心
	x, y = max(x-hx, 0), max(y-hy, 0)                          // 求圆心至矩形的最短距离矢量
	return x*x+y*y < float64(r*r)+eps
}

// todo 两圆的【两个交点的连线】与【圆心连线】的交点
//  https://leetcode.cn/circle/discuss/GNUiDD/view/Q4wDFD/

// 圆与扫描线
// todo https://blog.csdn.net/hzj1054689699/article/details/87861808
//  http://openinx.github.io/2013/01/01/plane-sweep-thinking/
//  http://poj.org/problem?id=2932
//  https://ac.nowcoder.com/acm/contest/7613/D
//  https://codeforces.com/problemset/problem/814/D

// 反演变换
// https://en.wikipedia.org/wiki/Inversive_geometry
// todo https://oi-wiki.org/geometry/inverse/

// 三角剖分
// todo https://oi-wiki.org/geometry/triangulation/
//  https://cp-algorithms.com/geometry/delaunay.html
//  http://poj.org/problem?id=2986

// 多边形相关
func _(abs func(int) int) {
	readVec := func(in io.Reader) vec {
		var x, y int
		fmt.Fscan(in, &x, &y)
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

	// todo 扫描线：线段求交 O(nlogn)
	// https://en.wikipedia.org/wiki/Sweep_line_algorithm
	// N 条线段求交的扫描线算法 http://johnhany.net/2013/11/sweep-algorithm-for-segments-intersection/
	// https://codeforces.com/problemset/problem/1359/F
	// 平面扫描思想在 ACM 竞赛中的应用 http://openinx.github.io/2013/01/01/plane-sweep-thinking/

	// 按 y 归并
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
	// 返回最近点对距离的平方
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/ClosestPair.java.html
	// 模板题 https://www.luogu.com.cn/problem/P1429
	// - https://www.luogu.com.cn/problem/P7883 加强版
	// - https://codeforces.com/problemset/problem/429/D
	// todo 字典序最小 LC2613 https://leetcode.cn/problems/beautiful-pairs/
	// 变形：改成求「距离为平面最近点对距离的两倍」的点对个数，做法应该是类似的
	// bichromatic closest pair 有两种类型的点，只需要额外判断类型是否不同即可 https://www.acwing.com/problem/content/121/ http://poj.org/problem?id=3714
	var closestPair func([]vec) int
	closestPair = func(ps []vec) int {
		// ！调用 closestPair 前需保证没有重复的点，并特判 n == 1 的情况
		// ！ps 必须按照 x 坐标升序：
		//   slices.SortFunc(ps, func(a, b vec) int { return a.x - b.x })
		n := len(ps)
		if n <= 1 {
			return math.MaxInt
		}
		m := n / 2
		x := ps[m].x
		d2 := min(closestPair(ps[:m]), closestPair(ps[m:]))
		copy(ps, merge(ps[:m], ps[m:])) // copy 是因为要修改 slice 的内容
		checkPs := []vec{}
		for _, pi := range ps {
			if (pi.x-x)*(pi.x-x) > d2 {
				continue
			}
			for j := len(checkPs) - 1; j >= 0; j-- {
				pj := checkPs[j]
				dy := pi.y - pj.y
				if dy*dy >= d2 {
					break
				}
				dx := pi.x - pj.x
				d2 = min(d2, dx*dx+dy*dy)
			}
			checkPs = append(checkPs, pi)
		}
		return d2
	}

	// 读入多边形
	// 输入的点必须按顺时针或逆时针顺序输入
	readPolygon := func(in io.Reader, n int) []line {
		ps := make([]vec, n)
		for i := range ps {
			fmt.Fscan(in, &ps[i].x, &ps[i].y)
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
		area := 0
		p0 := ps[0]
		for i := 2; i < len(ps); i++ {
			area += ps[i-1].sub(p0).det(ps[i].sub(p0))
		}
		return float64(area) / 2
	}

	// 求凸包 Andrew 算法（Andrew's monotone chain convex hull algorithm）
	// 使用单调栈，保存的向量是有极角序的
	// 求下凸包：从最左边的点开始遍历，同时用一根绳子逆时针绕圈，理想的顺序是下一个点的位置在绳子前进方向的左侧，如果某个点会导致绳子向右走，那么就需要出栈
	// 求上凸包就从倒数第二个点开始继续求
	// 允许 ps 里面有重复的点，算法会去重
	// https://en.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
	// https://en.wikipedia.org/wiki/Convex_hull_algorithms
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GrahamScan.java.html
	// NOTE: 坐标值范围不超过 M 的整点凸多边形的顶点数为 O(M^(2/3)) 个
	//
	// https://www.luogu.com.cn/problem/P2742 模板题
	// - LC587 https://leetcode.cn/problems/erect-the-fence/
	// LC3494 https://leetcode.cn/problems/find-the-minimum-amount-of-time-to-brew-potions/
	// https://codeforces.com/problemset/problem/1017/E 2400
	// https://codeforces.com/problemset/problem/1142/C 2400
	// https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A 限制区间长度的区间最大均值问题
	// https://atcoder.jp/contests/abc275/tasks/abc275_g 转换（求二维完全背包极限 lim_{maxW->∞} dp[maxW]/maxW）
	// 构造 LCP15 https://leetcode.cn/problems/you-le-yuan-de-mi-gong/
	// todo poj 2187 1113 1912 3608 2079 3246 3689
	convexHull := func(ps []vec) (q []vec) {
		slices.SortFunc(ps, func(a, b vec) int { return cmp.Or(a.x-b.x, a.y-b.y) })
		//sort.Slice(ps, func(i, j int) bool { return ps[i].less(ps[j]) })
		// 下凸包（从左到右）
		for _, p := range ps {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
		// 上凸包（从右到左）
		downSize := len(q)
		// 注意下凸包的最后一个点，已经是上凸包的（右边）第一个点了，所以从 n-2 开始遍历
		for i := len(ps) - 2; i >= 0; i-- {
			p := ps[i]
			for len(q) > downSize && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
		// 此时首尾是同一个点 ps[0]，需要去掉
		q = q[:len(q)-1]
		return
	}

	// 只计算上凸包（从左到右）
	// https://leetcode.cn/problems/find-the-minimum-amount-of-time-to-brew-potions/
	convexHull2 := func(ps []vec) (q []vec) {
		//slices.SortFunc(ps, func(a, b vec) int { return cmp.Or(a.x-b.x, a.y-b.y) })
		sort.Slice(ps, func(i, j int) bool { return ps[i].less(ps[j]) })
		for _, p := range ps {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) >= 0 { // 注意这里是 >=
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
		return
	}

	// 凸包哈希，用于判断两个凸包是否相等
	// 边长 - 夹角 - 边长 - 夹角 - 边长 - …… 绕一圈，必须有 n 个边长和 n 个夹角
	// https://codeforces.com/problemset/problem/1017/E 2400
	convexHullHash := func(a []vec) []int {
		// ！输入的 a 必须是凸包，且至少有两个点

		// 首尾是同一个点，方便计算
		a = append(a, a[0])

		res := make([]int, 1, len(a)*2-2)
		res[0] = a[0].dis2(a[1])
		for i := 2; i < len(a); i++ {
			// 用斜边长代替夹角（SSS 全等），避免浮点数
			// 取负号是为了区分边长与斜边长
			res = append(res, -a[i-2].dis2(a[i]), a[i-1].dis2(a[i]))
		}
		res = append(res, -a[len(a)-2].dis2(a[1]))

		return res // 然后计算 smallestRepresentation(res)，见 strings.go
	}

	// 动态凸包
	// 用平衡树加点
	// https://en.wikipedia.org/wiki/Dynamic_convex_hull
	// 模板题 https://codeforces.com/problemset/problem/70/D 2700
	// https://codeforces.com/problemset/problem/932/F 2700 启发式合并
	// https://www.luogu.com.cn/problem/P2521
	// https://www.luogu.com.cn/problem/P4027
	dynamicConvexHull := func(ps []vec) {
		det := func(x1, y1, x2, y2 int) int { return x1*y2 - x2*y1 }
		_remove := func(t *treapM[int, int], i int) bool {
			if i <= 0 || i+1 >= t.size() {
				return false
			}
			pre := t.kth(i - 1)
			cur := t.kth(i)
			nxt := t.kth(i + 1)
			if det(cur.key-pre.key, cur.value-pre.value, nxt.key-pre.key, nxt.value-pre.value) >= 0 {
				t.delete(cur.key)
				return true
			}
			return false
		}
		contains := func(t *treapM[int, int], x, y int) bool {
			i := t.lowerBoundIndex(x)
			if i == t.size() {
				return false
			}
			cur := t.kth(i)
			if cur.key == x {
				return y <= cur.value
			}
			if i == 0 {
				return false
			}
			pre := t.kth(i - 1)
			return det(cur.key-pre.key, cur.value-pre.value, x-pre.key, y-pre.value) <= 0
		}
		insert := func(t *treapM[int, int], x, y int) bool {
			if contains(t, x, y) {
				return false
			}
			t.put(x, y)
			idx := t.lowerBoundIndex(x)
			for j := idx + 1; _remove(t, j); {
			}
			for j := idx - 1; _remove(t, j); j-- {
			}
			return true
		}

		topC := newMap[int, int]()  // 上凸包
		downC := newMap[int, int]() // 下凸包
		addPoint := func(x, y int) bool {
			ok1 := insert(topC, x, y)
			ok2 := insert(downC, x, -y) // 上下都要插入
			return ok1 || ok2
		}
		hasPoint := func(x, y int) bool {
			return contains(topC, x, y) && contains(downC, x, -y)
		}

		_ = []any{addPoint, hasPoint}
	}

	// 旋转卡壳求最远点对（凸包直径） Rotating calipers
	// https://en.wikipedia.org/wiki/Rotating_calipers
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FarthestPair.java.html
	// 模板题 https://www.luogu.com.cn/problem/P1452
	rotatingCalipers := func(ps []vec) (p1, p2 vec) {
		ch := convexHull(ps)
		n := len(ch)
		if n == 2 {
			// maxD2 := ch[0].dis2(ch[1])
			return ch[0], ch[1]
		}
		i, j := 0, 0
		for k, p := range ch {
			if !ch[i].less(p) {
				i = k
			}
			if ch[j].less(p) {
				j = k
			}
		}
		maxD2 := 0
		for i0, j0 := i, j; i != j0 || j != i0; {
			if d2 := ch[i].dis2(ch[j]); d2 > maxD2 {
				maxD2 = d2
				p1, p2 = ch[i], ch[j]
			}
			if ch[(i+1)%n].sub(ch[i]).det(ch[(j+1)%n].sub(ch[j])) < 0 {
				i = (i + 1) % n
			} else {
				j = (j + 1) % n
			}
		}
		return
	}

	// todo 最小矩形覆盖/最小外接矩形
	// https://www.luogu.com.cn/problem/P3187
	// UVa 10173 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=13&page=show_problem&problem=1114

	// todo 闵可夫斯基和 Minkowski sum
	// https://en.wikipedia.org/wiki/Minkowski_addition
	// https://www.cnblogs.com/xzyxzy/p/10229921.html
	// https://www.luogu.com.cn/problem/P4557
	// https://codeforces.com/problemset/problem/87/E
	// https://codeforces.com/contest/1841/problem/F

	// todo 点集的最大四边形
	// https://www.luogu.com.cn/problem/P4166
	// https://codeforces.com/contest/340/problem/B

	// 凸包周长
	convexHullPerimeter := func(ps []vec) (l float64) {
		ch := convexHull(ps) // 注意 convexHull 需要去掉 q = q[:len(q)-1] 这行
		for i := 1; i < len(ch); i++ {
			l += ch[i].dis(ch[i-1])
		}
		return
	}

	// 半平面交
	// O(nlogn)，时间开销主要在排序上
	// 大致思路：首先极角排序，然后用一个队列维护半平面交的顶点，每添加一个半平面，就不断检查队首队尾是否在半平面外，是就剔除
	// 注意要先剔除队尾再剔除队首
	// 注：凸包的对偶问题很接近半平面交，所以二者算法很接近
	// https://oi-wiki.org/geometry/half-plane/
	// https://www.luogu.com.cn/blog/105254/dui-ban-ping-mian-jiao-suan-fa-zheng-que-xing-xie-shi-di-tan-suo
	// 模板题 https://www.luogu.com.cn/problem/P4196 https://www.acwing.com/problem/content/2805/
	// todo https://www.luogu.com.cn/problem/P3256 https://www.acwing.com/problem/content/2960/
	type lp struct {
		l lineF
		p vecF // l 与下一条直线的交点
	}
	halfPlanesIntersection := func(ls []lineF) []lp { // 规定左侧为半平面
		sort.Slice(ls, func(i, j int) bool { return ls[i].vec().polarAngle() < ls[j].vec().polarAngle() })
		q := []lp{{l: ls[0]}}
		for i := 1; i < len(ls); i++ {
			l := ls[i]
			for len(q) > 1 && !q[len(q)-2].p.onLeft(l) {
				q = q[:len(q)-1]
			}
			for len(q) > 1 && !q[0].p.onLeft(l) {
				q = q[1:]
			}
			if math.Abs(l.vec().det(q[len(q)-1].l.vec())) < eps {
				// 由于极角排序过，此时两有向直线平行且同向，取更靠内侧的直线
				if l.p1.onLeft(q[len(q)-1].l) {
					q[len(q)-1].l = l
				}
			} else {
				q = append(q, lp{l: l})
			}
			if len(q) > 1 {
				q[len(q)-2].p = q[len(q)-2].l.intersection(q[len(q)-1].l)
			}
		}
		// 最后用队首检查下队尾，删除无用半平面
		for len(q) > 1 && !q[len(q)-2].p.onLeft(q[0].l) {
			q = q[:len(q)-1]
		}

		if len(q) < 3 {
			// 半平面交不足三个点的特殊情况，根据题意来返回
			// 如果需要避免这种情况，可以先加入一个无穷大矩形对应的四个半平面，再求半平面交
			// 如果不足三个点，说明面积为 0
			return nil
		}

		// 补上首尾半平面的交点
		q[len(q)-1].p = q[len(q)-1].l.intersection(q[0].l)

		// 注：可以用三角剖分求半平面交的面积

		return q
	}

	// 点 p 是否在三角形 △abc 内
	// 一般是整数，浮点比较注意精度
	inTriangle := func(a, b, c, p vec) bool {
		pa, pb, pc := a.sub(p), b.sub(p), c.sub(p)
		return abs(b.sub(a).det(c.sub(a))) == abs(pa.det(pb))+abs(pb.det(pc))+abs(pc.det(pa))
	}

	// 判断点 p 是否在凸多边形 ps 内部 O(logn)
	// ps 需为逆时针顺序，ps[n-1] 无需等于 ps[0]
	// 【推荐】https://www.cnblogs.com/yym2013/p/3673616.html
	// https://cp-algorithms.com/geometry/point-in-convex-polygon.html
	// 其他 O(n) 方法 https://blog.csdn.net/WilliamSun0122/article/details/77994526
	//     判断点是否在所有边的左边（假设 ps 逆时针顺序）
	// EXTRA: 判断线段是否在凸多边形内：判断两端点是否均在凸多边形内即可
	inConvexPolygon := func(ps []vec, p vec) bool {
		o := ps[0]
		op := p.sub(o)
		// p 在凸多边形外：o-p 在最右射线 o-ps[1] 的右侧 or 在最左射线 o-ps[n-1] 的左侧
		// det: 正左负右
		if ps[1].sub(o).det(op) < 0 || ps[len(ps)-1].sub(o).det(op) > 0 { // 不允许点在边上则加上 =
			return false
		}
		// 二分找到一个点 ps[i]，使得 o-p 在 o-ps[i] 右侧或重合
		i := sort.Search(len(ps), func(mid int) bool { return ps[mid].sub(o).det(op) <= 0 })
		// p 是否在边 ps[i-1]-ps[i] 左侧或与这条边重合
		return ps[i].sub(ps[i-1]).det(p.sub(ps[i-1])) >= 0 // 不允许点在边上则改为 >
	}

	// 判断点 p 是否在多边形 ps 内部（不保证凸性）
	// https://oi-wiki.org/geometry/2d/#ray-casting-algorithm
	// https://leetcode.cn/contest/sf-tech/problems/uWWzsv/
	// http://acm.hdu.edu.cn/showproblem.php?pid=1756
	// 法一：射线法（光线投射算法 Ray casting algorithm）  奇内偶外
	// 由于转角法更方便，这里省略射线法的代码
	// 法二：转角法（绕数法、回转数法）
	// 【配图】https://blog.csdn.net/Form_/article/details/77855163
	// 代码参考《训练指南》
	// 这里统计绕数 Winding Number
	// 从 p 出发向右作射线，统计多边形穿过这条射线正反多少次
	// 【输入 ps 不要求是逆时针还是顺时针】
	inAnyPolygon := func(ps []vec, p vec) int {
		sign := func(x float64) int {
			if x < -eps {
				return -1
			}
			if x < eps {
				return 0
			}
			return 1
		}
		ps = append(ps, ps[0]) // 额外补一个点，方便枚举所有边
		wn := 0
		for i := 1; i < len(ps); i++ {
			p1, p2 := ps[i-1], ps[i]
			if p.onSeg(line{p1, p2}) {
				return -1 // 在边界上
			}
			// det: 正左负右
			k := sign(float64(p2.sub(p1).det(p.sub(p1)))) // 适配 int 和 float64
			d1 := sign(float64(p1.y - p.y))
			d2 := sign(float64(p2.y - p.y))
			if k > 0 && d1 <= 0 && d2 > 0 { // 逆时针穿过射线（p 需要在 p1-p2 左侧）
				wn++
			} else if k < 0 && d2 <= 0 && d1 > 0 { // 顺时针穿过射线（p 需要在 p1-p2 右侧）
				wn--
			}
		}
		if wn != 0 {
			return 1 // 在内部
		}
		return 0 // 在外部
	}

	// 判断任意两个多边形是否相离 O(n^2)
	// 属于不同多边形的任意两边都不相交，且一个多边形上的任意顶点都不被另一个多边形所包含
	// https://www.luogu.com.cn/problem/UVA10256

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

	// 矩形面积并（离散化+扫描线）
	// 见 segment_tree_rect.go

	// todo 矩形周长并

	// todo 三角形面积并
	// 扫描线算法
	// https://www.acwing.com/video/2218/
	// 模板题 https://www.luogu.com.cn/problem/P4406
	// https://www.luogu.com.cn/problem/P1222
	// https://www.luogu.com.cn/problem/P3219

	_ = []any{
		readVec, leftMostVec, rightMostVec,
		readPolygon, polygonArea,
		convexHull, convexHull2, convexHullHash, dynamicConvexHull,
		rotatingCalipers, convexHullPerimeter,
		halfPlanesIntersection,
		inTriangle, inConvexPolygon, inAnyPolygon,
		isRectangleAnyOrder, minAreaRect,
	}
}

/*
三维向量（点）
*/
type vec3 struct{ x, y, z int }

func (a vec3) less(b vec3) bool {
	return a.x < b.x || a.x == b.x && (a.y < b.y || a.y == b.y && a.z < b.z)
}

/*
三维直线（线段）
*/
type line3 struct{ p1, p2 vec3 }

// todo 计算几何三维入门 https://www.luogu.com.cn/blog/105254/ji-suan-ji-he-san-wei-ru-men
func vec3Collections() {
	var ps []vec3
	sort.Slice(ps, func(i, j int) bool {
		a, b := ps[i], ps[j]
		return a.x < b.x || a.x == b.x && (a.y < b.y || a.y == b.y && a.z < b.z)
	})

	// 三维凸包
	// todo 模板题 https://www.luogu.com.cn/problem/P4724
}

// 空间三角形相交

// 下面这些仅作为占位符表示，实际使用的时候复制上面的模板，类型改成 float64 同时 vecF 替换成 vec 等
// todo 怎么做更优雅一些呢？泛型好像不那么适配
type vecF struct{ x, y float64 }
type lineF struct{ p1, p2 vecF }
type vec3F struct{ x, y, z int }
type line3F struct{ p1, p2 vec3F }
type circleF struct {
	vecF
	r float64
}

func (vecF) add(vecF) (_ vecF)       { return }
func (vecF) sub(vecF) (_ vecF)       { return }
func (vecF) mul(float64) (_ vecF)    { return }
func (vecF) dot(vecF) (_ float64)    { return }
func (vecF) det(vecF) (_ float64)    { return }
func (vecF) len() (_ float64)        { return }
func (vecF) len2() (_ float64)       { return }
func (vecF) dis(vecF) (_ float64)    { return }
func (vecF) dis2(vecF) (_ float64)   { return }
func (vecF) polarAngle() (_ float64) { return }
func (lineF) vec() (_ vecF)          { return }
