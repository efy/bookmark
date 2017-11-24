package bookmark

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	var tests = []struct {
		in  string
		out Bookmark
		err error
	}{
		{
			``,
			Bookmark{},
			ErrBookmarkEmpty,
		},
		{
			`<DT><A HREF="https://goodbookmark.com/" ADD_DATE="1466009412">Good Bookmark</A>`,
			Bookmark{
				Title: "Good Bookmark",
				Url:   "https://goodbookmark.com/",
			},
			nil,
		},
	}

	for _, tr := range tests {
		bm, err := ParseLine(tr.in)

		if err != tr.err {
			t.Error(err)
		}
		if bm.Title != tr.out.Title {
			t.Error("expected", tr.out.Title, "got", bm.Title)
		}
		if bm.Url != tr.out.Url {
			t.Error("expected", tr.out.Url, "got", bm.Url)
		}
	}
}

func TestParseLines(t *testing.T) {
	lines := `
			<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009413">Regex Crossword</A>
			<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009412">Regex Crossword</A>
			<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009412">Regex Crossword</A>
  `

	got, err := ParseLines(lines)
	if err != nil {
		t.Error(err)
	}
	if len(got) != 3 {
		t.Error("expected 3 bookmarks got", len(got))
	}
}

func TestParse(t *testing.T) {

}
