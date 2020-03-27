package log_tail

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func NewTail(path string)  {
	t, _ := tail.TailFile(path, tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}




