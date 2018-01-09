package main

import (
	"flag"
	"fmt"
	"time"
)

var x *int = flag.Int("x", 12, "tak(x,y,z)")
var y *int = flag.Int("y", 6, "tak(x,y,z)")
var z *int = flag.Int("z", 0, "tak(x,y,z)")
var counter int

func tak(x, y, z int) int {
	counter++
	if x <= y {
		return y
	} else {
		return tak(tak(x-1, y, z), tak(y-1, z, x), tak(z-1, x, y))
	}
}

func main() {
	flag.Parse()
	s := time.Now()

	tak(*x, *y, *z)

	e := time.Since(s)
	fmt.Printf("%d times call\n", counter)
	fmt.Printf("%d nsec\n", e.Nanoseconds())
	fmt.Printf("%f sec\n", e.Seconds())
}
