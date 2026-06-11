package jobs

import (
	"bufio"
)

type SendMessageJob struct {
	Userdata string
	Reader   *bufio.Reader
}

func (j *SendMessageJob) Execute() error {
	return nil
}
