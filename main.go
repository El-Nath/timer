package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

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
		singleThread(arr, c)
	case "m":
		//multithread
		fmt.Print("Input number of concurrency (max = " + strconv.Itoa(len(arr)) + "): ")
		_, err := fmt.Scan(&cN)
		n, err := strconv.Atoi(cN)
		if err != nil || n > len(arr) {
			fmt.Println("--Not a valid option, terminating now--")
			break
		}
		multiThread(arr, n, c)
	default:
		//faultyinput
		fmt.Println("--Not a valid option, terminating now--")
		break
	}
}

func singleThread(arr []int, c chan os.Signal) {

	cT := time.Unix(time.Now().Unix(), 0)
	tT := 0

	for i := 0; i < len(arr); i++ {
		fmt.Print("Starting to sleep for " + strconv.Itoa(arr[i]) + " seconds at : " + time.Unix(time.Now().Unix(), 0).Format(time.UnixDate) + " -- ")
		tT += arr[i]
		time.Sleep(time.Duration(arr[i]) * time.Second)
		fmt.Println("Stopped sleeping at : " + time.Unix(time.Now().Unix(), 0).Format(time.UnixDate))

		go func(i int) {
			select {
			case sig := <-c:
				fmt.Printf("--Got %s signal. Aborting...\n--", sig)
				os.Exit(1)
			}
		}(i)
	}

	fmt.Println("")
	fmt.Println("Total time : " + strconv.Itoa(tT) + " seconds")
	fmt.Printf("Time elapsed : %.01f seconds", time.Since(cT).Seconds())
}

func multiThread(arr []int, n int, c chan os.Signal) {

	var wG sync.WaitGroup
	cT := time.Unix(time.Now().Unix(), 0)
	tT := 0

	wG.Add(n)

	for i := 0; i < len(arr); i++ {
		fmt.Println("Starting to sleep for " + strconv.Itoa(arr[i]) + " seconds at : " + time.Unix(time.Now().Unix(), 0).Format(time.UnixDate))
		tT += arr[i]
		go func(i int) {
			time.Sleep(time.Duration(arr[i]) * time.Second)
			fmt.Println("Stopped sleeping for " + strconv.Itoa(arr[i]) + " seconds at : " + time.Unix(time.Now().Unix(), 0).Format(time.UnixDate))
			defer wG.Done()
		}(i)
	}

	go func() {
		select {
		case sig := <-c:
			fmt.Printf("--Got %s signal. Aborting...\n--", sig)
			os.Exit(1)
		}
	}()

	wG.Wait()
	fmt.Println("")
	fmt.Println("Total time : " + strconv.Itoa(tT) + "seconds")
	fmt.Printf("Time elapsed : %.0f seconds", time.Since(cT).Seconds())

}
