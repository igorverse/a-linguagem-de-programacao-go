package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%d microseconds elapsed\n", time.Since(start).Microseconds())

	start = time.Now()
	for i, arg := range os.Args {
		fmt.Printf("%v %v\n", i, arg)
	}
	fmt.Printf("%d microseconds elapsed\n", time.Since(start).Microseconds())
}
