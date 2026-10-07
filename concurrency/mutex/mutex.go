package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main(){
	var bankBalance int = 0
	var balanceMut sync.Mutex

    fmt.Printf("Initial bank account balance : ₹%v", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main job", Amount: 5000},
		{Source: "Gifts", Amount: 50},
		{Source: "Side job", Amount: 500},
		{Source: "Investments", Amount: 5},
	}
	for i, income := range(incomes) {
        wg.Add(1)
		go func(i int, income Income){
			defer wg.Done()
			for week := 1 ; week <= 52; week++ {
                balanceMut.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balanceMut.Unlock()
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Bank balance after a year : ₹%v", bankBalance)
    
}
