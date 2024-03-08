package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	defaultBalance int = 100
)

type bankAccount struct {
	id      string
	balance int
	mtx     sync.Mutex
}

func newBankAccount(id string, balance int) *bankAccount {
	return &bankAccount{
		id:      id,
		balance: balance,
		mtx:     sync.Mutex{},
	}
}

func (ba *bankAccount) transfer(a *arbitrator, target *bankAccount, amount int, txID string) {
	a.lockAccounts(ba.id, target.id)
	defer a.unlockAccounts(ba.id, target.id)
	fmt.Printf("TX %q: Locking accounts %q, %q", txID, ba.id, target.id)
	ba.balance -= amount
	target.balance += amount
	fmt.Printf("TX %q: Unlocking accounts %q, %q", txID, ba.id, target.id)
}

type arbitrator struct {
	accountsInUse map[string]bool
	cond          *sync.Cond
}

func newArbitrator() *arbitrator {
	return &arbitrator{
		accountsInUse: map[string]bool{},
		cond:          sync.NewCond(&sync.Mutex{}),
	}
}

func (a *arbitrator) lockAccounts(ids ...string) {
	a.cond.L.Lock()
	defer a.cond.L.Unlock()
	for allAvailable := true; !allAvailable; {
		for _, id := range ids {
			if a.accountsInUse[id] {
				allAvailable = false
				a.cond.Wait()
			}
		}
	}
	for _, id := range ids {
		a.accountsInUse[id] = true
	}
}

func (a *arbitrator) unlockAccounts(ids ...string) {
	a.cond.L.Lock()
	defer a.cond.L.Unlock()
	for _, id := range ids {
		a.accountsInUse[id] = false
	}
	a.cond.Broadcast()
}

func main() {
	accnts := []*bankAccount{
		newBankAccount("Sam", defaultBalance),
		newBankAccount("Paul", defaultBalance),
		newBankAccount("Amy", defaultBalance),
		newBankAccount("Mia", defaultBalance),
	}
	accntsCount := len(accnts)
	wg := sync.WaitGroup{}
	arbitratr := newArbitrator()
	for idx := range accnts {
		wg.Add(1)
		go func(accnt *bankAccount) {
			for i := 0; i < 1000; i++ {
				randAccntIdx := rand.Intn(accntsCount)
				if accnts[randAccntIdx].id == accnt.id {
					randAccntIdx = rand.Intn(accntsCount)
				}
				destinationAccnt := accnts[randAccntIdx]
				amount := 10

				txID := fmt.Sprintf("src: %q, dest: %q, amount: %q", accnt.id, destinationAccnt.id, amount)
				accnt.transfer(arbitratr, destinationAccnt, amount, txID)
			}
			fmt.Printf("completed all rnd transactions from accnt %q\n", accnt.id)
			wg.Done()
		}(accnts[idx])
	}
	wg.Wait()
	fmt.Println("all finished.")
}
