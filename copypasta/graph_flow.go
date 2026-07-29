package copypasta

import (
	"container/heap"
	"math"
	"sort"
)

/* 网络流 Flow Network
https://en.wikipedia.org/wiki/Flow_network

Worst-Case Graphs for Maximum Flow Algorithms
https://codeforces.com/blog/entry/145343
https://github.com/leonard-weininger/worst-case-max-flow/

# 目录

最大流题目分类：
1. 收益-依赖-代价 -> 闭合子图/集合划分/最小割 Closed Subgraph / Graph Partitioning
2. 分配/匹配/填数 -> 二分图匹配/最大流 Bipartite Matching
3. 不相交/切断/阻碍 -> 拆点/最小割

费用流题目分类：
1. 带权二分图匹配 (Weighted Assignment)
2. K 路径覆盖与网格取数 (K-Path Cover & Grid Collection)
3. 凸函数/非线性代价的线性化 (Convex Cost Function Linearization)
4. 具有时间/状态属性的动态规划替代 (DP with Global Constraints)

如何区分最大流与费用流：
看目标函数：
- 只问「最多能匹配多少」-> 最大流
- 问「在匹配最多的前提下，最少花多少钱」-> 费用流
- 问「在总预算 C 以内，最多能匹配多少」-> 带预算的最大流（通常二分答案+费用流，或贪心费用流）
看贪心是否可行：
- 很多费用流问题看起来像贪心（比如每次选最便宜的）
- 如果当前的贪心选择会后悔（即阻碍了未来更大的收益），且无法通过简单的反悔堆解决，那么费用流就是系统化的反悔贪心

# 正文

网络流 24 题 https://loj.ac/p?tagIds=30 https://www.luogu.com.cn/problem/list?tag=332
线性规划与网络流 24 题 解题报告 https://byvoid.com/zhs/blog/lpf24-solution/

最大权闭合子图 “收益-依赖-代价” Maximum Weight Closure of a Graph
https://en.wikipedia.org/wiki/Closure_problem
  左部为我们需要决策选或不选（或者其他）的点，选择点 i 的收益为 earn[i]
  右部为点 i 的依赖，即选择点 i，也同时必须选点 g[i][j]
  选择右部的点 j 的代价为 cost[j]（或者说收益为 -cost[j]）
建图方式：
  超级源点 S 连 i，容量为 earn[i]
  j 连超集汇点 T，容量为 cost[j]
  g[i][j] 的容量为 inf（两部之间的边不能割）
利润 = sum(earn) - sum(没有选的 earn) - sum(选的点对应的 cost)
     = sum(earn) - 割掉没有选的 earn 以及左部选的点对应的右部点的 cost      割完后刚好可以把图分成两部分
最大权闭合子图 = sum(earn) - 最小割
对于左部，不割表示选，割表示不选，要反过来思考
没有思路？想一想【正难则反】：通过提前计算收益，把不选收益变成代价；通过提前计算代价（提前亏钱），把不选代价（不计入亏钱）变成收益
Solving Problems with Min Cut Max Flow Duality https://codeforces.com/blog/entry/136761 https://www.bilibili.com/video/BV1jt4y1t7pd/
https://codeforces.com/problemset/problem/311/E 2300 建图设计
https://codeforces.com/problemset/problem/1082/G 2400
- https://www.luogu.com.cn/problem/P4174 NOI06 最大获利
https://codeforces.com/problemset/problem/2026/E 2500
https://codeforces.com/problemset/problem/1666/K 3200
https://atcoder.jp/contests/abc326/tasks/abc326_g 2470=CF2579 经典建图设计
https://atcoder.jp/contests/arc085/tasks/arc085_c
https://atcoder.jp/contests/abc225/tasks/abc225_g
https://atcoder.jp/contests/abc259/tasks/abc259_g
https://atcoder.jp/contests/abc193/tasks/abc193_f
https://atcoder.jp/contests/arc142/tasks/arc142_e
https://www.luogu.com.cn/problem/P2762 https://loj.ac/p/6001 【网络流 24 题】太空飞行计划
https://qoj.ac/contest/1774/problem/9224

集合划分 Graph Partitioning
二元选择：每个元素有两个状态可选（例如：选/不选，属于A集合/属于B集合，黑色/白色）
代价/收益：选择某种状态会有基础收益；如果两个关联元素选择了“冲突”的状态（比如一个选A一个选B），会产生额外的“惩罚”或“代价”
目标：通常是求“最小的总代价”或“最大的总净收益”

二分图最大匹配 Maximum Bipartite Matching
找到看待问题的两种视角
超级源点连左部，右部连超级汇点，所有边的容量均为 1，最大流即为最大匹配
模板题 https://www.luogu.com.cn/problem/P3386 代码 https://www.luogu.com.cn/record/123020820
https://codeforces.com/problemset/problem/489/B 1200
https://codeforces.com/problemset/problem/78/E 2300
https://codeforces.com/problemset/problem/2026/E 2500 最大权闭合子图 / Hall 定理
https://codeforces.com/problemset/problem/628/F 2500 建图设计

带权二分图匹配
https://codeforces.com/problemset/problem/1437/C 1800

todo
 https://www.cnblogs.com/victorique/p/8560656.html
 https://blog.bill.moe/network-flow-models/
 NOI 一轮复习 I：二分图网络流 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-i-er-fen-tu-wang-lao-liu
 2016 国家集训队论文《网络流的一些建模方法》姜志豪 https://github.com/enkerewpo/OI-Public-Library/blob/master/IOI%E4%B8%AD%E5%9B%BD%E5%9B%BD%E5%AE%B6%E5%80%99%E9%80%89%E9%98%9F%E8%AE%BA%E6%96%87/%E5%9B%BD%E5%AE%B6%E9%9B%86%E8%AE%AD%E9%98%9F2016%E8%AE%BA%E6%96%87%E9%9B%86.pdf

todo 题单 https://www.zybuluo.com/xzyxzy/note/992041
 网络流从入门到入土 #1 https://www.luogu.com.cn/training/12097#problems
 网络流从入门到入土 #2 https://www.luogu.com.cn/training/12098#problems
 网络流从入门到入土 #3 https://www.luogu.com.cn/training/12099#problems
 网络流建模经典题 https://www.luogu.com.cn/training/1230#problems
 网络流经典题目 https://www.luogu.com.cn/training/3144#problems

Max-Flow in almost linear time https://codeforces.com/blog/entry/100510

CF Tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=flows
*/

