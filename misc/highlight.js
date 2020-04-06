// ==UserScript==
// @name         Codeforces Highlight
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  highlight some important words
// @author       EndlessCheng
// @match        https://codeforces.com/*
// @match        https://codeforces.ml/*
// @match        https://codingcompetitions.withgoogle.com/*
// @match        https://atcoder.jp/*
// @grant        none
// ==/UserScript==

(function () {
    'use strict';

    function replaceAll(str, find, replace) {
        return str.replace(new RegExp(find, 'g'), replace);
    }

    const color = "#f25e6b";
    const words = [
        // 注意前者要包含后者
        " not ", "don't", "didn't", "doesn't", "can't", "n't", " no ",
        " and all", " or all",
        " and ", " or ",
        " any", " all ", "every", "both ",
        "exactly", "always",
        "unique", "distinct",
        "must", "only",
        "same", "different",
        "more",

        // "Note", "note",
        "minimize", "maximize", "minimum", "maximum", "minimal", "maximal", "smallest", "largest",
        " small ", " big ",
        "at least", "at most",
        "non-zero", "integers", "integer", "pairwise",
        "permutations", "permutation",
        "lowercase", "uppercase",
        "lexicographically", "lexicographical",
        "expected value",

        "Initially", "initially", "guaranteed",
        "modulo",

        "operations", "Operations", "operation", "Operation",
    ];

    const tags = ['p', 'li'];
    for (var ti = 0; ti < tags.length; ti++) {
        var pNodes = document.getElementsByTagName(tags[ti]);
        for (var i = 0; i < pNodes.length; i++) {
            var text = pNodes[i].innerHTML;
            for (var j = 0; j < words.length; j++) {
                text = replaceAll(text, words[j], "<span style='color: " + color + "'>" + words[j] + "</span>");
            }
            pNodes[i].innerHTML = text;
        }
    }
})();
