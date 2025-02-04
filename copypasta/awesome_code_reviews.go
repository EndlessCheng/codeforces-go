package copypasta

/* 有思路但是不好实现
https://codeforces.com/problemset/problem/1974/D 1400
https://codeforces.com/problemset/problem/1781/C 1600
https://codeforces.com/problemset/problem/1861/C 1600

*/

/* 编程技巧
https://codeforces.com/problemset/problem/691/B
点评：类似紫书例题 3-3 UVa401

*/

/* 非常漂亮的解法
https://atcoder.jp/contests/abc294/submissions/44905290
点评：巧妙引入「剩余长度」这一概念，简化比较逻辑。

https://atcoder.jp/contests/abc300/submissions/44909568
点评：与其滑窗，不如用前缀和的思想，视作有「k+i」次修改机会，算完后再减去「本不应该修改的长度」。
极大地简化了判断逻辑！妙哉！

*/
