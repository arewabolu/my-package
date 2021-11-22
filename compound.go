package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

/*var (
	capital    float64
	length     int
	percentage float64
)

func compoundCalc1(capital float64, percentage float64) {
	length = 1
	baseVal := 1 + (percentage / 100)
	fmt.Println(baseVal)
	factor := math.Pow(baseVal, float64(length))
	DayOne := capital * (factor - 1)
	fmt.Println(DayOne)
}
*/
func main() {
	//fmt.Scanln(&capital)
	//fmt.Scanln(&percentage)
	//compoundCalc1(float64(capital), float64(percentage))
	capital := os.Args[1]
	length := os.Args[2]
	percentage := os.Args[3]
	cap, err := strconv.ParseFloat(capital, 64)
	if err != nil {
		log.Fatalln(err)
	}
	duration, err := strconv.ParseFloat(length, 64)
	if err != nil {
		log.Fatalln(err)
	}
	interest, err := strconv.ParseFloat(percentage, 64)
	if err != nil {
		log.Fatalln(err)
	}
	baseVal := 1.0 + (interest / 100)
	fact := math.Pow(baseVal, float64(duration))
	DayOne := (cap * fact) - 1
	fmt.Print("The compound interest for the duration is:", DayOne)

}
