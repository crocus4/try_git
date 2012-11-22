package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tip(19.98))	// 3.02
	fmt.Println(tip(29.23))	// 4.77
	fmt.Println(tip(7.54))	// 1.46
}

func tip(bill float64) (tip, total float64) {
	total = math.Ceil(1.15 * bill)
	tip = total - bill
	return
}
