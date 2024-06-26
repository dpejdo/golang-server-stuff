package main

import (
	"fmt"
	"sync"
)

type balance struct {
	amount   int
	currency string
	mu       sync.Mutex
}

func (b *balance) getBalance() string {
	b.mu.Lock()

	defer b.mu.Unlock()

	return fmt.Sprintf("User balance is %d%s", b.amount, b.currency)
}

func (b *balance) setBalance(val int) {
	b.mu.Lock()
	b.amount += val
	b.mu.Unlock()
}
