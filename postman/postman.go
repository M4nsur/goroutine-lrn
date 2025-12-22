package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, letter string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Postman (number):", n, "Finish")
			return
		default:
			fmt.Println("Postman (number):", n, "I picked up the letter")
			time.Sleep(1 * time.Second)
			fmt.Println("Postman (number):", n, "Delivered the letter to the post office", letter)
			transferPoint <- letter
			fmt.Println("Postman (number):", n, "Passed on the letter", letter)
		}
	}
}


func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := range postmanCount {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMain(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}


func postmanToMain(postmanNumber int) string {
	ptm := map[int]string{
	0: "Invitation to the blues",
	1: "Family greeting",
	2: "Invitation from a friend",
	3: "Information from the auto service",
}
	mail, ok := ptm[postmanNumber]

	if !ok {
		return "default letter"
	}

	return mail
}