package bookmark

import (
  "testing"
)

var rowTable = []struct {
  desc string
  in string
  out Bookmark
}{
  {
    "Chrome entry",
    `<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009412">Regex Crossword</A>`,
    Bookmark{},
  },
}

func TestParseRowChrome(t *testing.T) {
  row := `<DT><A HREF="https://regexcrossword.com/" ADD_DATE="1466009412">Regex Crossword</A>`
  bm, err := ParseRow(row)

  if err != nil {
    t.Fatal()
  }

  if bm.Title != "Regex Crossword" {
    t.Error("expected", "Regex Crossword", "got", bm.Title)
  }

  if bm.Url != "https://regexcrossword.com/" {
    t.Error("expected", "https://regexcrossword.com/", "got", bm.Url)
  }

  if bm.Created.String() != "2016-06-15 17:50:12 +0100 BST" {
    t.Error("expected", "2016-06-15 17:50:12 +0100 BST", "got", bm.Created.String())
  }
}

