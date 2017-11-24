package bookmark

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrBookmarkEmpty = fmt.Errorf("bookmark empty")
)

func ParseLine(r string) (Bookmark, error) {
	var bm Bookmark

	tr := regexp.MustCompile(`(?i)<a.*>(.*?)<\/a>`)
	ur := regexp.MustCompile(`(?i)href="(.*?)"`)
	tsr := regexp.MustCompile(`(?i)add_date="(.*?)"`)

	titlematch := tr.FindStringSubmatch(r)
	if len(titlematch) > 1 {
		bm.Title = titlematch[1]
	}

	urlmatch := ur.FindStringSubmatch(r)
	if len(urlmatch) > 1 {
		bm.Url = urlmatch[1]
	}

	ts := tsr.FindStringSubmatch(r)
	if len(ts) > 1 {
		tsi, err := strconv.ParseInt(ts[1], 10, 64)
		if err == nil {
			bm.Created = time.Unix(tsi, 0)
		}
	}

	if (Bookmark{}) == bm || bm.Url == "" {
		return bm, ErrBookmarkEmpty
	}

	return bm, nil
}

func Parse(r io.Reader) ([]Bookmark, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return []Bookmark{}, err
	}

	return ParseLines(string(b))
}

func ParseLines(str string) ([]Bookmark, error) {
	lines := strings.Split(sanatize(str), "\n")
	var bms []Bookmark

	for _, line := range lines {
		// Skip empty newlines
		if line == "" {
			continue
		}
		bm, err := ParseLine(line)
		if err != nil {
			continue
		}
		bms = append(bms, bm)
	}

	return bms, nil
}

// Normalizes the bookmark file contents
func sanatize(str string) string {
	// Trim spaces and and newlines from beginning and end
	s := strings.Trim(str, " \n\t")

	// Remove carriage returns
	s = strings.Replace(s, "\r", "", -1)

	// Replace tabs with a space
	s = strings.Replace(s, "\t", " ", -1)
	return s
}
