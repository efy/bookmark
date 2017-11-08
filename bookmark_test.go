package bookmark

import (
	"testing"
	"time"
)

func TestBookmarkType(t *testing.T) {
	bm := Bookmark{"google.com", "http://google.com", time.Time{}}

	if bm.Title != "google.com" {
		t.Error("expected", "google.com", "got", bm.Title)
	}

	if bm.Url != "http://google.com" {
		t.Error("expected", "http://google.com", "got", bm.Url)
	}

	if bm.Created.String() != "0001-01-01 00:00:00 +0000 UTC" {
		t.Error("expected", "0001-01-01 00:00:00 +0000 UTC", "got", bm.Created)
	}
}
