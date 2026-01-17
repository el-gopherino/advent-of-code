package main

import (
	"fmt"
)

type celsius float64

type temperature struct {
	celsius
}

func (c celsius) String() string {
	return fmt.Sprintf("%.2f ºC", c)
}

func main() {
	c := celsius(10.0)
	fmt.Println(c)

	t := temperature{c}
	fmt.Println(t)

}