/* 最大流·建模·转换
Max-Flow，简称 MF
https://en.wikipedia.org/wiki/Maximum_flow

可视化 https://visualgo.net/zh/maxflow
选择左下的图示 - CS4234 MF Demo 或者 CP4 8.15*，然后选择左下的 Dinic - 前进

最大流·建模·转换
https://www.luogu.com.cn/problem/P2891 http://poj.org/problem?id=3281
https://loj.ac/p/6005 https://www.luogu.com.cn/problem/P2766 【网络流 24 题】最长不降子序列
- 注意这题用到了操纵超级源点的技巧：容量限制与解除容量限制
NWERC07 B https://codeforces.com/gym/100723 http://poj.org/problem?id=3498 UVa12125 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=243&page=show_problem&problem=3277
https://codeforces.com/problemset/problem/1360/G 1900 网格模型
https://codeforces.com/problemset/problem/498/C 2100 建模题，推荐
https://codeforces.com/problemset/problem/546/E 2100
https://codeforces.com/problemset/problem/653/D 2200 转换
https://codeforces.com/problemset/problem/1913/E 2400 todo
https://atcoder.jp/contests/arc085/tasks/arc085_c 转换 todo
http://poj.org/problem?id=1149 转换

把点拆为入点和出点（v 和 v+n），可把点上的约束变成边上的约束

顶点上有容量
将顶点拆成两个（入顶点 x 和出顶点 y），入点向 x 连边，y 向出点连边，x 向 y 连边，容量为顶点的容量

无向图
视作两条容量均为 cap 的有向边（具体实现见下面代码中 addEdge 的注释）

多源汇最大流
建立超级源点 S 和超级汇点 T，S 向所有源点连边，所有汇点向 T 连边，每条边的容量为 inf 或对应源汇的容量限制
https://www.acwing.com/problem/content/2236/

只能经过这条边一次 ⇔ 容量为 1
http://poj.org/problem?id=2455 https://www.acwing.com/problem/content/2279/

上下界可行流·总结
https://oi-wiki.org/graph/flow/bound/
https://www.acwing.com/solution/content/17067/
https://zhuanlan.zhihu.com/p/324507636
todo 题单 https://www.luogu.com.cn/training/8462

无源汇上下界可行流（循环流）
假设存在一个流量守恒的解 f，通过将每条边的流量减去 low，得到一个新图的流，但其不一定满足流量守恒
对于每个顶点 v，记 d(v) = ∑lowIn(v) - ∑lowOut(v)
- 若 d(v) > 0，说明流入减去的更多，则需将 v 的流入量增加 d(v)，这可以通过新建超级源点 S，并增加 S->v，容量为 d(v) 的边做到
- 若 d(v) < 0，说明流出减去的更多，则需将 v 的流出量增加 d(v)，这可以通过新建超级汇点 T，并增加 v->T，容量为 -d(v) 的边做到
跑从 S 到 T 的最大流，若满流（即最大流等于从 S 出发的容量之和），则说明可以让新图的流量守恒，从而说明原图存在可行流 f，其每条边的流量为 low 加上新图中每条边的流量；若不满流则无解
模板题 https://loj.ac/p/115 https://www.acwing.com/problem/content/2190/

有源汇上下界可行流
从汇点向源点连一条容量为 inf 的边，即转换成了无源汇上下界可行流

有源汇上下界最大流
1. 跑一遍有源汇上下界可行流，若有解，记此时源点到汇点的流量为 f1（通过汇点向源点的反向边的流量得到）
2. 删去汇点到源点的边（或将其容量置为 0，具体实现时可以将汇点->源点边最后加入，或者使用指针记录该边及其反向边）
3. 在残余网络上继续增广，记额外的最大流为 f2，那么答案即为 f1+f2
模板题 https://loj.ac/p/116 https://www.luogu.com.cn/problem/P5192

有源汇上下界最小流
将上面第 3 步改成退流，即减去残余网络上从汇点到源点的最大流
模板题 https://loj.ac/p/117 https://www.luogu.com.cn/problem/P4843

分层图
注意：可以在原图的基础上添加边/增加容量，然后继续寻找增广路增广
【网络流 24 题】星际转移 https://loj.ac/p/6015 https://www.luogu.com.cn/problem/P2754

关键边
关键边 v-w 需满足，在跑完最大流后：
1. 这条边的流量等于其容量
2. 在残余网络上，从源点可以到达 v，从 w 可以到达汇点（即从汇点顺着反向边可以到达 w）
http://poj.org/problem?id=3204 https://www.acwing.com/problem/content/2238/
具体实现见下面代码中的 EXTRA
*/

