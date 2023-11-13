package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

// Methods and interfaces
func main() {
	// Methods are functions
	// + Pointer indirection
	v := Vertex3{3, 4}
	var pV = &v
	fmt.Println("v := Vertex3{3, 4}")
	fmt.Println("var pV = &v")
	// If a method is expecting a value as a receiver
	// you can pass a pointer to that value
	fmt.Println("v.Abs():", v.Abs())
	fmt.Println("pV.Abs():", pV.Abs())
	// If a method is expecting a value as argument
	// you cannot pass a pointer to that value
	fmt.Println("AbsFunc(v):", AbsFunc(v))
	fmt.Println("AbsFunc(pV): compile error")

	// Methods continued
	f := MyFloat(-math.Sqrt2)
	fmt.Println("f := MyFloat(-math.Sqrt2)")
	fmt.Println("f.Abs():", f.Abs())

	// Pointer receivers
	v.Scale(10)
	fmt.Println("v.Scale(10)")
	fmt.Println("v.Abs():", v.Abs())
	pV.Scale(10)
	fmt.Println("pV.Scale(10)")
	fmt.Println("v.Abs():", v.Abs())
	ScaleFunc(pV, 10) // == ScaleFunc(&v, 10)
	fmt.Println("ScaleFunc(pV, 10)")
	// ScaleFunc(v, 10)
	fmt.Println("ScaleFunc(v, 10) --> Not done, will compile error")
	fmt.Println("v.Abs():", v.Abs())

	// Interface
	//   A value of interface type can hold any value that implements
	//   those methods.
	var a Abser
	fmt.Println("var a Abser")
	f2 := MyFloat(-math.Sqrt2)
	v2 := Vertex3{3, 4}

	a = f2 // a MyFloat implements Abser
	fmt.Println("a = f (MyFloat)")
	a = &v2 // a *Vertex implements Abser
	fmt.Println("a = &v (*Vertex)")
	// a = v2
	fmt.Println("a = v (Vertex) --> not working")
	fmt.Println("a.Abs():", a.Abs())

	// Interface implemented implicitly, values
	var i I = T{"hello"}
	fmt.Println("var i I = T{\"hello\"}")
	// Calling a method on an interface value executes
	// the method of the same name on its underlying type.
	fmt.Println("i.M():")
	i.M()

	// Interface values with nil underlying values
	var i2 I
	// calling i2.M() here will result in a run-time error
	var t *T2
	i2 = t
	fmt.Println("var i2 I")
	fmt.Println("var t *T2")
	fmt.Println("i2 = t")
	fmt.Println("describe(i2):")
	describe(i2)
	fmt.Println("i2.M():")
	i2.M() // would be an error without method M() for type T2
	i2 = &T2{"hello you"}
	fmt.Println("i2 = &T{\"hello you\"}")
	fmt.Println("describe(i2):")
	describe(i2)
	fmt.Println("i2.M():")
	i2.M()

	// The empty interface
	var i3 interface{}
	fmt.Println("var i3 interface{}")
	fmt.Println("describeAny(i3):")
	describeAny(i3)
	i3 = 42
	fmt.Println("i3 = 42")
	fmt.Println("describeAny(i3):")
	describeAny(i3)
	i3 = "hello"
	fmt.Println("i3 = \"hello\"")
	fmt.Println("describeAny(i3):")
	describeAny(i3)

	// Type assertions
	var i4 interface{} = "hello"
	fmt.Println("var i4 interface{} = \"hello\"")
	s4 := i4.(string)
	fmt.Println("s4 := i4.(string)")
	fmt.Println("s4:", s4)
	s4, ok := i4.(string)
	fmt.Println("s, ok := i4.(string)")
	fmt.Println("s:", s4, "| ok:", ok)
	f4, ok := i4.(float64)
	fmt.Println("f4 ok := i4.(float64)")
	fmt.Println("f4:", f4, "| ok:", ok)
	fmt.Println("f4 = i.(float64) --> will trigger a panic")

	// Type switches (see method)
	do(21)
	do("hello")
	do(true)

	// Stringer (the interface for String() method)
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}
	r, e := Sqrt(2)
	fmt.Printf("Sqrt(2): %v | %v\n", r, e)
	r, e = Sqrt(-2)
	fmt.Printf("Sqrt(-2): %v | %v\n", r, e)

	// Readers
	newReader := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	displayReader(newReader, b)

	myReader := &MyReader{}
	b = make([]byte, 10)
	displayReader2(myReader, b)

	myOtherReader := &rot13Reader{r: strings.NewReader("Lbh penpxrq gur pbqr!")}
	b = make([]byte, 20)
	displayReader(io.Reader(myOtherReader), b)

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}
	io.Copy(os.Stdout, &r2)
	fmt.Println("")

	// Images
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	m2 := Image{Width: 100, Height: 100} // Set the width and height
	fmt.Println("m2 := Image{Width: 100, Height: 100}")
	fmt.Println("m2:", m2)
	fmt.Println("m2.Bounds()", m2.Bounds())
	fmt.Println("m2.ColorModel():", m2.ColorModel())
	fmt.Println("m2.At(10, 10):", m2.At(10, 10))
	pic.ShowImage(m2) /* from "golang.org/x/tour/pic" */
}

type Vertex3 struct {
	X, Y float64
}

/*
Go does not have classes. However, you can define methods
on types.
A method is a function with a special receiver argument.
*/
func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax
*T for some type T. (Also, T cannot itself be a pointer
such as *int.)
*/
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// These two method are equivalent
func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
An interface type is defined as a set of method signatures.
*/
type Abser interface {
	Abs() float64
}

/*
A type implements an interface by implementing its methods.
There is no explicit declaration of intent, no "implements"
keyword.
*/
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type T2 struct {
	S string
}

func (t *T2) M() {
	// cannot use == nil if was using 'T2' value
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
An empty interface may hold values of any type.
*/
func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
A type switch is a construct that permits several
type assertions in series.
*/
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type IPAddr [4]byte

// implementing the Stringer interface
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",
		ip[0], ip[1], ip[2], ip[3])
}

/*
The error type is a built-in interface similar to
fmt.Stringer:

	type error interface {
	    Error() string
	}
*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return 0, nil
	}
	return 0, &MyError{
		time.Now(),
		fmt.Sprintf("\ncannot Sqrt negative number: %v", x),
	}
}

func displayReader(reader io.Reader, b []byte) {
	for {
		n, err := reader.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (m MyReader) Read(b []byte) (i int, e error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func displayReader2(myReader *MyReader, b []byte) {
	n, err := myReader.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
}

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13Reader.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	}
	return b
}

/*
The image interface:

	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}
*/
type Image struct {
	Width, Height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}
