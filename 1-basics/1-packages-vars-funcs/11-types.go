package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Pi = 3.14

	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it raight again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	// var ii = 42
	// var ff = float64(ii)
	// var z uint = uint(ff)
	// simple version
	ii := 42
	ff := float64(ii)
	u := uint(ff)
	fmt.Println(ii, ff, u)

	// more conversion
	var xx, yy int = 3, 4
	var fff float64 = math.Sqrt(float64(xx*xx + yy*yy))
	var z uint = uint(fff)
	fmt.Println(xx, yy, fff, z)

	// constants
	const World = ""
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
