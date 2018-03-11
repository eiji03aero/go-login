package util

import (
	"fmt"
	"time"

	"github.com/c-bata/go-prompt"
)

func WaitMS(s time.Duration) {
	time.Sleep(s * time.Millisecond)
}

func NewLine() {
	fmt.Println("")
}

func DefaultCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
