package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/igorverse/a-linguagem-de-programacao-go/estrutura-dos-programas/measureconv"
	"github.com/igorverse/a-linguagem-de-programacao-go/estrutura-dos-programas/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		cm := measureconv.Centimeter(v)
		p := measureconv.Pound(v)

		fmt.Println("Temperatures")
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Println("Measures")
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", cm, measureconv.CMToFT(cm), cm, measureconv.CMToM(cm), p, measureconv.PToKG(p))
	}

}
