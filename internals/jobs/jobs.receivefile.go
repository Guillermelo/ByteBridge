package jobs

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ReceiveFileJob struct {
	Type     string `json:"type"`
	Size     int64  `json:"size"`
	Filename string `json:"filename"`
	Userdata string `json:"userdata"`
	Reader   *bufio.Reader
}

func (j *ReceiveFileJob) ReflectType() string {
	return "ReceiveFileJob"
}

func (j *ReceiveFileJob) Execute() error {
	// logic of ReceiveFileJob
	wheretosave := "./files/server/"
	out, _ := os.Create(wheretosave + "received_" + j.Filename)
	_, err := io.CopyN(out, j.Reader, j.Size)
	if err != nil {
		fmt.Println("failed to copy")
		return err
	}
	fmt.Println("file received")
	return nil
}
