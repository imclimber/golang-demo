package main

import "testing"

// 两个引用类型相等的前提是：两者的hash值相等，两者的字面量相等。
func TestT(t *testing.T) {
	type S struct {
		ID int
	}
	s1 := S{ID: 1}
	s2 := S{ID: 1}

	var h = map[S]int{}
	h[s1] = 1
	t.Log(h[s1])
	t.Log(h[s2])
	t.Log(s1 == s2)

	var hh = map[*S]int{}
	hh[&s1] = 1
	t.Log(hh[&s1])
	t.Log(hh[&s2])
	t.Log(s1 == s2)
}
