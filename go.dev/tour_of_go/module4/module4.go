// Package module4
package module4

import (
	"fmt"
	"math"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func module4_method() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// Remember: a method is just a function with a receiver argument
// Here's Abs written as a regular function with no change in functionality.
func module4_method_1() {
	v := Vertex{3, 4}
	fmt.Println((Abs(v)))
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare a method on non-struct types, too
func module4_method_continued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receiver
func module4_pointer_receiver() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// == we pass the strucutre itself and not the a copy of the value by passing the reference to it
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// pointer and function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func module4_point_n_function() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// method and pointer indirection
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func module4_method_pointer_indirection() {
	v := Vertex{3, 4}
	// here go automatically understand (&v).Scale() this is an indirection
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// method and pointer indirection 2
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func module4_method_pointer_indirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

// choose  a value or pointer receiver
func module4_pointer_or_receiver() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling : %+v, Abs: %v\n", v, v.Abs())
}

// interface
type Abser interface {
	Abs() float64
}

func module4_interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 5}

	a = f  // a MyFloat implement Abser
	a = &v // a *Vertex implements Abser

	// In the followin line, v is a Vertex (not *Vertex)
	// and does NOT implement abser

	a = v

	fmt.Println(a.Abs())
}

// == interface implement implicitly

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the infterface implement I,
// but we don't need to explicitly declare that it does so.

func (t T) M() {
	fmt.Print(t.S)
}

func module4_interface_implicitly() {
	var i I = T{"hello"}
	i.M()
}

// == interface value
type F float64

func (f F) M() {
	fmt.Println(f)
}

func module4_interface_value() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// == interface value with nil underlying values
type V interface {
	C()
}

func (t *T) C() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func module4_interface_value_nil() {
	var i V

	var t *T
	i = t
	describe2(i)
	i.C()

	i = &T{"hello"}
	describe2(i)
	i.C()
}

func describe2(i V) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ==nil interface value
func module4_interface_value_nil_call() {
	// this should not work
	// cause M method isn't initialized
	type I interface {
		M()
	}

	var i I
	describe(i)
	i.M()
}

// == empty interface
// used when we don't now the incoming type value
func module4_empty_interface() {
	var i interface{}

	describe3(i)

	i = 42
	describe3(i)

	i = "hello"
	describe3(i)
}

// describe function
func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// type assertions
func module4_type_assertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// will trigger a panic cause the underlying value does not exit, and we didn't check for it
	// like above
	f = i.(float64) // panic~error

	fmt.Println(f)
}

// type switch
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)

	}
}

func module4_type_switch() {
	do(21)
	do("Hello")
	do(true)
}

// stringers
// A stringer is a type that can descrbe itself as a string
func module4_stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// stringers exercices
func module4_stringer_exercice() {
	hosts := map[string]IPAddr{
		"loopbakc":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// error
func module4_error() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// error exercices
func Sqrt(x float64) (float64, error) {
	if (x < 0)
		
	return 0, nil
}

func module4_exercie_error() {
	fmt.Println(Sqrt((2)))
	fmt.Println(Sqrt((-2)))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() {
	fmt.Printf("cannot Sqrt negative number: %v", e)
}

func Module4() {
	// module4_method()
	// module4_method_1()
	// module4_pointer_receiver()
	// module4_point_n_function()
	// module4_method_pointer_indirection()
	// module4_method_pointer_indirection2()
	// module4_pointer_or_receiver()
	// module4_interface()
	// module4_interface_implicitly()
	// module4_interface_value()
	// module4_interface_value_nil()
	// module4_interface_value_nil_call()
	// module4_empty_interface()
	// module4_type_assertions()
	// module4_type_switch()
	// module4_stringer()
	// module4_stringer_exercice()
	// module4_error()
	module4_exercie_error()
}