/* 最小割·建模·转换
https://en.wikipedia.org/wiki/Max-flow_min-cut_theorem
最大流等于最小割的证明 https://seineo.github.io/%E5%9B%BE%E8%AE%BA%EF%BC%9A%E6%9C%80%E5%A4%A7%E6%B5%81%E6%9C%80%E5%B0%8F%E5%89%B2%E8%AF%A6%E8%A7%A3.html
最小割模型汇总 https://blog.csdn.net/qq_35649707/article/details/77482691
下面的 topic 参考胡伯涛《最小割模型在信息学竞赛中的应用》（PDF 见 https://github.com/EndlessCheng/cp-pdf）

如何输出最小割（只需要求一个解）
1. 求最大流
2. 从源点出发在残余网络上 DFS，标记所有能够到达的点
3. 遍历原边集 edges，若其中一端有标记，另一端没有标记，则这条边为最小割上的边

技巧 1：用容量为 inf 的边来防止割断
技巧 2：给边权加上很大的数来约束删除次数
https://codeforces.com/problemset/problem/700/C 2600 推荐

建模·转换
https://atcoder.jp/contests/arc074/tasks/arc074_d 网格图标准题
https://www.luogu.com.cn/problem/P1345
https://atcoder.jp/contests/abc326/tasks/abc326_g
https://www.acwing.com/problem/content/2282/
平均边权最小 https://www.acwing.com/problem/content/2281/
点连通度 SEERC04 F https://codeforces.com/gym/101461 http://poj.org/problem?id=1966 UVa1660 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4535
   https://en.wikipedia.org/wiki/Connectivity_(graph_theory)
   https://en.wikipedia.org/wiki/Menger%27s_theorem
LCP38/21春·战队赛F https://leetcode.cn/problems/7rLGCR/
todo https://atcoder.jp/contests/arc085/tasks/arc085_c

todo 最小割必经边？

最大密度子图 Maximum Density Subgraph
https://en.wikipedia.org/wiki/Dense_subgraph
参考 https://www.luogu.com.cn/problem/solution/UVA1389
二分上下界：最小密度为 1/n，最大密度为 m
二分精度：任意两个密度不同的子图，其密度差 >= 1/n^2
todo NEERC06 H https://codeforces.com/gym/100287 https://codeforces.com/gym/100532 http://poj.org/problem?id=3155 UVa1389 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4135

二分图最小点权覆盖集 Minimum Weight Vertex Covering Set (MinWVCS) in a Bipartite Graph
二分图最大点权独立集 Maximum Weight Vertex Independent Set (MaxWVIS) in a Bipartite Graph
建立一个源 s，向 X 部每个点连边；建立一个汇 t，从 Y 部每个点向汇 t 连边，把二分图中的边看成是有向的，
则任意一条从 s 到 t 的路径，一定具有 s->v->w->t 的形式（v∈X, w∈Y）。
割的性质是不存在一条从 s 到 t 的路径。故路径上的三条边 s-v, v-w, w-t 中至少有一条边在割中。
若人为地令 v-w 不可能在最小割中，即令其容量为正无限，
可将条件简化为 s-v, w-t 中至少有一条边在最小割中，这正好与点覆盖集限制条件的形式相符（边的两端点中至少一个在覆盖集内），
而目标是最小化点权之和，这恰好也是最小割的优化目标。
对于最大点权独立集，其等价于点权之和减去最小点权覆盖集。
【网络流 24 题】骑士共存 https://loj.ac/p/6226 https://www.luogu.com.cn/problem/P3355
todo https://codeforces.com/contest/808/problem/F
NEERC03 D https://codeforces.com/gym/100725 https://codeforces.com/gym/101651 http://poj.org/problem?id=2125
黑白染色转化成二分图 https://www.acwing.com/problem/content/2328/
todo https://atcoder.jp/contests/abc285/tasks/abc285_g

最小割的可行边和必须边（所有最小割集的并集和交集）
跑最大流，然后求整个残余网络的 SCC，则有：
- 可行边：两端不在一个 SCC 内，即不存在另一条从 v 到 w 的路径
- 必须边：一端在 S 的 SCC 内,另一端在 T 的 SCC 内
AHOI09 https://www.luogu.com.cn/problem/P4126
*/

