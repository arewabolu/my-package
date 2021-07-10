package StrtoInt

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var Usinput string
	fmt.Scanln(&Usinput)
	Endres := Receiver(Usinput)
	fmt.Println(Endres)

}

func Receiver(inp string) []int {
	// Code here or
	SplitString := strings.Split(inp, "")
	ints := make([]int, len(SplitString))
	for i, k := range SplitString {
		ints[i], _ = strconv.Atoi(k)
	}
	return ints
}
