package bookmark

import (
	"fmt"
	"os"
	"reflect"
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
			`<DT><A HREF="https://goodbookmark.com/" ADD_DATE="1466009412" ICON="testiconstring">Good Bookmark</A>`,
			Bookmark{
				Title: "Good Bookmark",
				Url:   "https://goodbookmark.com/",
				Icon:  "testiconstring",
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
		if bm.Icon != tr.out.Icon {
			t.Error("expected", tr.out.Icon, "got", bm.Icon)
		}
	}
}

func TestParseLineWithTags(t *testing.T) {
	var tt = []struct {
		in  string
		out Bookmark
		err error
	}{
		{
			`<DT><A HREF="https://goodbookmark.com/" TAGS="		tabs		">Bad Bookmark</A>`,
			Bookmark{
				Title: "Bad Bookmark",
				Url:   "https://goodbookmark.com/",
				Tags: []string{
					"tabs",
				},
			},
			nil,
		},
		{
			`<DT><A HREF="https://goodbookmark.com/" TAGS="one">Good Bookmark</A>`,
			Bookmark{
				Title: "Good Bookmark",
				Url:   "https://goodbookmark.com/",
				Tags: []string{
					"one",
				},
			},
			nil,
		},
		{
			`<DT><A HREF="https://goodbookmark.com/" TAGS="one,two, three">Good Bookmark</A>`,
			Bookmark{
				Title: "Good Bookmark",
				Url:   "https://goodbookmark.com/",
				Tags: []string{
					"one",
					"two",
					"three",
				},
			},
			nil,
		},
	}

	for _, tr := range tt {
		bm, err := parseLine(tr.in)

		if err != tr.err {
			t.Error(err)
		}

		if !reflect.DeepEqual(bm.Tags, tr.out.Tags) {
			t.Error("expected", tr.out.Tags)
			t.Error("got     ", bm.Tags)
		}
	}
}

func TestParseTotals(t *testing.T) {
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
		t.Run(tr.file, func(t *testing.T) {
			file, err := os.Open(tr.file)
			defer file.Close()
			if err != nil {
				t.Error(err)
			}

			got, err := Parse(file)
			if len(got) != tr.count {
				t.Error("expected", tr.count, "got", len(got))
			}
		})
	}
}

func TestParsePropgatesReaderError(t *testing.T) {
	br := badReader{}
	_, err := Parse(br)
	if err.Error() != "reader error" {
		t.Errorf("expected %q to be propogated", "reader error")
	}
}

type badReader struct{}

func (b badReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("reader error")
}

func TestFiles(t *testing.T) {
	tt := map[string][]struct {
		index    int
		bookmark Bookmark
	}{
		"testfiles/firefox_nested.htm": {
			{
				4,
				Bookmark{
					Url: "http://xkcd.com/1332/",
				},
			},
		},
		"testfiles/chromium_nested.htm": {
			{
				0,
				Bookmark{
					Url: "http://www.jabber.org/",
				},
			},
			{
				6,
				Bookmark{
					Url: "https://checkio.org/",
				},
			},
			{
				16,
				Bookmark{
					Url: "https://github.com/shaarli/Shaarli",
				},
			},
		},
		"testfiles/delicious.htm": {
			{
				4,
				Bookmark{
					Url: "http://fontfamily.io/",
				},
			},
		},
		"testfiles/netscape_nested.htm": {
			{
				0,
				Bookmark{
					Url: "http://nest.ed/1",
				},
			},
		},
		"testfiles/netscape_multiline.htm": {
			{
				1,
				Bookmark{
					Url: "http://multi.li.ne/2",
				},
			},
		},
		"testfiles/shaarli.htm": {
			{
				4,
				Bookmark{
					Url: "https://github.com/shaarli/Shaarli/wiki",
				},
			},
		},
	}

	for k, v := range tt {
		t.Run(k, func(t *testing.T) {
			file, err := os.Open(k)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			got, err := Parse(file)
			if err != nil {
				t.Fatal(err)
			}

			for _, tr := range v {
				if got[tr.index].Url != tr.bookmark.Url {
					t.Error("expected", tr.bookmark.Url)
					t.Error("got     ", got[tr.index].Url)
				}
			}
		})
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
