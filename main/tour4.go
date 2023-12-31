package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"strconv"
	"strings"
	"time"
)

// Concurrency:
//    Goroutines
//    Channels
//    Buffered Channels
//    Range and Close
//    Select

func main() {
	// A Goroutines is a lightweight thread managed
	// by the Go runtime.
	// starts a new goroutine running
	go say("world")
	say("hello")

	// Channels are a typed conduit through which you
	// can send and receive values with the channel operator, <-.
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println("s := []int{7, 2, 8, -9, 4, 0}")
	fmt.Println("c := make(chan int)")
	fmt.Println("go sum(s[:len(s)/2], c)")
	fmt.Println("go sum(s[len(s)/2:], c)")
	fmt.Println("x, y := <-c, <-c ")
	fmt.Println("x, y, x+y:", x, y, x+y)

	// Buffered Channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3  // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Range and Close
	/*
		A sender can close a channel to indicate that no more values will be sent.
		Receivers can test whether a channel has been closed by assigning a second
		parameter to the receive expression
	*/
	cha := make(chan int, 10)
	go fibonacci4(cap(cha), cha)
	fmt.Println("go fibonacci4(cap(cha), cha)")
	// The loop for i := range c receives values from the channel repeatedly until
	// it is closed.
	for i := range cha {
		fmt.Println(i)
	}

	// Select
	/*
		The select statement lets a goroutine wait on multiple communication operations.
		A select blocks until one of its cases can run, then it executes that case.
		It chooses one at random if multiple are ready.
	*/
	c5 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("<-c5:", <-c5)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("quit <- 0")
		quit <- 0
	}()
	fmt.Println("go fibonacci5(c5, quit)")
	fibonacci5(c5, quit)

	// Default Selection (see fibonacci5)

	// Exercise: Equivalent Binary Trees
	// using "golang.org/x/tour/tree"

	// test
	root := &tree.Tree{
		Value: 1,
		Left: &tree.Tree{
			Value: 0,
		},
		Right: &tree.Tree{
			Value: 2,
			Right: &tree.Tree{
				Value: 3,
			},
		},
	}
	printTreeComplexImproved(root)

	tree1 := tree.New(5)
	tree2 := tree.New(5)
	fmt.Println("printTreeComplexImproved(tree1)")
	printTreeComplexImproved(tree1)
	fmt.Println("printTreeComplexImproved(tree2)")
	printTreeComplexImproved(tree2)
	fmt.Println("Same(tree1, tree2):", Same(tree1, tree2))
	defer testLast()
}

func testLast() {
	tree1 := tree.New(5)
	tree2 := tree.New(10)
	fmt.Println("printTreeComplexImproved(tree1)")
	printTreeComplexImproved(tree1)
	fmt.Println("printTreeComplexImproved(tree2)")
	printTreeComplexImproved(tree2)
	fmt.Println("Same(tree1, tree2):", Same(tree1, tree2))
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
Note: Only the sender should close a channel, never the receiver.
Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming,
such as to terminate a range loop.
*/
func fibonacci4(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci5(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("case c <- x: x, y = y, x+y")
			x, y = y, x+y
		case <-quit:
			fmt.Println("<-quit: quitting method fibonacci5(c, quit)")
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

/*
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println("Walk in a tree")
	if t.Left != nil {
		fmt.Println("going left")
		Walk(t.Left, ch)
	} else {
		fmt.Println("take value")
		ch <- t.Value
	}
	if t.Right != nil {
		fmt.Println("going right")
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	fmt.Println("Same() method start")
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		select {
		case x := <-c1:
			y := <-c2
			fmt.Printf("x := <-c1 (x=%v) y := <-c2 (y=%v)\n", x, y)
			if x != y {
				return false
			}
		case <-c1:
			fmt.Println("<-c1 Same() method end")
			return true
		}
	}
}

func printTree(t *tree.Tree) {
	if t == nil {
		return
	}

	printTree(t.Left)    // Traverse left subtree
	fmt.Println(t.Value) // Print the value
	printTree(t.Right)   // Traverse right subtree
}

type NodeInfo struct {
	Node   *tree.Tree
	Parent int
}

func printTreeComplex(root *tree.Tree) {
	if root == nil {
		return
	}

	var nodes []*tree.Tree
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		size := len(nodes)
		currentLevel := []string{}

		for i := 0; i < size; i++ {
			node := nodes[0]
			nodes = nodes[1:]

			if node != nil {
				currentLevel = append(currentLevel, strconv.Itoa(node.Value))
				nodes = append(nodes, node.Left)
				nodes = append(nodes, node.Right)
			} else {
				currentLevel = append(currentLevel, "nil")
			}
		}

		fmt.Println(strings.Join(currentLevel, " - "))
	}
}

func printTreeComplexImproved(root *tree.Tree) {
	if root == nil {
		return
	}

	nodes := []NodeInfo{{Node: root, Parent: -1}} // -1 indicates no parent

	for len(nodes) > 0 {
		size := len(nodes)
		currentLevel := []string{}

		for i := 0; i < size; i++ {
			nodeInfo := nodes[0]
			nodes = nodes[1:]

			var nodeStr string
			if nodeInfo.Node != nil {
				nodeStr = strconv.Itoa(nodeInfo.Node.Value)
				if nodeInfo.Parent != -1 {
					nodeStr = fmt.Sprintf("(%d) %s", nodeInfo.Parent, nodeStr)
				}

				nodes = append(nodes, NodeInfo{Node: nodeInfo.Node.Left, Parent: nodeInfo.Node.Value})
				nodes = append(nodes, NodeInfo{Node: nodeInfo.Node.Right, Parent: nodeInfo.Node.Value})
			} else {
				if nodeInfo.Parent != -1 {
					nodeStr = fmt.Sprintf("(%d) nil", nodeInfo.Parent)
				} else {
					nodeStr = "nil"
				}
			}
			currentLevel = append(currentLevel, nodeStr)
		}

		fmt.Println(strings.Join(currentLevel, " - "))
	}
}
