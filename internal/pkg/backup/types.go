package backup

import (
	"context"
	"encoding/json"
	"eud_backup/pkg/cache"
)

const (
	redisBackupKey  = "backup_stats"
	redisExpiration = 0
)

// Config represents the config which holds databases list
type Config struct {
	Databases []string `yaml:"databases"`
}

// Stats represents the backup stats
type Stats struct {
	BackupTime     string `json:"backup_time"`
	NextBackupTime string `json:"next_backup_time"`
	Success        bool   `json:"success"`
}

func (stats Stats) store() error {
	// ctx represents the context
	var ctx = context.Background()

	// client represents the Redis client
	var client = cache.New()

	bytes, err := json.Marshal(stats)

	if err != nil {
		return err
	}

	client.Set(ctx, redisBackupKey, bytes, redisExpiration)

	return nil
}
