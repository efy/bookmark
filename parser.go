package bookmark

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func ParseRow(r string) (Bookmark, error) {
	var bm Bookmark

	tr := regexp.MustCompile(`(?i)<a.*>(.*?)<\/a>`)
	ur := regexp.MustCompile(`(?i)href="(.*?)"`)
	tsr := regexp.MustCompile(`(?i)add_date="(.*?)"`)

	title := tr.FindStringSubmatch(r)[1]
	url := ur.FindStringSubmatch(r)[1]
	ts := tsr.FindStringSubmatch(r)[1]

	tsi, err := strconv.ParseInt(ts, 10, 64)

	if err != nil {
		return bm, errors.New("Could not parse timestamp to integer")
	}

	created := time.Unix(tsi, 0)

	bm = Bookmark{title, url, created}

	return bm, nil
}
