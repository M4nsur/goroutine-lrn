package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()
	select {
	case <- ctx.Done():
		return 
	default:
		for {
			fmt.Println("Miner:", n, "start work")
			time.Sleep(1 * time.Second)
			fmt.Println("Miner:", n, "mined  coal", power)
			transferPoint <- power
			fmt.Println("Miner:", n, "passed coal to transferPoint", power)
		}
	}
}


func MinerPool(ctx context.Context, minerCount int) <-chan int {
	transferCoal := make(chan int)
	wg := &sync.WaitGroup{}
	for i := range minerCount {
		wg.Add(1)
		go Miner(ctx, wg, transferCoal, i, i*10)
	}
	
	go func() {
		wg.Wait()
		close(transferCoal)
	}()

	return transferCoal
}