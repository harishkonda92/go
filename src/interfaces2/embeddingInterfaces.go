package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalulator interface {
	CalculateLeaves() int
}

type EmployeeOperatiosn interface {
	SalaryCalculator
	LeaveCalulator
}

type Employee struct {
	firstname, lastname                    string
	basicpay, pf, totalLeaves, leavestaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s salary is %d \n", e.firstname, (e.basicpay + e.pf))
}

func (e Employee) CalculateLeaves() int {
	return e.totalLeaves - e.leavestaken
}

type Dasher interface {
	Hello()
}

func main() {
	emp := Employee{
		firstname:   "Harish",
		lastname:    "konda",
		basicpay:    33000,
		pf:          100,
		leavestaken: 12,
		totalLeaves: 30,
	}

	var empOpe EmployeeOperatiosn = emp
	empOpe.DisplaySalary()
	fmt.Printf("%s has leaves %d leaves left \n", emp.firstname, empOpe.CalculateLeaves())

	var ds Dasher
	// ds.Hello()
	if ds == nil {
		fmt.Printf("ds is nil and has typr %T and value %v \n", ds, ds)
	}
}
