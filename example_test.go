package multiline_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/byounghoonkim/multiline"
)

func ExampleMultiLine() {
	for i := 0; i < 10; i++ {
		line := multiline.GetLine(fmt.Sprintf("%d job - ", i))
		go func(line *multiline.Line) {
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

	multiline.Print()
}
