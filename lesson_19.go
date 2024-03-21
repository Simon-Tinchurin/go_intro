package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// practical example of goroutines

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

// ----------SYNC----------- 501.523063ms
// func getUserProfile(id int) (*UserProfile, error) {
// 	comments, err := getComments(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	likes, err := getLikes(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	friends, err := getFriends(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &UserProfile{
// 		ID:       id,
// 		Comments: comments,
// 		Likes:    likes,
// 		Friends:  friends,
// 	}, nil
// }

// func getComments(id int) ([]string, error) {
// 	time.Sleep(time.Millisecond * 200)
// 	comments := []string{
// 		"Hey!",
// 		"Yep",
// 		"Ok",
// 	}
// 	return comments, nil
// }

// func getLikes(id int) (int, error) {
// 	time.Sleep(time.Millisecond * 200)
// 	return 10, nil
// }

// func getFriends(id int) ([]int, error) {
// 	time.Sleep(time.Millisecond * 100)
// 	return []int{11, 34, 543, 123}, nil
// }

// func main() {
// 	start := time.Now()
// 	userProfile, err := getUserProfile(10)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%+v\n", userProfile)
// 	fmt.Println("Fetching the user profile took:", time.Since(start))
// }

// -------------ASYNC-------------- 200.865229ms
type Response struct {
	data any
	err  error
}

func getUserProfile(id int) (*UserProfile, error) {
	var (
		respch    = make(chan Response, 3)
		waitGroup = &sync.WaitGroup{}
	)

	// we are doing 3 requests inside this goroutine
	go getComments(id, respch, waitGroup)
	go getLikes(id, respch, waitGroup)
	go getFriends(id, respch, waitGroup)
	// adding 3 to the wait group
	waitGroup.Add(3)
	waitGroup.Wait() // blick until the wait group counter == 0
	close(respch)

	userProfile := UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg

		}
	}
	return &userProfile, nil
}

func getComments(id int, respch chan Response, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey!",
		"Yep",
		"Ok",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}
	// work is done
	waitGroup.Done()
}

func getLikes(id int, respch chan Response, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: 100,
		err:  nil,
	}
	// work is done
	waitGroup.Done()
}

func getFriends(id int, respch chan Response, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	friedndsIds := []int{11, 34, 543, 123}
	respch <- Response{
		data: friedndsIds,
		err:  nil,
	}
	// work is done
	waitGroup.Done()
}

func main() {
	start := time.Now()
	userProfile, err := getUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", userProfile)
	fmt.Println("Fetching the user profile took:", time.Since(start))
}
