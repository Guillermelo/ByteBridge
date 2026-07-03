package jobs

import (
	"bufio"
)

type SendMessageJob struct {
	Userdata string
	Writer   *bufio.Writer
}

func (j *SendMessageJob) Execute() error {
	return nil
}
