package main

import (
	"fmt"
	"log"
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

	PrintResult("FPTP", seats)

	pr := DivideAmong(650, SumTotals(constits))
	PrintResult("Total Sum", pr)
}

func PrintResult(title string, seats map[string]int) {
	sum := 0
	fmt.Println(title)
	for k, v := range seats {
		sum += v
		fmt.Println("\t", k, ":", v)
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
