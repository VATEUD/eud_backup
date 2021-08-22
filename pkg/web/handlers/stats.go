package handlers

import (
	"context"
	"encoding/json"
	"eud_backup/pkg/cache"
	"log"
	"net/http"
)

const redisBackupKey = "backup_stats"

// ctx represents the context
var ctx = context.Background()

// client represents the Redis client
var client = cache.New()

// Stats returns the cached information about the last backup
func Stats(w http.ResponseWriter, r *http.Request) {
	data, err := client.Get(ctx, redisBackupKey).Result()

	if err != nil {
		resp, err := notFound()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write([]byte("error occurred")); err != nil {
				log.Println(err.Error())
			}
			return
		}

		if _, err = w.Write(resp); err != nil {
			log.Println(err.Error())
		}
		return
	}

	if _, err = w.Write([]byte(data)); err != nil {
		log.Println(err.Error())
	}
}

// notFound marshals the 404 response
func notFound() ([]byte, error) {
	data := map[string]interface{}{
		"code":    404,
		"message": "stats could not be fetched",
	}

	return json.Marshal(data)
}