/* 费用流·建模·转换
https://en.wikipedia.org/wiki/Minimum-cost_flow_problem MCFP
https://en.wikipedia.org/wiki/Assignment_problem
https://en.wikipedia.org/wiki/Network_simplex_algorithm

NOTE: 对于修改容量的情况，由于 EK 是基于最短路的贪心算法，不能像最大流那样直接在残余网络上继续跑，必须重新建图重新跑 EK
todo https://codeforces.com/problemset/problem/362/E

建模·转换
从源点到集合 A 中各点连边，容量为 1，费用为 0
从集合 B 中各点到汇点连边，容量为 1，费用为 0
集合 A 和 B 两两连边，容量为 inf（或者题目指定），费用为 F(Ai,Bj)（题目指定）
这样跑 MCMF 得到的结果是匹配全部 A（或 B）的最小花费
代表题目 https://codeforces.com/problemset/problem/237/E 2000 | 代码 https://codeforces.com/problemset/submission/237/241222973
https://codeforces.com/problemset/problem/1437/C 1800
LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/ 2392
【网络流 24 题】运输问题 https://loj.ac/p/6011 https://www.luogu.com.cn/problem/P4015
【网络流 24 题】数字梯形 https://loj.ac/p/6010 https://www.luogu.com.cn/problem/P4013
【网络流 24 题】深海机器人 https://loj.ac/p/6224 https://www.luogu.com.cn/problem/P4012
k 取方格数 https://www.luogu.com.cn/problem/P2045 http://poj.org/problem?id=3422
    关键技巧：拆点时，从入点向出点连两条边，第一条边容量为 1，费用为点权，第二条边容量为 k-1，费用为 0
    这表示第一次经过该点时，可以把数取走，之后再经过时就不再计算
【网络流 24 题】餐巾计划 https://loj.ac/p/6008 https://www.luogu.com.cn/problem/P1251

最大费用
将每条边的费用反向，答案即为 -MCMF

无源汇上下界最小费用可行流
建图和上面的「无源汇上下界可行流」一样
NOI08 志愿者招募 https://www.luogu.com.cn/problem/P3980（也可以用线性规划做）
- 由于没有上界，建图的时候可以不用减去下界
- 把每天的人数要求看成是边的流量下界（从 i 天向 i+1 天连边）
- 由于要满足流量守恒，对于每个人 i，需要从结束日期向开始日期连边，容量为 inf，费用为 ci。这相当于每个人在流网络的一单位的流量流过了一个环
- 代码实现 https://www.luogu.com.cn/record/56398769
AHOI14/JSOI14 支线剧情 https://www.luogu.com.cn/problem/P4043
-「看完所有剧情」可以转换成每条边的流量下界为 1，容量为 inf，费用为过剧情花费的时间
-「开始新的游戏」可以转换成每个点向点 1 连边，容量为 inf，费用为 0
- 代码实现 https://www.luogu.com.cn/record/56402617

流通问题 circulation problem
最小费用流通问题 minimum-cost-circulation problem
https://en.wikipedia.org/wiki/Circulation_problem
The circulation problem and its variants are a generalisation of network flow problems,
with the added constraint of a lower bound on edge flows,
and with flow conservation also being required for the source and sink (i.e. there are no special nodes).
《算法导论》思考题 29-5
todo https://codeforces.com/contest/1455/problem/E
 https://codeforces.com/blog/entry/85186?#comment-728533
*/

/* 网络流建模技巧/转换技巧
todo 整合到其它 blocks
todo 重新看一下挑战

标准建模（指派问题）：
	http://poj.org/problem?id=2175
	http://poj.org/problem?id=3686
边容量减少：
	若 flow<=cap' 则最大流不变；若 flow>cap' 需要将多出的流退回去 todo
    最小割+退流 https://www.luogu.com.cn/problem/P3308
流量任意：
	todo
容量为负数：
	todo
费用为负数：
	todo 挑战:228
求最小割划分成两个集合：
	Dual Core CPU http://poj.org/problem?id=3469
无重复边的往返最短路：
	http://poj.org/problem?id=2135
	转换成流量为 2 的最小费用流
点边转换
   将点拆为入点和出点（v 和 v+n），即可把点的属性变成边的属性，从而方便应用最大流、最小割等算法
   将边的中间加一个节点，把边的属性体现在中间的点上
上下界费用流
	对每条边新增一条边 e'
	e.cap-=minCap
	e'.cap=minCap
	e'.cost=e.cost-M // 一个足够大的常数
	跑完 MCMF 后加上 M*∑minCap

Disjoint paths
Edge-disjoint paths: It turns out that the maximum number of edge-disjoint paths equals the maximum flow of the graph, assuming that the capacity of each edge is one.
Node-disjoint paths: 拆点法

路径覆盖问题 Path cover + 打印
todo https://zhuanlan.zhihu.com/p/125759333
todo Competitive Programmer’s Handbook Ch.20
todo 线性规划与网络流 24 题 - 最小路径覆盖问题 https://byvoid.com/zhs/blog/lpf24-3/

给一 DAG，求它的最大反链大小、一组最大反链构造以及所有最大反链的并
https://yhx-12243.github.io/OI-transit/records/lydsy1143%3Blg4298.html
*/

