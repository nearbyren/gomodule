package main

import (
	"fmt"
)

const name string = "小王"
const sex bool = false
const age int = 27
const price float64 = 12.5

/***
	break          default         func        interface          select
    case           defer           go          map                struct
    chan           else            goto        package            switch
    const          fallthough      if          range              type
    continue       for             import      return             var
*/
func main() {

	//删除scile中元素,删除下标为2的元素
	test := []int{10, 20, 30, 40, 50, 100}
	test = append(test[:2], test[5:]...)
	fmt.Println(test) //[10 20 40 50 100]

	//[10 20 10 20 30]

	fmt.Println("i = ", name, sex, age, price)
	for i := 0; i < 20; i++ {
		go test01(i)
	}
	var user User
	user.UserAdd = "s"
	user.UserAge = 12
	user.Username = "哈哈哈"

	user2 := User{
		Username: "结构体",
		UserAge:  28,
	}

	user3 := User{"jaja", 20, "海南"}

	fmt.Println(user, "\n", user2, "\n", user3, "\n")
	p := Person{
		"小米",
		1,
		user2,
	}
	fmt.Println( p)
}

type Person struct {
	string
	int
	User
}

type User struct {
	Username string
	UserAge  int
	UserAdd  string
}

func test01(i int) {
	var a = "hahah"
	fmt.Println("i = ", &a, i, name, sex, age, price)
}
