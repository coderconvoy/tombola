package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var crapF map[string]int = map[string]int{
	"Labour": 2,
	"SNP":    3,
	"MNP":    4,
	"UKIP":   9,
	"Other":  10,
	"Cons":   14,
	"LibDem": 16,
	"Green":  17,
	"DUP":    18,
	"UUP":    19,
	"SF":     20,
}

type Constit struct {
	name  string
	votes map[string]int
}

func (c Constit) FptpWinner() (string, int) {
	winner := ""
	score := 0
	for k, v := range c.votes {
		if v > score {
			winner = k
			score = v
		}
	}
	return winner, score
}

func NewConstit(line string, cName int, plist map[string]int) (Constit, error) {
	ss := splitQoth(line)
	votes := make(map[string]int)
	if len(ss) < 20 {
		return Constit{name: "", votes: votes}, errors.New("No Data")
	}

	for k, v := range plist {
		n, err := strconv.Atoi(ss[v])
		if err != nil {
			continue
		}
		votes[k] = n
	}
	return Constit{name: ss[cName], votes: votes}, nil
}

func ReadConstits(r io.Reader, cName int, plist map[string]int) ([]Constit, error) {
	res := []Constit{}
	sc := bufio.NewScanner(r)
	sc.Scan() //Throw away csv titles
	line := -1
	for sc.Scan() {
		line++
		t := sc.Text()
		c, err := NewConstit(t, cName, plist)
		if err != nil {
			fmt.Printf("Error on line %d : %s \n\t%s", line, err, t)
			continue
		}
		res = append(res, c)
	}
	return res, nil
}

func ReadConstitsFile(fname string, cName int, plist map[string]int) ([]Constit, error) {
	f, err := os.Open(fname)
	if err != nil {
		return []Constit{}, err
	}
	defer f.Close()
	return ReadConstits(f, cName, plist)
}

//Comma separated with comas in data quoted
func splitQoth(t string) []string {
	res := []string{}
	curr := ""
	qoth := false
	for _, v := range t {
		if v == '"' {
			qoth = !qoth
			continue
		}
		if v == ',' && !qoth {
			res = append(res, curr)
			curr = ""
			continue
		}
		curr += string(v)
	}
	res = append(res, curr)
	return res
}