// 最大流 Dinic's algorithm O(n^2 * m)  二分图上为 O(m√n)
// https://en.wikipedia.org/wiki/Dinic%27s_algorithm
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// [Tutorial] My way of understanding Dinitz's ("Dinic's") algorithm https://codeforces.com/blog/entry/104960
// https://www.bilibili.com/video/BV1j64y1R7yK/
// 时间复杂度证明 https://www.zhihu.com/question/34374412
// - 一次增广令至少一条边饱和（cap = 0），所以有至多 O(m) 条增广路，每条增广路的长度为 O(n)，所以多路增广的时间是 O(nm)
// - 由于每次 BFS 都会使 d[end] 变大，至多变大 O(n) 次，所以总的时间复杂度为 O(n) * O(nm) = O(n^2 * m)
// 关于二分图上的时间复杂度，见 https://www.cnblogs.com/Itst/p/12556871.html
//
// 模板题 https://www.luogu.com.cn/problem/P3376
//       https://www.luogu.com.cn/problem/P2740
// https://codeforces.com/problemset/problem/498/C 2100 最大流 建模
// https://codeforces.com/problemset/problem/546/E 2100 模板题 输出具体方案
// https://atcoder.jp/contests/practice2/tasks/practice2_d 骨牌/瓷砖铺设 输出具体方案
func (*graph) maxFlowDinic(n, st, end int, edges [][]int, a, b []int) int {
	type neighbor struct{ to, rid, cap, eid int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, eid int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, eid})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -1}) // 无向图上 0 换成 cap
	}
	for i, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap, i)
	}

	{
		// 二分图最大匹配的建图（忽略上面的代码）
		st := len(a) + len(b)
		end := st + 1
		type neighbor struct{ to, rid, cap int }
		g := make([][]neighbor, end+1)
		addEdge := func(from, to, cap int) {
			g[from] = append(g[from], neighbor{to, len(g[to]), cap})
			g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
		}

		// 超级源点连左部，右部连超级汇点，所有边的容量均为 1，最大流即为最大匹配
		for i, v := range a {
			for j, w := range b {
				// 和题目有关，满足该约束即可匹配 a[i] 和 b[j]
				if v+w < 100 {
					// 一般 cap=math.MaxInt 表示无限制/不允许割（下面我们对 st end 做了限制，所以一般 cap=math.MaxInt）
					// cap=1 表示一对一
					// 如果题目约束 a[i] 只能匹配 x 个 b[j]，那么 cap=x
					addEdge(i, j+len(a), math.MaxInt)
				}
			}
			// cap=1 表示每个 a[i] 只能选一次（匹配一次）
			// 如果题目允许一对多，比如一对二，把 1 改成 2
			addEdge(st, i, 1)
		}
		for j := range b {
			// cap=1 表示 b[j] 只能被一个 a[i] 独占（匹配）
			// 如果题目允许多对一，比如二对一，把 1 改成 2
			addEdge(j+len(a), end, 1)
		}

		// 算完最大流后，如果要输出具体方案，可以遍历左部 -> 右部的边，cap == 0 的边就是在最大匹配中的边
	}

	dis := make([]int, len(g))
	bfs := func() bool {
		clear(dis) // dis[i] = 0 表示没有访问过
		dis[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && dis[w] == 0 {
					dis[w] = dis[v] + 1
					q = append(q, w)
				}
			}
		}
		return dis[end] > 0
	}
	// 当前弧，在其之前的边已经没有用了，避免多次检查没有用的边
	iter := make([]int, len(g))
	// 寻找增广路（多路增广）
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == end {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			w := e.to
			if e.cap > 0 && dis[w] > dis[v] {
				f := dfs(w, min(totalFlow-curFlow, e.cap))
				if f == 0 {
					continue
				}
				e.cap -= f
				g[w][e.rid].cap += f
				curFlow += f
				if curFlow == totalFlow {
					break
				}
			}
		}
		return
	}
	maxFlow := 0
	for bfs() {
		clear(iter)
		maxFlow += dfs(st, math.MaxInt)
	}

	// EXTRA: 容量复原（不存原始容量的写法）
	for _, es := range g {
		for i, e := range es {
			if e.eid >= 0 { // 正向边
				es[i].cap += g[e.to][e.rid].cap
				g[e.to][e.rid].cap = 0
			}
		}
	}

	// EXTRA: 求流的分配方案（即反向边上的 cap）
	// https://codeforces.com/problemset/problem/546/E 2100
	// https://loj.ac/p/115 https://www.acwing.com/problem/content/2190/
	ans := make([]int, len(edges))
	for _, es := range g { // v
		for _, e := range es {
			w, i := e.to, e.eid
			if i >= 0 { // 正向边
				ans[i] = g[w][e.rid].cap
			}
		}
	}

	// EXTRA: 求关键边（扩容后可以增加最大流的边）的数量
	// 关键边 v-w 需满足，在跑完最大流后：
	// 1. 这条边的流量等于其容量
	// 2. 在残余网络上，从源点可以到达 v，从 w 可以到达汇点（即从汇点顺着反向边可以到达 w）
	// http://poj.org/problem?id=3204 https://www.acwing.com/problem/content/2238/
	{
		// 在残余网络上跑 DFS，看看哪些点能从源点和汇点访问到（从汇点出发的要判断反向边的流量）
		vis1 := make([]bool, len(g))
		var dfs1 func(int)
		dfs1 = func(v int) {
			vis1[v] = true
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && !vis1[w] {
					dfs1(w)
				}
			}
		}
		dfs1(st)

		vis2 := make([]bool, len(g))
		var dfs2 func(int)
		dfs2 = func(v int) {
			vis2[v] = true
			for _, e := range g[v] {
				if w := e.to; !vis2[w] && g[w][e.rid].cap > 0 {
					dfs2(w)
				}
			}
		}
		dfs2(end)

		ans := 0
		for v, es := range g {
			if !vis1[v] {
				continue
			}
			for _, e := range es {
				// 原图的边，流量为 0（说明该边满流），且边的两端点能分别从源汇访问到
				if e.eid >= 0 && e.cap == 0 && vis2[e.to] {
					ans++
				}
			}
		}
	}

	return maxFlow
}

