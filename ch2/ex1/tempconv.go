// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin
// is −273.15°C and a difference of 1K has the same magnitude as 1°C.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Celsius = -273.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsius     { return Celsius(k) + ZeroK }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
