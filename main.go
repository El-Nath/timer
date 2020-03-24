package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	type Inputs []string
	var input, cN string

	fmt.Print("Sleep single threaded or multi threaded (s/m): ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err.Error())
	}
}
