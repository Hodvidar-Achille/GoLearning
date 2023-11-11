package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// variable declarations may be "factored" into blocks,
// as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", rand.Intn(10))

	// Using format
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// In Go, a name is exported if it begins with a capital letter
	fmt.Println("math.Pi =", math.Pi)

	// A function can take zero or more arguments.
	fmt.Println("add method:", add(42, 13))

	// When two or more consecutive named function parameters
	// share a type, you can omit the type from all but the last.
	fmt.Println("multiply method:", multiply(300, 2))

	// A function can return any number of results.
	a, b := swap("hello", "world")
	fmt.Println("Swap method returns:", a, b)

	// Go's return values may be named. If so, they are treated
	// as variables defined at the top of the function.
	fmt.Print("Split method for 34: ")
	fmt.Println(split(34))

	// The var statement declares a list of variables;
	// as in function argument lists, the type is last
	var c, python, java bool
	var i int
	fmt.Println("Some variables:", i, c, python, java)

	// A var declaration can include initializers, one per variable.
	// If an initializer is present, the type can be omitted
	var c2, python2, java2 = true, false, "no!"
	var i2, j2 int = 1, 2
	fmt.Println("Some variables (2):", i2, j2, c2, python2, java2)

	// Inside a function, the := short assignment statement can be
	// used in place of a var declaration with implicit type.
	k := 3
	fmt.Println("Some variables (3):", k)

	// Variables can be declared in a block (see var block)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value
	var i3 int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("Default variables type value %v %v %v %q\n", i3, f, b2, s)

	// T(v) convert the value v to the T type
	i4 := 42
	f2 := float64(i)
	u2 := uint(f)
	fmt.Println("Some variables (4):", i4, f2, u2)

	// the variable's type is inferred from the value on the right hand side.
	i5 := 42          // int
	f3 := 3.142       // float64
	g := 0.867 + 0.5i // complex128
	fmt.Println("Some variables (5):", i5, f3, g)

	// Constants are declared like variables, but with the const keyword.
	const World = "世界"
	fmt.Println("Hello", World)

	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	fmt.Println("needInt(Small):", needInt(Small))
	fmt.Println("needFloat(Small):", needFloat(Small))
	fmt.Println("needFloat(Big):", needFloat(Big))
}

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return /* x, y */
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
