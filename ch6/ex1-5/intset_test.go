// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Example_three() {
	var x IntSet
	fmt.Println(x.Len())
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(1)

	fmt.Println(x.Len())

	// Output:
	// 0
	// 3
}

func Example_four() {
	var x IntSet
	x.Remove(30)
	fmt.Println(&x)

	x.Add(1)
	x.Add(144)
	x.Remove(5)

	fmt.Println(&x)

	x.Remove(144)
	fmt.Println(&x)

	// Output:
	// {}
	// {1 144}
	// {1}
}

func Example_five() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Clear()
	fmt.Println(&x)
	x.Add(35)
	fmt.Println(&x)

	// Output:
	// {}
	// {35}
}

func Example_six() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	y := x.Copy()
	x.Add(33)
	fmt.Println(&x)
	fmt.Println(y)
	// Output:
	// {1 33 144}
	// {1 144}
}

func Example_seven() {
	var x IntSet
	x.AddAll(1, 2, 3)
	fmt.Println(&x)
	// Output:
	// {1 2 3}
}

func Example_eight() {
	var x, y IntSet
	x.IntersectWith(&y)
	fmt.Println(&x)
	x.AddAll(1, 100, 300)
	y.AddAll(2, 300, 500)
	x.IntersectWith(&y)
	fmt.Println(&x)
	// Output:
	// {}
	// {300}
}

func Example_nine() {
	var x, y IntSet
	x.DifferenceWith(&y)
	fmt.Println(&x)
	x.AddAll(1, 100, 300)
	y.AddAll(2, 300, 500)
	x.DifferenceWith(&y)
	fmt.Println(&x)
	// Output:
	// {}
	// {1 100}
}

func Example_ten() {
	var x, y IntSet
	x.SymmetricDifference(&y)
	fmt.Println(&x)
	x.AddAll(1, 100, 300)
	y.AddAll(2, 300, 500)
	x.SymmetricDifference(&y)
	y.AddAll(501)
	fmt.Println(&x)
	// Output:
	// {}
	// {1 2 100 500}
}

func Example_eleven() {
	var x IntSet
	fmt.Println(x.Elems())
	x.AddAll(1, 65, 128)
	fmt.Println(x.Elems())
	// Output:
	// []
	// [1 65 128]
}
