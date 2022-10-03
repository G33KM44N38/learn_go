package module3

import (
	"fmt"
	"math"
	"strings"
	// "golang.org/x/tour/wc"
)

func module3_1() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func module3_struct_1() {
	fmt.Println(Vertex{1, 3})
}

func module3_struct_2() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func module3_struct_3() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func module3_struct_4() {
	fmt.Println(v1, p, v2, v3)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func module3_array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 6, 11, 13}
	fmt.Println(primes)
}

// A slice does not store any data, it just describes a section of an underlying module3_array
// Changing the elements of a slice modifies the corresponding elements of its underlying array
func module3_slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:5]
	fmt.Println(s)
}

func module3_slice_1() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// slice literal is like an array literal without the length
func module3_slice_literals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

func module3_slice_default() {
	s := []int{2, 3, 4, 5, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func module3_slice_length_and_capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Use one or the other part

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap =%d %v\n", len(s), cap(s), s)
}

func module3_nil_slices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil !")
	}
}

// Slices can be created with the built-in make function;
// this is how you create dynamically-sized arrays
// The make function allocates a zeroed array and returns a slice that refers to that array
func module3_slice_make() {
	a := make([]int, 5)
	printSlice_make("a", a)

	b := make([]int, 0, 5)
	printSlice_make("b", b)

	c := b[:2]
	printSlice_make("c", c)

	d := c[2:5]
	printSlice_make("d", d)
}

// print a slice and his capacity
func printSlice_make(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Slices can contain any type, including other slices
func module3_slice_of_slice() {
	// Create a tit-tac-toe board

	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	// The Players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// appending to a slice
// don't forget that it refer to an array
// if there is not enough space append func will allocate more space and pointed to a new array
func module3_appending_to_slice() {
	var s []int
	printSlice(s)

	// apend works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as nedded.
	s = append(s, 1)
	printSlice(s)

	// We can ad dmore than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// The range form of the for loop iterates over a slice or map
// When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.
func module3_slice_range() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// like range above
// You can skip the index or value by assigning to _
func module3_slice_range_continued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// Implement Pic. It should return a slice of length
// dy, each element of which is a slice of dx 8-bit
// unsigned integers. When you run the program, it
// will display your picture, interpreting the integers as
// grayscale (well, bluescale) values.
// The choice of image is up to you. Interesting
// functions include (x+y)/2, x*y, and x^y.
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// (Use uint8(intValue) to convert between types.)

func Pic(dx, dy int) [][]uint8 {
	tab := [][]uint8{}
	return tab
}

func module3_exercice_slices() {
}

// a map maps keys to values
type Vertez struct {
	Lat, Long float64
}

var m map[string]Vertez

// A map maps keys to values
// The zero value of a map is nil.
// A nil map has no keys, nor can keys be added.
// The make functions returns a map of the given type, initialized and ready for use.
func module3_map() {
	m = make(map[string]Vertez)
	m["Bell Labs"] = Vertez{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// Map literals are like struct literals, but the keys are required
func module3_map_literal() {
	m := map[string]Vertez{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
		"kylian": {
			484.8949, 8943.448,
		},
	}
	fmt.Println(m)
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
func module3_map_literal_continued() {
	m := map[string]Vertez{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func module3_mutating_map() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present", ok)
}

// Functions are values too. They can be passed around just like other values.
func module3_function_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(math.Hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Go functions may be closures.
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables
func module3_function_closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// exerccie on fibonacci
func module3_fibonacci_closure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func Module3() {
	// module3_1()
	// module3_struct_1()
	// module3_struct_2()
	// module3_struct_3()
	// module3_struct_4()
	// module3_array()
	// module3_slice()
	// module3_slice_1()
	// module3_slice_literals()
	// module3_slice_default()
	// module3_slice_length_and_capacity()
	// module3_nil_slices()
	// module3_slice_make()
	// module3_slice_of_slice()
	// module3_appending_to_slice()
	// module3_slice_range()
	// module3_slice_range_continued()
	// module3_map()
	// module3_map_literal()
	// module3_map_literal_continued()
	// module3_mutating_map()
	// module3_function_value()
	// module3_function_closure()
	// module3_fibonacci_closure()
}