// ISAP, Improved Shortest Augmenting Path O(n^2 * m)
// https://oi-wiki.org/graph/flow/max-flow/#isap
// https://www.renfei.org/blog/isap.html
// 测试了一下性能和 Dinic 差不多
func (*graph) maxFlowISAP(n, st, end int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for _, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap)
	}

	// 计算从汇点 end 出发的距离
	d := make([]int, n)
	for i := range d {
		d[i] = -1
	}
	d[end] = 0
	cd := make([]int, n+1) // 注意有 d[i] == n 的情况
	q := []int{end}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		cd[d[v]]++
		for _, e := range g[v] {
			if w := e.to; d[w] < 0 {
				d[w] = d[v] + 1
				q = append(q, w)
			}
		}
	}
	if d[st] < 0 {
		return -1
	}

	// 寻找增广路
	const inf int = 1e18
	maxFlow := 0
	iter := make([]int, n)
	type pair struct{ v, i int }
	fa := make([]pair, n)
o:
	for v := st; d[st] < n; {
		if v == end {
			minF := inf
			for v := end; v != st; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := end; v != st; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			v = st
		}
		for i := iter[v]; i < len(g[v]); i++ {
			e := g[v][i]
			if w := e.to; e.cap > 0 && d[w] < d[v] {
				fa[w] = pair{v, i}
				iter[v] = i
				v = w
				continue o
			}
		}
		if cd[d[v]] == 1 {
			break // gap 优化
		}
		cd[d[v]]--
		minD := n - 1
		for _, e := range g[v] {
			if e.cap > 0 && d[e.to] < minD {
				minD = d[e.to]
			}
		}
		d[v] = minD + 1
		cd[d[v]]++
		iter[v] = 0
		if v != st {
			v = fa[v].v
		}
	}
	return maxFlow
}

// 最高标号预流推进 (HLPP, High Level Preflow Push)   O(n^2 * √m)
// 注：虽然在复杂度上比增广路方法进步很多，但是预流推进算法复杂度的上界是比较紧的，因此有时差距并不会很大
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm#Highest_label_selection_rule
// https://oi-wiki.org/graph/flow/max-flow/#hlpp
// https://www.luogu.com.cn/blog/ONE-PIECE/jiu-ji-di-zui-tai-liu-suan-fa-isap-yu-hlpp
// 模板题 https://loj.ac/p/127 https://www.luogu.com.cn/problem/P4722
// todo deque 优化 + 全局重贴标签等 https://www.luogu.com.cn/problem/solution/P4722
type hlppHeap struct {
	sort.IntSlice
	d []int
}

func (h hlppHeap) Less(i, j int) bool { return h.d[h.IntSlice[i]] > h.d[h.IntSlice[j]] } // 处于堆中的节点的 d 值不会改变，所以可以直接比较
func (h *hlppHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hlppHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hlppHeap) push(v int)        { heap.Push(h, v) }
func (h *hlppHeap) pop() int          { return heap.Pop(h).(int) }

