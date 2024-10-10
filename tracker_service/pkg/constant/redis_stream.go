package constant

import "time"

const (
	REDIS_KEY_STREAM    = "live_tracking:%s"
	REDIS_NO_DATA       = "no data"
	REDIS_EXPIRE_STREAM = 24 * time.Hour
)
