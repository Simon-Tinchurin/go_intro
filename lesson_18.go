package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// package context

type result struct {
	userId string
	err    error
}

func thirdpartyHTTPCall() (string, error) {
	// context deadline exceeded
	// time.Sleep(time.Millisecond * 500)

	time.Sleep(time.Millisecond * 90)
	return "some response", nil
}

func fetchResponse(ctx context.Context) (string, error) {
	// ctx - fron context
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("the value from the context ->", val)

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdpartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done()
	// 1. -> the context timout is exceeded
	// 2. -> the context has been manually canceled -> Canceled()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func less_18() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "john")
	res, err := fetchResponse(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v. result: %+v\n", time.Since(start), res)
}
