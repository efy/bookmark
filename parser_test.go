package bookmark

import (
	"testing"
)

var tests = []struct {
	desc string
	in   string
	out  Bookmark
}{
	{
		"Chromium",
		`<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009412">Regex Crossword</A>`,
		Bookmark{
			Title: "Regex Crossword",
			Url:   "https://regexcrossword.com/",
		},
	},
}

func TestParseRow(t *testing.T) {
	for _, tt := range tests {
		bm, err := ParseRow(tt.in)
		if err != nil {
			t.Error(tt.desc, err)
		}
		if bm.Title != tt.out.Title {
			t.Error(tt.desc, "expected", tt.out.Title, "got", bm.Title)
		}
		if bm.Url != tt.out.Url {
			t.Error(tt.desc, "expected", tt.out.Url, "got", bm.Url)
		}
	}
}
