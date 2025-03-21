package main

import "fmt"

////////////////////////////////////////////////////////////
// Method
// 1. Method를 사용해서 데이터와 기능을 묶어 응집도를 높인다
// 2. 코드 재사용성을 높여 모듈화로 코드 관리성과 가독성이 향상
////////////////////////////////////////////////////////////

type Account struct {
	balance int
}

// Nomal function : func MethodName(Receiver, Parameter) { ... }
func Func(a *Account, amount int) {
	a.balance -= amount
}

// Method function : func (Receiver) MethodName(Parameter) { ... }
func (a *Account) Method(amount int) {
	a.balance -= amount
}

func main() {
	a := &Account{100} // Account pointer variable create
	Func(a, 30)
	a.Method(30)

	fmt.Printf("%d\n", a.balance)
}
