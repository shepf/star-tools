package main

import (
	"fmt"
	"runtime"
)

func foo() {
	for i := 0; i < 50; i++ {
		go func() {
			fmt.Print(i, " ")

		}()
	}
	runtime.Gosched()

}

func main() {
	foo()
}
