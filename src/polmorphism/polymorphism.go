package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func calculateNetIncome(ic []Income) {
	netincome := 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "project 2", biddedAmount: 3000}
	project3 := TimeAndMaterial{projectName: "Time and Material", noOfHours: 160, hourlyRate: 25}
	project4 := Advertisement{adName: "Bluechew", CPC: 5, noOfClicks: 300}
	incomeStreams := []Income{project1, project2, project3, project4}
	calculateNetIncome(incomeStreams)
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}
