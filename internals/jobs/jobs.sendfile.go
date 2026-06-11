package jobs

import "bufio"

type SendFileJob struct {
	Size     int64
	Filename string
	Userdata string
	Reader   *bufio.Reader
}

func (j *SendFileJob) Execute() error {
	// logic for sending files
	return nil
}
