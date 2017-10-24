package main

import (
	"fmt"
	"sort"
)

type ptmod struct {
	p string
	m int
}

type modc []ptmod

func (m modc) Len() int           { return len(m) }
func (m modc) Less(a, b int) bool { return m[a].m > m[b].m } //Sorts descending
func (m modc) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }

type moda []ptmod

func (m moda) Len() int           { return len(m) }
func (m moda) Less(a, b int) bool { return m[a].m < m[b].m } //Sorts ascending
func (m moda) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }

func SortMap(a map[string]int, asc bool) []ptmod {
	res := []ptmod{}
	for k, v := range a {
		res = append(res, ptmod{k, v})
	}
	if asc {
		sort.Sort(moda(res))
		return res
	}
	sort.Sort(modc(res))
	return res
}

func DivideAmong(n int, mp map[string]int) map[string]int {
	res := make(map[string]int)
	mods := []ptmod{}
	sum := 0
	tot := 0
	//Simple Average divide
	for _, v := range mp {
		sum += v
	}
	for k, v := range mp {
		seats := (v * n) / sum
		res[k] = seats
		tot += seats
		mods = append(mods, ptmod{m: (v * n) % sum, p: k})
	}
	//Leftovers

	sort.Sort(modc(mods))

	for k, v := range mods {
		if tot+k >= n {
			return res
		}
		res[v.p] += 1
	}
	return res
}

//Compare returns an int comparing differences in vote representation,
//A low score is close votes.
func Compare(a, b map[string]int) int {
	score := 0
	for k, av := range a {
		bv, _ := b[k]
		if av-bv > 0 {
			score += av - bv
		} else {
			score += bv - av
		}
	}

	for k, bv := range b {
		if _, ok := a[k]; !ok {
			score += bv
		}
	}

	return score
}

func OrderPreserve(a, b map[string]int) bool {
	as := SortMap(a, true)
	bs := SortMap(b, true)

	if len(a) < len(b) {
		fmt.Println(a)

		return false
	}
	for k, v := range bs {
		if as[k].p != v.p {
			return false
		}
	}
	return true

}
