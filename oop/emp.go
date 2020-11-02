package oop

import "fmt"

type Emp struct {
	name    string
	sallary int64
	id      int
}

type Manager struct {
	Empl       Emp
	department string
}

func (e Emp) printData() {
	fmt.Println("data - ", e.id, e.name, e.sallary)
}

func newEmp() Emp {
	return Emp{}
}

func newManeger() Manager {
	return Manager{}
}

func (e *Emp) setName(name string) {
	e.name = name
}
func (e *Emp) setSalary(salary int64) {
	e.sallary = salary
}

func (e *Emp) raiseSalary(money int64) {
	e.sallary += money
}
