package multiline

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

func TestMultiLine(t *testing.T) {
	ml := New()

	for i := 0; i < 10; i++ {
		line := ml.GetLine(fmt.Sprintf("%d job - ", i))
		go func(line *Line) {
			defer line.Close()

			line.WriteS("🚚 Preparing ...")
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

			for j := 0; j < 10; j++ {
				line.WriteS("⛏️  " + msgList[rand.Intn(len(msgList))])
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}

			line.WriteS("✅ DONE")
		}(line)
	}

	ml.Print()
}