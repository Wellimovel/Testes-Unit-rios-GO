package main

import "fmt"

func main() {
	x := soma(1, 5, 3)
	fmt.Println(x)

}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	fmt.Println(total)
	return total
}
