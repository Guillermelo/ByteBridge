package jobs

import (
	"bufio"
)

type SendMessageJob struct {
	Userdata string
	Reader   *bufio.Writer
}

func (j *SendMessageJob) Execute() error {
	return nil
}
