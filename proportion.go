package main

import "sort"

type ptmod struct {
	p string
	m int
}

type modc []ptmod

func (m modc) Len() int           { return len(m) }
func (m modc) Less(a, b int) bool { return m[a].m > m[b].m } //Sorts descending
func (m modc) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }

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
