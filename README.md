multiline
=========

Small library to print multiple lines for job status from goroutines.

[![asciicast](https://asciinema.org/a/UX57WjsHOHqO9q3o7EC2ZTd8V.svg)](https://asciinema.org/a/UX57WjsHOHqO9q3o7EC2ZTd8V)

## Quick Started

Install.
```
go get github.com/byounghoonkim/multiline
```


Example Code.

``` golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/byounghoonkim/multiline"
)

func main() {
	ml := multiline.New()

	for i := 0; i < 10; i++ {
		line := ml.GetLine(fmt.Sprintf("%d job - ", i))
		go func(line *Line) {
			defer line.Close()

			fmt.Fprint(line, "🚚 Preparing ...")
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

			for j := 0; j < 10; j++ {
				fmt.Fprintf(line, "⛏️  %s", msgList[rand.Intn(len(msgList))])
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}

			fmt.Fprint(line, "✅ DONE")
		}(line)
	}

	ml.Print()
}

var msgList = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
	"Phasellus semper mauris eu tellus eleifend, sit amet gravida massa pellentesque. ",
	"Praesent ornare lacinia odio ut cursus. ",
	"Proin quis est leo. ",
	"Cras egestas elit eget dui laoreet blandit. ",
	"Etiam mattis mattis viverra. ",
	"Donec a elit eget massa vestibulum luctus sit amet a mauris. ",
	"Maecenas scelerisque pretium pellentesque. ",
	"Donec in nisl eget velit dapibus feugiat ac non ligula.",
}
```
