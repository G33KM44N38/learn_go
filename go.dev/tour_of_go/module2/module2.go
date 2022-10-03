package module2

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

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

func Module2() {
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
}
