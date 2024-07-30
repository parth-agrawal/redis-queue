package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)


var totalGlobalClickCount int


func startQueueCleaner() {
	fmt.Println("starting cleaner")

	ticker := time.NewTicker(10 * time.Second)
	ctx := context.Background()

	for range ticker.C {
		fmt.Println("starting cleaning, iterating users")
		users, err := store.Keys(ctx, "*").Result() // Simplified pattern to match all keys
		if err != nil {
			fmt.Printf("error getting user keys: %v", err)
			continue
		}

		fmt.Println("starting to remove items")

		for _, user := range users {
			_, err := store.RPop(ctx, user).Result()
			if err != nil && err != redis.Nil {
				fmt.Printf("error removing item from user %s: %v\n", user, err)
			}
		}
	}
}

func GetTotalClicks () int {
	return totalGlobalClickCount

}

func slurpClickTimed() error { 

	fmt.Println("starting slurpClickTimed")

	ticker := time.NewTicker(2 * time.Second)
	ctx := context.Background()

	for range ticker.C {
		poppedData, err := store.RPop(ctx, "taskQueue").Result()
		if err != nil {
			if err == redis.Nil {
				// No items in the task queue, continue to the next iteration
				continue
			}
			fmt.Printf("error popping taskqueue: %v", err)
			return err // Return the error if it's not redis.Nil
		}
		fmt.Println("popped data: ", poppedData)
		totalGlobalClickCount++
		fmt.Println(totalGlobalClickCount)
	}

	return nil // Return nil if everything went well
}
