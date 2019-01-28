
package main

import (
	"./tempconf"
	"flag"
	"fmt"
)

var temp = tempconf.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
