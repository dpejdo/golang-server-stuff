package main

import (
	"sync"
	"testing"
)

func TestBalance_Concurrency(t *testing.T) {
	b := &balance{amount: 1000, currency: "USD"}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.setBalance(100)
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.getBalance()
		}()
	}

	wg.Wait()

	expected := 1000 + (100 * 100) // Initial amount + 100 Add(100) operations
	if b.amount != expected {
		t.Errorf("expected %d, got %d", expected, b.amount)
	}
}
