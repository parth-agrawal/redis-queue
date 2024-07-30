package backend

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ClickHandler(user string, timestamp int) error {
	// the click handler should essentially take in the user
	// add the timestamp to the UserQueue
	// the max size is 10
	// there should be an if statement - if it was successfully added
	// it should go to the next thing and add to the task queue

	// if it was not sucessfully added it should immediately return an error which should be returned
	
	store := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})


	var ctx = context.Background()

	
	// jsonData, err := json.Marshal("foo")
	// if err != nil {
	// 	return err
	// }
	
	err := store.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

	val, err := store.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

	return err


	
}
