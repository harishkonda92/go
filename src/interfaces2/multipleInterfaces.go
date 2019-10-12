package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalulator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstname   string
	lastname    string
	basicpay    int
	pf          int
	totalleaves int
	leavestaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary %d \n", e.firstname, e.lastname, (e.basicpay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalleaves - e.leavestaken
}

func main() {
	e := Employee{
		firstname:   "harish",
		lastname:    "konda",
		basicpay:    33000,
		pf:          100,
		totalleaves: 30,
		leavestaken: 12,
	}

	var sc SalaryCalculator = e

	sc.DisplaySalary()

	var leaveCal LeaveCalulator = e
	leavesleft := leaveCal.CalculateLeavesLeft()
	fmt.Printf("%s has %d leaves left \n", e.firstname, leavesleft)

}
