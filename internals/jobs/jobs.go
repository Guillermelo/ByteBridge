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

func FlushConnJobs(Queue chan<- *Job, Conn net.Conn) {
	// en una rutina distinta tengo que estar escuchando
	// processando la info para llenar el Queue por el momento
	// lo hago aca en una funcion

	// for job := range &Queue {
	// 	switch job.Type {
	// 	case "Message":
	// 	case "ReceiveFile":
	//
	// 	}
	// }
}

func NewJob(p *Packet) *Job {
	currentJob := Job{
		ID:     "null",
		Type:   p.Type,
		Packet: p,
	}

	return &currentJob
}

func FillConnJobs(Queue chan<- *Job, conn net.Conn) {
	// here we create the jobs based on type in loop, and insert them in the Queue
	reader := bufio.NewReader(conn)
	// read the `json:"
	header, _ := reader.ReadBytes('\n')

	var CurrentPacket Packet
	err := json.Unmarshal(header, &CurrentPacket)
	if err != nil {
		fmt.Println(err)
		err = nil
	}

	currentJob := NewJob(&CurrentPacket)
	Queue <- currentJob

	fmt.Println(CurrentPacket)
	wheretosave := "./files/server"
	out, _ := os.Create(wheretosave + "received_" + CurrentPacket.Filename)
	defer out.Close()
	io.CopyN(out, reader, CurrentPacket.Size)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file received")
}

func RunJob(job JobsContract) error {
	return job.Execute()
}

type ReceiveFile struct {
	size     int64
	Filename string
}

func NewFile(packet Packet) *ReceiveFile {
	job := ReceiveFile{packet.Size, packet.Filename}
	return &job
}

func (job *ReceiveFile) Execute() {
}
