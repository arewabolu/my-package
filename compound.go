package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const rebase int = 3 // rebase every 8 hours i.e 3rebases per day

type initVal struct {
	Principal float64 //principal
	Time      int     //duration in days
	Interest  float64 //percentage of interest
}

func (values *initVal) compoundCalc() float64 {

	k := float64(rebase) * float64(values.Time)
	//
	baseVal := 1 + (values.Interest / 100)
	//totalValue := math.Pow(compVal, k)
	return values.Principal * math.Pow(baseVal, k)
}

func main() {
	P := os.Args[1]

	P2, err := strconv.ParseFloat(P, 64)
	if err != nil {
		log.Fatalln(err)
	}

	T := os.Args[2]

	T2, err := strconv.ParseFloat(T, 64)
	if err != nil {
		log.Fatalln(err)
	}

	n := os.Args[3]
	n2, err := strconv.ParseFloat(n, 64)
	if err != nil {
		log.Fatalln(err)
	}

	compVal := initVal{P2, int(T2), n2}

	fmt.Sprintf("The compound value is:%.2f", compVal.compoundCalc())

}
