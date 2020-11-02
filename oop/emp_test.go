package oop

import (
	"testing"
)

func TestEmp(t *testing.T) {
	e := newEmp()
	e.setName("suman")
	e.setSalary(1000)
	e.printData()
	e.raiseSalary(5000)
	e.printData()
}
func TestEmpNew(t *testing.T) {
	e := newEmp()
	e.setName("saurav")
	e.setSalary(3000)
	e.printData()
	e.raiseSalary(10000)
	e.printData()
}

func TestManager(t *testing.T) {
	e := newEmp()
	m := newManeger()
	m.department = "tech"
	m.Empl = e
	//fmt.Println(m.id)
	//e.setName("saurav")
	//e.setSalary(3000)
	//e.printData()
	//e.raiseSalary(10000)
	//e.printData()
}
