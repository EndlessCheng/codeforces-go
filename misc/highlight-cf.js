// ==UserScript==
// @name         高亮关键字 - CF
// @namespace    https://github.com/EndlessCheng
// @version      0.1
// @description  highlight some important words
// @author       灵茶山艾府
// @match        https://codeforces.com/contest/*
// @match        https://codeforces.com/group/*
// @match        https://codeforces.com/gym/*
// @match        https://codeforces.com/problemset/problem/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=codeforces.com
// ==/UserScript==

(function () {
    'use strict';

    // **题干**
    // 添加按钮
    $("div[class='problem-statement']").each(function () {
        $(this).parent().before(
            "<div class='html2md-panel'> <button class='html2mdButton highlightButton'>高亮</button> </div>"
        );
    });

    $(".highlightButton").click(function () {
        let target = $(this).parent().next().get(0);
        if (target.highlightContent) {
            return;
        }
        target.highlightContent = $(target).html();

        // console.log(target.highlightContent );

        target.highlightContent = target.highlightContent
            // 避免换行
            .replace(". ", ".") // 标题
            .replaceAll("Mr. ", "Mr.")
            .replaceAll("mr. ", "mr.")
            .replaceAll("Mrs. ", "Mrs.")
            .replaceAll("Ms. ", "Ms.")
            .replaceAll("Dr. ", "Dr.")
            .replaceAll("Co. ", "Co.")
            .replaceAll("Ltd. ", "Ltd.")
            .replaceAll("i. e. ", "i.e.")
            .replaceAll("i. e. ", "i.e.") // https://codeforces.com/contest/1535/problem/A
            .replaceAll("i.e. ", "i.e.")
            .replaceAll("I.e. ", "I.e.")
            .replaceAll("E. g. ", "E.g.") // https://codeforces.com/contest/1551/problem/E
            .replaceAll("E.g. ", "E.g.")
            .replaceAll(". ", ".</p><p>") // 加个换行（英文）
            .replaceAll("。", "。</p><p>") // 加个换行（中文）

        target.highlightContent = target.highlightContent.replace(".", ". ") // 复原标题

        String.prototype.highlight = function (substr, color) {
            return this.replaceAll(substr, "<span style='color: " + color + "'>" + substr + "</span>");
        };

        // 关键词高亮
        const color = "#49cc54"; // 49cc54  008dde
        // 高亮的词，一部分类似 Python 的关键字，另一部分是一些术语、修饰词之类
        // 注意若两个词有包含关系，把长的放前面
        const words = [
            // 描述
            "Initially", "initially", "guaranteed", "No matter", "no matter",
            "Not ", " not ", "don't", "didn't", "doesn't", "can't", "isn't", "aren't", "No ", " no ", "Non-", "non-", "without", "forbidden", "invalid", " nothing",
            "if and only if", "as long as",
            "Each ", " each ", // each 是个挺有趣的词，高亮它能帮助快速定位后面所描述的对象
            "every", " both ", " other ",
            "Also", "also",
            // " all ", "All ",
            "any number of", "Any number of",
            "equally", "equal", "same", "duplicate", "different", "unique", "distinct", "strictly", "overlapping", "overlap",
            "Only", "only", "just", "Exactly", "exactly", "always", "indeed", "precisely",
            "pairwise", "adjacent", "neighbour", "in a row", "consecutive", "continuous", "contiguous", "one after another", "disjoint", "as possible",
            "more than", "less than", "greater than",
            "except",
            // "must",

            // 求解
            "choose",
            "minimize", "maximize", "minimum", "maximum", "minimal", "maximal", "smallest", "largest", "shortest", "longest", "cheapest", "fastest",
            // " small", " big", " large", " few",
            "At least", "at least", "At most", "at most",

            // 特殊描述
            "substring", "subarray", // "subsequence",
            "Empty", " empty",
            "leading zero", "zero", "positive", "negative",
            "decreasing", "descending", "increasing", "ascending", "sorted",
            "permutation",
            "lowercase", "lower case", "uppercase", "upper case",
            "lexicographical", "palindrome",
            "undirected", "directed", "bidirectional", "direct",
            "independent",
            "expected value",
            " circle", " ring",
            "counterclockwise", "counter-clockwise", "clockwise",
            // "origin",
            "initial",
            "infinite",
            "leaf",
            "even integer",

            "没有", "不是", "不同", "不会", "互不",
            "相同",
            "独一无二", "唯一", "只有", "两个",
            "所有", "每", "任何", "任意", "或", "之一", "反之", "必须", "仅", "其他",
            "最小", "最大", "最少", "最多", "最短", "最长", "最早", "最晚", "最后", "第一", // todo regex 最...
            "至少", "至多", "恰好", "刚好",
            "非空", "连续", "子数组", "子区间", "区间", "子序列", "子字符串", "子串",
            "严格", /*"递增",*/ "递减", "升序", "降序", "字典序",
            "重复", "重新", "相邻",
            "小写", "大写", "回文",
            // "排列",
            "叶子", "叶节点",
            "单向", "双向",
            "本身",
            "独立",
            "返回",
            // "计算",
        ];

        for (let j = 0; j < words.length; j++) {
            target.highlightContent = target.highlightContent.highlight(words[j], color);
        }

        // 额外高亮
        const colorRed = "#ff0000";
        target.highlightContent = target.highlightContent
            .highlight("取模", colorRed)
            .highlight("取余", colorRed)
            .highlight("重复边", colorRed)
            .highlight("重边", colorRed)
            .highlight("3 seconds", colorRed)
            .highlight("4 seconds", colorRed)
            .highlight("5 seconds", colorRed)
            .highlight("6 seconds", colorRed)
            .highlight("7 seconds", colorRed)
            .highlight("8 seconds", colorRed)
            .highlight("9 seconds", colorRed)
            .highlight("10 seconds", colorRed);

        $(target).html(target.highlightContent);
    });
})();
