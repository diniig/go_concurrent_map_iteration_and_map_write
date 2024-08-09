package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

type ResourceSnapshot struct {
}

func main() {
	println("Hello world!")
	var res map[string]string = map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for true {
			for k, v := range res {
				fmt.Printf("\nk: %v, v: %v", k, v)
				runtime.Gosched()
			}
		}
	}()

	go func() {
		i := 0
		for true {
			i = i + 1
			v := strconv.Itoa(i)
			setRes(res, v)
			runtime.Gosched()
		}
	}()

	wg.Wait()

	println("\nBye world!")
}

func setRes(res map[string]string, v string) {
	res["C"] = "c" + v
	res["B"] = "b" + v
	res["A"] = "a" + v
}
