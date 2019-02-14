package beater

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/elastic/beats/libbeat/logp"
)

// Executor exec command logcat
type Executor struct {
	Command string
	Stdout  io.Reader
}

// NewExecutor init a new Executor
func NewExecutor(command string) *Executor {
	return &Executor{
		Command: command,
	}
}

// Run start
func (e *Executor) Run(msgs chan string) {
	logp.Info(e.Command)
	cmd := exec.Command("sh", "-c", e.Command)
	e.Stdout, _ = cmd.StdoutPipe()
	f := bufio.NewReader(e.Stdout)
	if err := cmd.Start(); err != nil {
		logp.Err(err.Error())
	}
	for {
		line, e := f.ReadString('\n')
		switch e {
		case io.EOF:
			err := cmd.Wait()
			if err != nil {
				logp.Err("An error occured while executing command: %v", err)
				return
			}
		case nil:
			msgs <- line
		default:
			continue
		}
	}
}
