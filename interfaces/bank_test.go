package interfaces

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	v, _ := NewAccount(12345)
	fmt.Println(v)
	v.Deposit(100)
	fmt.Println(v)

	v1, _ := NewAccount(23456)
	fmt.Println(v1)
	v1.Deposit(100)
	fmt.Println(v1)

}
