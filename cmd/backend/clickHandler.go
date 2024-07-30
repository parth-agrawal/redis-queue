package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var store *redis.Client


func init () { 
	store = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	go startQueueCleaner()
}

func startQueueCleaner() {

	fmt.Println("starting cleaner")

	ticker := time.NewTicker(10 * time.Second)
	ctx := context.Background()

	for range ticker.C { 
		
		fmt.Println("starting cleaning, iterating users")
		users, err := store.Keys (ctx, "*").Result()
		if err != nil {
			fmt.Printf("error getting user key: %v", err)
			continue
		}

		fmt.Println("starting to remove items")
		
		for _, user := range users { 
			_, err := store.RPop(ctx, user).Result()
			if err != nil && err != redis.Nil { 
				fmt.Printf("error removing item %v", err)
			}
		}
	}


}


func ClickHandler(user string, timestamp int) error {
	// the click handler should essentially take in the user
	// add the timestamp to the UserQueue
	// the max size is 10
	// there should be an if statement - if it was successfully added
	// it should go to the next thing and add to the task queue

	// if it was not sucessfully added it should immediately return an error which should be returned
	



	var ctx = context.Background()

	err := store.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set key-value pair: %w", err)
	}

	// Create a queue called fruits and add apple and banana to it
	err = store.LPush(ctx, "fruits", "apple", "banana").Err()
	if err != nil {
		return fmt.Errorf("failed to create fruits queue: %w", err)
	}

	fmt.Println("we're here in the clickHandler. user and timestamp are ", user, timestamp)

	// read queue size
	queueLength, err := store.LLen(ctx, user).Result()
	if err!=nil {
		return fmt.Errorf("error reading queue for user %s", user)
	}


	if(queueLength >= 10) {
		return fmt.Errorf("rate limit for user %s exceeded", user)
	}

	// if it's the right length, less than 10, add it to the userQueue

	err= store.LPush(ctx, user, timestamp).Err()
	if err != nil {
		return fmt.Errorf("failed to add timestamp to queue: %w", err)
	}





	return err


	
}

