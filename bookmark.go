package bookmark

import (
	"time"
)

type Bookmark struct {
	Title   string
	Url     string
	Created time.Time
}
