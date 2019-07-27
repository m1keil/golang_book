
package main

import (
	"golang/chapter7/ex7-6/tempconf"
	"flag"
	"fmt"
)

var temp = tempconf.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
