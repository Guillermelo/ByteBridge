// Package jobs is for dividingthe work dpending on the type
package jobs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

type Job struct {
	ID      string
	Payload []byte
	Type    string
}

type Packet struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type Jobs interface {
	Execute() error
}

func FlushConnJobs(Queue <-chan Job, Conn net.Conn) {
	// en una rutina distinta tengo que estar escuchando
	// processando la info para llenar el Queue por el momento
	// lo hago aca en una funcion

	for job := range Queue {
		switch job.Type {
		case "Message":
			ReceiveMessage(job)
		case "ReceiveFile":
			ReceiveFile(job)

		}
	}
}

func FillConnJobs(Queue <-chan Job, conn net.Conn) {
	// here we create the jobs based on type in loop, and insert them in the Queue
	reader := bufio.NewReader(conn)

	// read the `json:"
	header, _ := reader.ReadBytes('\n')

	var packet Packet
	json.Unmarshal(header, &packet)
	fmt.Println(packet)
	out, _ := os.Create("received_" + packet.Filename)
	defer out.Close()
	io.CopyN(out, reader, packet.Size)
	fmt.Println("file received")
}

type FileTransferJob struct {
	Payload []byte
}

func (job FileTransferJob) Execute() error {
	fmt.Println("funciona?")
	return nil
}

func ReceiveMessage(job Job) {
}

func ReceiveFile(job Job) {
}

func IsValidJob(Job Job) bool {
	return true
}
