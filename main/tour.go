package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
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

	// Only looping construct, the for loop (see git push -u -f origin master)
	fmt.Println("sum10Times(12):", sum10Times(12))

	//  In the loop init and post statements are optional.
	fmt.Println("sumWithLoop(400):", sumWithLoop(400))

	// Using if (see method)
	fmt.Println("sqrt(2):", sqrt(2), "sqrt(-4):", sqrt(-4))

	// Using if with condition (see method)
	// Using else (see method)
	fmt.Println("pow(3, 2, 10):", pow(3, 2, 10),
		"pow(3, 3, 20):", pow(3, 3, 20))

	fmt.Println("sqrtFinder(81):", sqrtFinder(81))

	// Switch case (see method)
	fmt.Println("getOS():", getOS())

	// Switch case with evaluate case
	fmt.Println("When's Saturday?", FindSaturday())
	fmt.Println(greetings())

	// A defer statement defers the execution of a function
	// until the surrounding function returns.
	defer fmt.Println("LAST MESSAGE")

	countingUsingDefer()
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

func sum10Times(x int) (sum int) {
	for i := 0; i < 10; i++ {
		sum += x
	}
	return sum
}

func sumWithLoop(x int) (sum int) {
	sum = x
	// In Go while = for with init and post statements
	for sum < 1000 {
		sum += sum
	}
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//  Like for, the if statement can start with a short
	// statement to execute before the condition.
	// ---
	// Variables declared by the statement are only in scope
	// until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// Variables declared inside an if short statement
		// are also available inside any of the else blocks.
		fmt.Println("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Loop and if
func sqrtFinder(x float64) float64 {
	z := float64(1)
	previous := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("i=", i, "|--> z=", z)
		if math.Abs(z-previous) < 0.001 {
			return z
		}
		previous = z
	}
	return z
}

// Switch case
func getOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X."
	case "linux":
		return "Linux."
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return os
	}
}

func FindSaturday() string {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		return "Today."
	case today + 1:
		return "Tomorrow."
	case today + 2:
		return "In two days."
	default:
		return "Too far away."
	}
}

func greetings() string {
	t := time.Now()
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long
	// if-then-else chains.
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon."
	default:
		return "Good evening."
	}
}

func countingUsingDefer() {
	fmt.Println("countingUsingDefer start")
	for i := 0; i < 10; i++ {
		// bad practice, can cause issues
		defer fmt.Println("countingUsingDefer:", i)
	}
	fmt.Println("countingUsingDefer done")
}
