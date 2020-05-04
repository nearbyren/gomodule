package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var a float64
	fv :=reflect.ValueOf(a)
	var c = &fv
	c.SetFloat(3.0)
	fmt.Printf("%v1\n",c)


	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}

		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
