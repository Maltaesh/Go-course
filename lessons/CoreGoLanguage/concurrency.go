package main

import (
	utils "example/project/mypackage"
	"sync"
	"time"
)

type Acconut struct {
	balance int
	lock    sync.Mutex
}

func (a *Acconut) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Acconut) Withdraw(w int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if w > a.balance {
		utils.Pl("Not enough money in account")
	} else {
		utils.F("%d Withdrawn: Balance: %d\n", w, a.balance)
		a.balance -= w
	}
}

func main() {
	var acct Acconut
	acct.balance = 100
	utils.Pl("Balance:", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}

	time.Sleep(2 * time.Second)
}
