package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var store *redis.Client


func init () { 
	totalGlobalClickCount = 0
	store = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})
	go startQueueCleaner()
	go slurpClickTimed()
}



func ClickHandler(user string, timestamp int) error {
	// the click handler should essentially take in the user
	// add the timestamp to the UserQueue
	// the max size is 10
	// there should be an if statement - if it was successfully added
	// it should go to the next thing and add to the task queue

	// if it was not sucessfully added it should immediately return an error which should be returned
	



	var ctx = context.Background()


	fmt.Println("we're here in the clickHandler. user and timestamp are ", user, timestamp)

	// read queue size
	queueLength, err := store.LLen(ctx, user).Result()
	if err!=nil {
		return fmt.Errorf("error reading queue for user %s", user)
	}


	if(queueLength >= 10) {
		return fmt.Errorf("rate limit for user exceeded", user)
	}

	// if it's the right length, less than 10, add it to the userQueue

	err= store.LPush(ctx, user, timestamp).Err()
	if err != nil {
		return fmt.Errorf("failed to add timestamp to queue: %w", err)
	}

	err = loadTaskQueue(user, timestamp)
	if err != nil{ 
		return fmt.Errorf("failed to add to taskqueue: %w", err)
	}

	

	return err


	
}


func loadTaskQueue (user string, timestamp int) error { 

	ctx := context.Background()

	// load into taskqueue

	taskData := map[string]interface{}{
		"user":      user,
		"timestamp": timestamp,
	}
	jsonData, err := json.Marshal(taskData)
	if err != nil {
		return fmt.Errorf("failed to marshal task data: %w", err)
	}
	err = store.LPush(ctx, "taskQueue", jsonData).Err()

	return err


}

