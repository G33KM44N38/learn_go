package main

import (
	"fmt"
	// "log"
	"math"
	"math/cmplx"
	// "math/rand"
	// "example.com/greetings"
	"runtime"
	"time"
)

// ======================================================= module 1
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

// ======================================================= module 2
// return the square of a number
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// return a number elevate the his pow
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// exercerice loop functions
func Sqrt(x float64) float64 {
	z_r := float64(1)

	for i := 0; i < 10; i++ {
		z_r -= (z_r*z_r - x) / (2 * z_r)
		fmt.Println(z_r)
	}
	return z_r
}

func main() {
	// ======================================================= module 1
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	/* == comment module 1
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	// some of my test
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

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

	== end comment module 1*/

	// ======================================================= module 2

	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// // other version of the for loop withtout the post statement
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// For is Go's "while"
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// if statement
	fmt.Println(sqrt(4), sqrt(-4))

	// if with a short statement
	// the if statement can start with a short statement to execute before the condition
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// exercerice loop functions
	fmt.Println(Sqrt(2))

	// switch
	// in Go's switch the break is putted by default
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Godd evening")
	}

	// defer wait for the return before executing the next funcionts
	defer fmt.Println("world")
	fmt.Println("hello")

	// Stacking defers
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println((i))
	}

	fmt.Println("done")
}