func (*graph) maxFlowHLPP(n, st, end int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for _, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap)
	}

	// 计算从汇点 end 出发的距离
	d := make([]int, n)
	for i := range d {
		d[i] = -1
	}
	d[end] = 0
	cd := make([]int, 2*n)
	_q := []int{end}
	for len(_q) > 0 {
		v := _q[0]
		_q = _q[1:]
		cd[d[v]]++
		for _, e := range g[v] {
			if w := e.to; d[w] < 0 {
				d[w] = d[v] + 1
				_q = append(_q, w)
			}
		}
	}
	if d[st] < 0 {
		return -1
	}
	d[st] = n

	exFlow := make([]int, n)
	q := hlppHeap{d: d}
	inQ := make([]bool, n)
	push := func(v, f int, e *neighbor) {
		w := e.to
		e.cap -= f
		g[w][e.rid].cap += f
		exFlow[v] -= f
		exFlow[w] += f
		if w != st && w != end && !inQ[w] {
			q.push(w)
			inQ[w] = true
		}
	}
	// 将源点的所有边都满流地推送出去
	for i := range g[st] {
		if e := &g[st][i]; e.cap > 0 {
			push(st, e.cap, e)
		}
	}
	for len(q.IntSlice) > 0 {
		v := q.pop()
		inQ[v] = false
	o:
		for {
			for i := range g[v] {
				if e := &g[v][i]; e.cap > 0 && d[e.to] < d[v] {
					push(v, min(e.cap, exFlow[v]), e)
					if exFlow[v] == 0 {
						break o
					}
				}
			}
			dv := d[v]
			cd[dv]--
			if cd[dv] == 0 { // gap 优化
				for i, h := range d {
					if i != st && i != end && dv < h && h <= n {
						d[i] = n + 1 // 超过 n，从而尽快将流量推回 st
					}
				}
			}
			// relabel
			minD := int(1e9)
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] < minD {
					minD = d[w]
				}
			}
			d[v] = minD + 1
			cd[d[v]]++
		}
	}
	return exFlow[end]
}

// 无向图全局最小割
// Stoer-Wagner 算法 O(nm+n^2logn)
// https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GlobalMincut.java.html
// todo 模板题 https://www.luogu.com.cn/problem/P5632 http://poj.org/problem?id=2914
func (*graph) minimumCutStoerWagner(dist [][]int) int {
	panic("todo")
}

