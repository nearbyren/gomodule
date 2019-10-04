package main


import (
	"fmt"
	"github.com/nearbyren/gomodule/router"
)


func main() {
	fmt.Println("This works.")
	r := router.InitRouter()
	r.Run()
}

