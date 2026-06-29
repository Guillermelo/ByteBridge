package jobs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type SendFileJob struct {
	Userdata string // this should go tu a decorator or some stuff like it
	Writer   *bufio.Writer
	Filepath string
}

// func (j *SendFileJob) NewJob() (Job, error) {
// 	return &SendFileJob{
// 		Writer:   j.Writer,
// 		Filepath: "/random/path/for/now",
// 	}, nil
// }

func (j *SendFileJob) Execute() error {
	// logic for sending files
	file, err := os.Open(j.Filepath)
	if err != nil {
		fmt.Println("error opening the file at Execute")
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	packet := Packet{
		Type:     "ReceiveFileJob",
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
	}
	jsondata, err := json.Marshal(packet)
	if err != nil {
		return err
	}
	jsondata = append(jsondata, '\n')
	_, err = j.Writer.Write(jsondata)
	if err != nil {
		return err
	}
	_, err = io.Copy(j.Writer, file)
	if err != nil {
		return err
	}
	fmt.Println(fileInfo.Name(), "sent successfully")
	return nil
}
