package main

import (
	"context"
	"fmt"
	"github/m4nsur/concurrency-lrn/miner"
	"github/m4nsur/concurrency-lrn/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coals atomic.Int64
	var mails []string
	var mu sync.Mutex
	contextMiner, cancelMiner := context.WithCancel(context.Background())
	contextPostman, cancelPostman := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		cancelMiner()
	}()
		go func() {
		time.Sleep(6 * time.Second)
		cancelPostman()
	}()

	
	coalTransferPoint := miner.MinerPool(contextMiner, 10)
	mailsTransferPoint := postman.PostmanPool(contextPostman, 10)

	wg := sync.WaitGroup{}

	wg.Go(func() {
		for coal := range coalTransferPoint {
			coals.Add(int64(coal))
		}
	})

	wg.Go(func() {
		for letter := range mailsTransferPoint {
			mu.Lock()
			mails = append(mails, letter)
			mu.Unlock()
		}
	})	

	wg.Wait()
	
	fmt.Println("Coals:", coals.Load())
	fmt.Println("Mails:", len(mails))

}