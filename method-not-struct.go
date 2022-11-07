package main

import (
	"fmt"
	"math"
)

type convertRound float64

func (m convertRound) ceil() float64 {
	return math.Ceil(float64(m))
}

func (m convertRound) round() float64 {
	return math.Round(float64(m))
}

func main() {
	around := convertRound(45.47)
	fmt.Printf("number around %f\n", around.ceil())
	fmt.Printf("number around %f\n", around.round())
}