// 最小费用流 MCFP
// 最小费用最大流 MCMF（即满流时的费用）
// 将 Edmonds-Karp 中的 BFS 改成 SPFA O(fnm) 或 Dijkstra O(fmlogn)
// 要求初始网络中无负权圈
// 性能对比（洛谷 P3381，由于数据不强所以 SPFA 很快）：SPFA 1.05s(max 365ms)   Dijkstra 1.91s(max 688ms)
// https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm
// https://oi-wiki.org/graph/flow/min-cost/
// https://cp-algorithms.com/graph/min_cost_flow.html
// 最小费用流的不完全算法博物馆 https://www.luogu.com.cn/blog/Atalod/zui-xiao-fei-yong-liu-di-fou-wan-quan-suan-fa-bo-wu-guan
//
// 模板题 https://www.luogu.com.cn/problem/P3381
// https://codeforces.com/problemset/problem/237/E 2000
//
// 常见建模方式（下面代码按照这种建模写的）
// 建模的时候，一般可以理解成在一个矩阵 a 上，每行每列至多选 K 个数，问所选数字之和的最小值
// 创建一个（完全）二分图，左部为行，右部为列
// - 行 -> 列，容量为 1，费用为 grid[i][j]（如果同一个元素可以重复选，则容量为 inf；如果至多选 mx[i][j] 个，则容量为 mx[i][j]）
// - 超级源点 S -> 行，容量为 K，费用为 0（K 表示每行至多选 K 个数）
// - 列 -> 超级汇点 T，容量为 K，费用为 0（K 表示每列至多选 K 个数）
// 如果要限制总共至多选 lim 个元素，可以在超级源点前面再加一个节点 S0，连到超级源点，容量为 lim，费用为 0（相当于超级源点的流出量至多为 lim）
//     如果满流，则表示恰好选了 lim 个元素
//
// https://codeforces.com/problemset/problem/1426/E 1800 完全二分图 + 多对多
// https://codeforces.com/problemset/problem/1437/C 1800 完全二分图 + 一对一
// https://codeforces.com/problemset/problem/1525/D 1800 完全二分图 + 一对一
// https://codeforces.com/problemset/problem/802/N 2400 时间约束：从第 i 天连向第 i+1 天，容量 inf（延后无限制），费用 0（等待不花钱）
// https://codeforces.com/problemset/problem/802/O 2900 模拟费用流/反悔贪心
// https://projecteuler.net/problem=345 完全二分图 + 一对一
// 完全二分图 + 一对一 LC3376 https://leetcode.cn/problems/minimum-time-to-break-locks-i/
// 完全二分图 + 一对多 LC2850 https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/
// 完全二分图 + 至多选 k=3 个数 LC3257 https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/
// 多对一 LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/
// 二分图 + 稀疏矩阵 LC3276 https://leetcode.cn/problems/select-cells-in-grid-with-maximum-score/
// https://codeforces.com/problemset/problem/1107/F 2600
// 最大费用流 + 输出具体方案 https://atcoder.jp/contests/practice2/tasks/practice2_e
// - 所有 cost 取相反数，变成最小费用流
// - 注意可能某次增广算出来的最短路是正数，导致满流的时候不是最大费用
// - 所以要先跑一遍 MCMF，记录增广过程中最大的费用对应的流 tarFlow；然后再跑一遍 MCMF，限制总共至多选 tarFlow 个元素
// - 最终答案为 -minCost
// - 代码 https://atcoder.jp/contests/practice2/submissions/64239460
func (*graph) minCostFlowSPFA(a [][]int) (int, int) {
	n := len(a)
	m := len(a[0])
	st := n + m
	end := st + 1

	// rid 为反向边在邻接表中的下标
	type nb struct{ to, rid, cap, cost int } // 如果输入的是 edges，可以额外记录边的下标
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
	}
	for i, row := range a {
		for j, v := range row {
			// 如果是多对多，改 cap（一般为 math.MaxInt）
			// 如果求最大，cost 改成 -v
			addEdge(i, n+j, 1, v) 
		}
		// 如果是一对多，改 cap=这个点可以选择的最大次数/可用次数
		addEdge(st, i, 1, 0) 
		// 注：特别地，如果这一行的所有 v 都相同，可以把 st->i 的 cost 改成 v，i->n+j 的 cost 改成 0
	}
	for j := range a[0] {
		// 如果是多对一，改 cap=这个点能被匹配的最大次数/可用次数
		addEdge(n+j, end, 1, 0) 
	}
	// 如果要限制至多选 k 个元素，把超级源点分裂成两个点，加 cap，相当于在点上限制流量
	//addEdge(st, st2, k, 0) // 上面 addEdge 中的 st 改成 st2

	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[st] = 0
		inQ[st] = true
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[end] < math.MaxInt
	}

	maxFlow := 0 // 可选
	minCost := 0
	for spfa() {
		// 沿 st-end 的最短路尽量增广
		// 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
		minF := math.MaxInt
		for v := end; v != st; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		maxFlow += minF
		minCost += dis[end] * minF
	}

	// 输出具体方案
	for _, to := range g[:n] {
		for _, e := range to[:m] {
			if e.cap == 0 { // 或者小于其初始值
				// 选了 a[i][j]
			}
		}
	}

	return maxFlow, minCost
}

// 基于原始对偶方法 (primal-dual method)
// https://blog.xehoth.cc/DurationPlan-Primal-Dual/
func (*graph) minCostFlowDijkstra(n, st, end, flowLimit int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for _, e := range edges {
		v, w, edgeCap, edgeCost := e[0], e[1], e[2], e[3]
		addEdge(v, w, edgeCap, edgeCost)
	}

	h := make([]int, len(g)) // 顶点的势
	dist := make([]int, len(g))
	type pair struct{ v, i int }
	fa := make([]pair, len(g))
	dijkstra := func() bool {
		const _inf int = 1e18
		for i := range dist {
			dist[i] = _inf
		}
		dist[st] = 0
		q := dijkstraHeap{{st, 0}}
		for len(q) > 0 {
			p := q.pop()
			v := p.v
			if p.dis > dist[v] {
				continue
			}
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost + h[v] - h[w]; newD < dist[w] {
					dist[w] = newD
					fa[w] = pair{v, i}
					q.push(dijkstraPair{w, newD})
				}
			}
		}
		return dist[end] < _inf
	}
	minCost := 0
	for flowLimit > 0 && dijkstra() {
		for i, d := range dist {
			h[i] += d
		}
		minF := flowLimit // inf
		for v := end; v != st; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minF {
				minF = c
			}
			v = p.v
		}
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		flowLimit -= minF        // maxFlow += minF
		minCost += h[end] * minF // 注意这里是 h 不是 dist
	}
	if flowLimit > 0 {
		return -1
	}
	return minCost
}

// todo 基于 Capacity Scaling 的弱多项式复杂度最小费用流算法 
// https://ouuan.github.io/post/%E5%9F%BA%E4%BA%8E-capacity-scaling-%E7%9A%84%E5%BC%B1%E5%A4%9A%E9%A1%B9%E5%BC%8F%E5%A4%8D%E6%9D%82%E5%BA%A6%E6%9C%80%E5%B0%8F%E8%B4%B9%E7%94%A8%E6%B5%81%E7%AE%97%E6%B3%95/

// ZKW 费用流
// https://artofproblemsolving.com/community/c1368h1020435
