package main

import (
	"context"
	"fmt"
	"github/m4nsur/concurrency-lrn/miner"
	"github/m4nsur/concurrency-lrn/postman"
	"time"
)

func main() {
	var coals int
	var mails []string

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

	isCoalClosed := false 
	isMailsClosed := false 

	for !isCoalClosed || !isMailsClosed {
		select {
		case coal, ok := <- coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}

			coals += coal

		case letter, ok := <- mailsTransferPoint:
			if !ok {
				isMailsClosed = true
				continue
			}

			mails = append(mails, letter)
		}

	}

	fmt.Println("Coals:", coals)
	fmt.Println("Mails:", len(mails))

}