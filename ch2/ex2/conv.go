/* Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or
from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and
Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like. */
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64
type Celsius float64
type Fahrenheit float64

func main() {
	var input string

	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		_, err := fmt.Fscan(os.Stdin, &input)
		if err != nil {
			fmt.Printf("error! %v\n", err)
			os.Exit(1)
		}
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("not a number")
	}

	fmt.Printf("%v m = %v ft\n", Meter(num), Meter(num).toFeet())
	fmt.Printf("%v kg = %v lb\n", Kilogram(num), Kilogram(num).toPounds())
	fmt.Printf("%v °C = %v °F\n", Celsius(num), Celsius(num).toFahrenheit())
}

func (m Meter) toFeet() Feet {
	return Feet(m / 0.3048)
}

func (k Kilogram) toPounds() Pound {
	return Pound(k * 2.2046)
}

func (c Celsius) toFahrenheit() Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}
