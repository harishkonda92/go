package interfacesI

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary () int
}

type Permenant struct{
	basicPay int
	pf int
	empid int
}

type Contract struct {
	basicPay int
	empid int
}
// calculates the salary of permanant emp
func (p Permenant) CalculateSalary() int  {
	return p.basicPay + p.pf
}
// calculates the salary of contract emp
func (c Contract)CalculateSalary()int  {
	return c.basicPay
}

func totalExpenses(s []SalaryCalculator) int  {
	expense := 0
	for _, sal := range s {
		expense = expense + sal.CalculateSalary()
	}
	return expense
}

func CalulateCompanyExpenses()  {
	// var salCal SalaryCalculator
	perm1 := Permenant{empid: 1, basicPay: 100, pf: 30,}
	perm2 := Permenant{empid: 2, basicPay: 110, pf: 35,}
	contr := Contract{empid: 2, basicPay: 110,}

	employess := []SalaryCalculator{perm1, perm2 , contr}
	totalExpnses := totalExpenses(employess)
	fmt.Println("total expenses : ", totalExpnses,"$")
}