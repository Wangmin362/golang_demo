package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/

// 解题思路：

func findAnagrams(s string, p string) []int {
	pMap := ['z' - 'a' + 1]int{}
	for _, pChar := range p {
		pMap[pChar-'a'] += 1
	}

	var res []int
	for i := 0; i <= len(s)-len(p); i++ {
		tmp := pMap
		isValid := true
		for j := i; j < i+len(p); j++ {
			cnt := tmp[s[j]-'a']
			if cnt == 0 {
				isValid = false
				break
			} else {
				tmp[s[j]-'a'] = cnt - 1
			}
		}
		if isValid {
			res = append(res, i)
		}
	}

	return res
}

func findAnagrams01(s string, p string) []int {
	freCnt := func(str string) [26]int {
		cnt := [26]int{}
		for idx := range str {
			cnt[str[idx]-'a']++
		}
		return cnt
	}
	if len(s) < len(p) {
		return nil
	}
	pCnt := freCnt(p)
	slow, fast := 0, len(p)-1
	var res []int
	for fast < len(s) && fast-slow+1 == len(p) {
		sCnt := freCnt(s[slow : fast+1])
		if reflect.DeepEqual(pCnt, sCnt) {
			res = append(res, slow)
		}
		slow++
		fast++
	}
	return res
}

func TestFindAnagrams(t *testing.T) {
	var testdata = []struct {
		s      string
		p      string
		expect []int
	}{
		{s: "cbaebabacd", p: "abc", expect: []int{0, 6}},
		{s: "abab", p: "ab", expect: []int{0, 1, 2}},
	}

	for _, test := range testdata {
		get := findAnagrams01(test.s, test.p)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("s:%s, p:%v, expect:%v, get:%v", test.s, test.p, test.expect, get)
		}
	}

}
