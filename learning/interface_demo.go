package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}
type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

/********************************************/
//工资计算器
type SalaryCalculator interface {
	CalculateSalary() int
}

//长期员工
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

//合同员工
type Contract struct {
	empId    int
	basicpay int
}

//长期员工
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//合同员工
func (c Contract) CalculateSalary() int {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

/********************************************/

type Test interface {
	Tester()
}
type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}
func describe(t Test) {
	fmt.Printf("interface type %T value %v1\n", t, t)
}
func main() {
	/*name := MyString("Sam Anderson")
	var v1 VowelsFinder
	v1 = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v1.FindVowels())*/

	/*pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)*/

	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
