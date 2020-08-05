package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/peterh/liner"
)

func prompt() string {
	return fmt.Sprintf("<%s> ", time.Now().Format("Jan 2 15:04:05.000"))
}

func main() {

	// Initialize commands
	commandsInit()

	// Start loop with history and completion
	_ = interactiveLoop()
}

func interactiveLoop() error {
	s := liner.NewLiner()
	s.SetTabCompletionStyle(liner.TabPrints)
	s.SetCompleter(func(line string) (ret []string) {
		for _, c := range commandKeys {
			if strings.HasPrefix(c, line) {
				ret = append(ret, c)
			}
		}
		return
	})
	defer func() { _ = s.Close() }()
	for {
		p, err := s.Prompt(prompt())
		if err == io.EOF {
			return nil
		}
		if err != nil {
			panic(err)
		}
		if executeCommand(p) {
			s.AppendHistory(p)
		}
	}
}
