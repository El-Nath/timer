package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	type Inputs []string
	var input, cN string

	fmt.Print("Sleep single threaded or multi threaded (s/m): ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "s":
		//singlethread
	case "m":
		//multithread
		fmt.Print("Input number of concurrency: ")
		_, err := fmt.Scan(&cN)
		n, err := strconv.Atoi(cN)
		if err != nil {
			fmt.Println("--Not a valid option, terminating now--")
			break
		}

	default:
		//faulty input
		fmt.Println("--Not a valid option, terminating now--")
		break
	}
}
