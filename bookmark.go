package bookmark

import (
	"time"
)

type Bookmarks []Bookmark

type Bookmark struct {
	Title   string
	Url     string
	Created time.Time
}
