package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
		fmt.Println("Miner (number):", n, "Finish")
			return 
		default:
			fmt.Println("Miner (number):", n, "Start work")
			time.Sleep(1 * time.Second)
			fmt.Println("Miner (number):", n, "Mined  coal", power)
			transferPoint <- power
			fmt.Println("Miner (number):", n, "Passed coal to transferPoint", power)
		}
	}
}


func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}
	for i := range minerCount {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}
	
	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}