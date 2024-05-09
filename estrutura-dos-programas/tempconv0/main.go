package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoillingC     Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g\n", BoillingC-FreezingC)
	fmt.Printf("%v -> type: %T\n", CToF(100), CToF(100))
	fmt.Printf("%v -> type: %T\n", FToC(100), FToC(100))
}
