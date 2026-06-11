package jobs

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ReceiveFileJob struct {
	Size     int64
	Filename string
	Userdata string
	Reader   *bufio.Reader
}

func (j *ReceiveFileJob) Execute() error {
	// logic of ReceiveFileJob
	wheretosave := "./files/server"
	out, _ := os.Create(wheretosave + "received_" + j.Filename)
	_, err := io.CopyN(out, j.Reader, j.Size)
	if err != nil {
		fmt.Println("failed to copy")
		return err
	}
	fmt.Println("file received")
	return nil
}
