package handlers

import (
	"context"
	"encoding/json"
	"eud_backup/pkg/cache"
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
			w.Write([]byte("error occurred"))
			return
		}

		w.Write(resp)
		return
	}

	w.Write([]byte(data))
}

// notFound marshals the 404 response
func notFound() ([]byte, error) {
	data := map[string]interface{}{
		"code":    404,
		"message": "stats could not be fetched",
	}

	return json.Marshal(data)
}
