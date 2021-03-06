
[![Go Report Card](https://goreportcard.com/badge/github.com/efy/bookmark)](https://goreportcard.com/report/github.com/efy/bookmark)

Package bookmark provides a function Parse that is capable of parsing exported bookmarks
from a variety of web browsers and bookmaring services. Uses the same approach as
the PHP [netscape-bookmark-parser](https://github.com/kafene/netscape-bookmark-parser)

# Example

```go

package main

import (
  "fmt"
  "os"
  "github.com/efy/bookmarks"
)

func main() {
  file, _ os.Open()

  bookmarks, err := bookmarks.Parse(file)

  for _, b := range bookmarks {
    fmt.Println(b.Title)
  }
}

```

```go
// Bookmark type

Bookmark{
  Title   string
  Url     string
  Icon    string
  Created time.Time
  Tags    []string
}

```
