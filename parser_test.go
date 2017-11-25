package bookmark

import (
	"os"
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
		bm, err := parseLine(tr.in)

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

func TestParse(t *testing.T) {
	tt := []struct {
		file  string
		count int
	}{
		{
			"testfiles/chromium_flat.htm",
			9,
		},
		{
			"testfiles/firefox_flat.htm",
			24,
		},
		{
			"testfiles/internet_explorer_11_flat.htm",
			18,
		},
		{
			"testfiles/netscape_basic.htm",
			2,
		},
		{
			"testfiles/netscape_nested.htm",
			8,
		},
		{
			"testfiles/shaarli.htm",
			6,
		},
	}

	for _, tr := range tt {
		file, err := os.Open(tr.file)
		if err != nil {
			t.Error(err)
		}

		got, err := Parse(file)
		if len(got) != tr.count {
			t.Error("expected", tr.count, "got", len(got))
		}
	}
}

func BenchmarkParse(b *testing.B) {
	file, err := os.Open("testfiles/chromium_flat.htm")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = Parse(file)
	}
}
