package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	constits, err := ReadConstitsFile("test_data/EconomistUK2017.csv", 5, crapF)
	if err != nil {
		log.Fatal(err)
	}

	seats := make(map[string]int)

	for k, v := range constits {
		w, sc := v.FptpWinner()
		if sc == 0 {
			fmt.Println("Error on line ", k, ":", v.name)
		}
		seats[w] += 1
	}

	stots := SumTotals(constits)

	pr := DivideAmong(650, stots)
	PrintResult("Total Sum", pr)

	PrintResult("FPTP", seats)
	fmt.Println("SCORE: ", Compare(pr, seats))

	tomb := Tombola(constits)
	bscore := Compare(pr, tomb)
	for i := 0; i < 1000; i++ {
		t2 := Tombola(constits)
		sc2 := Compare(pr, t2)
		if sc2 < bscore { //&& OrderPreserve(stots, t2) {
			fmt.Printf(".")
			tomb = t2
			bscore = sc2
		}
	}
	PrintResult("Tombola", tomb)
	fmt.Println("SCORE: ", bscore)
}

func PrintResult(title string, seats map[string]int) {
	sum := 0
	fmt.Println(title)

	sorter := []ptmod{}
	for k, v := range seats {
		sorter = append(sorter, ptmod{m: v, p: k})
		sum += v
	}
	sort.Sort(modc(sorter))
	for _, v := range sorter {
		fmt.Println("\t", v.p, ":", v.m)
	}
	fmt.Println("Total:", sum)
}

func SumTotals(cc []Constit) map[string]int {
	res := make(map[string]int)
	for _, v := range cc {
		for kk, vv := range v.votes {
			res[kk] += vv
		}
	}
	return res
}

func Tombola(cc []Constit) map[string]int {
	res := make(map[string]int)
	for _, v := range cc {
		balls := v.GetBalls(0, 100)
		winner := SelectBall(balls)
		res[winner]++
	}
	return res
}
