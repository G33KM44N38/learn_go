package module1

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// make a simple addition between 2 integer
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// naked retrurn
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// declare variable
var c, python, java bool = true, false, true

// Go's basic types are
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// const
const Pi = 3.14

// const number
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func Module1() {
	fmt.Println("my favorite number is ", rand.Intn(10))

	// all wor that begin with a capital letter is exported
	fmt.Println(math.Pi)

	// print the result of the function
	fmt.Println("addition between 2 num :", add(13, 32))

	// returning two string
	a, b := swap("hello", "world")
	fmt.Println("Test of returning 2 string", a, b)

	// naked returns
	fmt.Println(split(17))

	// use  declared function
	var i int = 12
	fmt.Println(i, c, python, java)

	// Short variable declarations
	var n, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(n, j, k, c, python, java)

	// basic type
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value.
	// The zero value is:
	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
	var u int
	var f float64
	var v bool
	var s string
	fmt.Printf("%v %v %v %q\n", u, f, v, s)

	// types conversion
	x := 2
	y := float64(x)
	z := uint64(y)

	fmt.Println(x, y, z)

	// type inference
	w := 28.8 // change me!
	fmt.Printf("v is of type %T\n", w)

	// constant
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules ?", Truth)

	// Numeric Constants
	fmt.Println("Small needInt", needInt(Small))
	fmt.Println("Small needFloat", needFloat(Small))
	fmt.Println("Big needFloat", needFloat(Big))
}
