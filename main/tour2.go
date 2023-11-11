package main

import (
	"fmt"
	"math"
)

func main() {
	// Go has pointers. A pointer holds the memory
	// address of a value.
	// The type *T is a pointer to a T value.
	// Its zero value is nil.
	pointers()

	// Struct fields are accessed using a dot
	vertex()
	vertex2()
	vertex3()

	// Arrays & Slices
	arrays()
	slices()
	slices2()
	slices3()
	printSlice([]int{2, 3, 5, 7, 11, 13})
	emptySlice()
	slicesRange()

	// Maps
	maps()
	mapLiterals()
	mapMutating()

	// Functions as values
	useFunctionAsValue()
	functionClosures()
	fibonacciClosure()
}

func pointers() {
	i, j := 42, 2701
	p := &i                // point to i
	fmt.Println("*p:", *p) // read i through the pointer
	*p = 21                // set i through the pointer
	fmt.Println("i:", i)   // see the new value of i
	p = &j                 // point to j
	*p = *p / 37           // divide j through the pointer
	fmt.Println("j:", j)   // see the new value of j
}

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func vertex() {
	fmt.Println("Vertex{1, 2}:", Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("Vertex x:", v.X)
}

func vertex2() {
	/*
		Struct fields can be accessed through a struct pointer.
		To access the field X of a struct when we have the
		struct pointer p we could write (*p).X.
		However, that notation is cumbersome, so the language
		permits us instead to write just p.X, without the explicit
		dereference.
	*/
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("Vertex{X, Y}:", v)
}

func vertex3() {
	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p := &Vertex{1, 2} // has type *Vertex
	fmt.Println("v1, p, v2, v3:", v1, p, v2, v3)
}

func arrays() {
	/*
		An array's length is part of its type, so arrays
		cannot be resized. This seems limiting, but don't worry;
		Go provides a convenient way of working with arrays.
	*/
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array
	var s []int = primes[1:4]            // slice
	fmt.Println("primes[1:4]:", s)
}

func slices2() {
	/*
		A slice does not store any data, it just describes a section
		of an underlying array.
		Changing the elements of a slice modifies the corresponding
		elements of its underlying array.
		Other slices that share the same underlying array will see
		those changes.
	*/
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names:", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("names[0:2]:", a, "names[1:3]:", b)

	b[0] = "XXX"
	fmt.Println("b[0] = \"XXX\"")
	fmt.Println("names[0:2]:", a, "names[1:3]:", b)
	fmt.Println("names:", names)
}

func slices3() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("[]int{2, 3, 5, 7, 11, 13}:", s)
	// The make function allocates a zeroed array
	// and returns a slice that refers to that array.
	// append is a built-in function to append elements
	// to a slice.
	s = append(make([]int, 0), 2, 3, 5, 7, 11, 13)
	fmt.Println("s", s)
	s = s[1:4]
	fmt.Println("s[1:4]:", s)
	s = []int{2, 3, 5, 7, 11, 13}
	s = s[:2]
	fmt.Println("s[2:]:", s)
	s = []int{2, 3, 5, 7, 11, 13}
	s = s[1:]
	fmt.Println("s[1:]:", s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func emptySlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func slicesRange() {
	/*
		The range form of the for loop iterates over a slice or map.
		When ranging over a slice, two values are returned for each
		iteration. The first is the index, and the second is a copy of
		the element at that index.
	*/
	fmt.Println("slicesRange():")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("index and value:")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println("value only:")
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

type Vertex2 struct {
	Lat, Long float64
}

func maps() {
	var m map[string]Vertex2
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	/*
		Map literals are like struct literals,
		but the keys are required.
	*/
	var m = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mapMutating() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func useFunctionAsValue() {
	fmt.Println("useFunctionAsValue():")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosures() {
	fmt.Println("functionClosures():")
	/*
		Go functions may be closures. A closure is a function
		value that references variables from outside its body.
		The function may access and assign to the referenced
		variables; in this sense the function is "bound" to
		the variables.
	*/
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	/*
		Implement a fibonacci function that returns a function
		(fibonacciClosure) that returns an int.
	*/
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacciClosure() {
	fmt.Println("fibonacciClosure():")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
